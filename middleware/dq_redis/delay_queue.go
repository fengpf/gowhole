package dq_redis

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"gowhole/middleware/dq_redis/myredis"
	"gowhole/middleware/log"
	"strconv"
	"sync"
	"time"
)

type Job struct {
	Topic     string `json:"topic"`     //Job类型。可以理解成具体的业务名称
	ID        int64  `json:"id"`        //Job的唯一标识。用来检索和删除指定的Job信息
	Delay     int64  `json:"delay"`     //Job需要延迟的时间。单位：秒。（服务端会将其转换为绝对时间）
	TimeStamp int64  `json:"timestamp"` //Job需要延迟的绝对时间戳）
	TTR       int64  `json:"ttr"`       //（time-to-run)：Job执行超时时间。单位：秒。 TTR的设计目的是为了保证消息传输的可靠性
	Body      string `json:"body"`      //Job的内容，供消费者做具体的业务处理，以json格式存储。
}

//消息状态转换
//每个Job只会处于某一个状态下：

//ready：可执行状态，等待消费。
//delay：不可执行状态，等待时钟周期。
//reserved：保留，已被消费者读取，但还未得到消费者的响应（delete、finish）。
//deleted：已被消费完成或者已被删除。

const (
	_ uint8 = iota
	Ready
	Delay
	Reserved
	Deleted
)

const (
	_bucketPrefix = "bucket-"
)

type Config struct {
	BucketCount int
}

type Bucket struct {
	timestamp int64
	jobID     string
}

type DQ struct {
	tickers     []*time.Ticker
	bucketCount int //桶数量

	redisPool *redis.Pool

	ReadyCount    int64 //可执行状态，等待消费。
	DelayCount    int64 //不可执行状态，等待时钟周期。
	ReservedCount int64 //保留，已被消费者读取，但还未得到消费者的响应（delete、finish）。
	DeletedCount  int64 //已被消费完成或者已被删除。

	mux  sync.RWMutex
	wg   sync.WaitGroup
	exit chan struct{}
}

func NewDQ(c *Config) *DQ {
	dq := &DQ{
		bucketCount: c.BucketCount,
		redisPool:   myredis.New(),
		exit:        make(chan struct{}),
	}

	dq.monitoring()

	return dq
}

func (dq *DQ) Start() {
	dq.wg.Add(dq.bucketCount)

	dq.tickers = make([]*time.Ticker, dq.bucketCount)
	for i := 0; i < dq.bucketCount; i++ {
		dq.tickers[i] = time.NewTicker(1 * time.Second)

		log.Info("start:%s", dq.getBucketName(i))
		go dq.runTick(dq.tickers[i], dq.getBucketName(i))
	}

}

func (dq *DQ) Close() {
	dq.exit <- struct{}{}
}

func (dq *DQ) runTick(ticker *time.Ticker, bucketName string) {
	defer dq.wg.Done()

	for {
		select {
		case t := <-ticker.C:
			dq.scanBucket(t, bucketName)
		case <-dq.exit:
			log.Warn("runTick bucket:%s exit", bucketName)
			return
		}
	}
}

// 扫描桶, 取出延迟时间小于当前时间的Job
func (dq *DQ) scanBucket(t time.Time, bucketName string) {
	for {
		bucket, err := dq.getBucket(bucketName)
		if err != nil {
			log.Error("scanBucket getFromBucket bucketName(%v) err(%v)", bucketName, err)
			return
		}

		if bucket == nil { //桶中没有数据
			return
		}

		if bucket.timestamp > t.Unix() { //延迟时间未到
			return
		}

		// 延迟时间小于等于当前时间, 取出Job元信息并放入ready queue
		job, err := dq.getJob(bucket.jobID)
		if err != nil {
			log.Error("scanBucket getJob jobID(%v) bucketName(%v) err(%v)", bucket.jobID, bucketName, err)
			continue
		}

		if job == nil { // job元信息不存在, 从bucket中删除
			_, err = dq.removeFromBucket(bucketName, bucket.jobID)
			if err != nil {
				log.Error("scanBucket removeFromBucket jobID(%v) bucketName(%v) err(%v)", bucket.jobID, bucketName, err)
			}
			continue
		}

		if job.TimeStamp > t.Unix() { // 再次确认元信息中delay是否小于等于当前时间
			_, err = dq.removeFromBucket(bucketName, bucket.jobID)
			if err != nil {
				log.Error("scanBucket removeFromBucket jobID(%v) bucketName(%v) err(%v)", bucket.jobID, bucketName, err)
			}

			// 重新计算delay时间并放入bucket中
			job.TimeStamp = job.Delay + time.Now().Unix()
			log.Info("job-%v 重新计算delay时间并放入bucket-%s中, 过期时间是：%s, 已到达执行时间：%s",
				job.ID, bucketName,
				time.Unix(job.TimeStamp, 0).Format("2006-01-02 15:04:05"),
				time.Now().Format("2006-01-02 15:04:05"))

			_, err = dq.pushToBucket(bucketName, job.TimeStamp, bucket.jobID)
			continue
		}

		log.Info("scanBucket pushToReadyQueue 从桶中取出:%s, job-%v: 准备进入就绪队列:%s, 过期时间是：%s, 已到达执行时间：%s",
			bucketName,
			bucket.jobID,
			job.Topic,
			time.Unix(bucket.timestamp, 0).Format("2006-01-02 15:04:05"),
			time.Unix(t.Unix(), 0).Format("2006-01-02 15:04:05"))

		_, err = dq.pushToReadyQueue(job.Topic, bucket.jobID)
		if err != nil {
			log.Error("scanBucket pushToReadyQueue job(%+v) bucketName(%v) err(%v)", job, bucketName, err)
			continue
		}

		dq.mux.Lock()
		dq.ReadyCount++
		dq.DelayCount--
		dq.mux.Unlock()

		_, err = dq.removeFromBucket(bucketName, bucket.jobID) // 从bucket中删除旧的jobId
		if err != nil {
			log.Error("scanBucket removeFromBucket jobID(%v) bucketName(%v) err(%v)", bucket.jobID, bucketName, err)
			continue
		}

		log.Info("scanBucket dq.removeFromBucket jobID(%v) bucketName(%v)", bucket.jobID, bucketName)
	}
}

func (dq *DQ) monitoring() { //监控延时队列运行任务状况
	go func() {
		for {
			dq.mux.RLock()
			log.Info("monitoring 当前时间:%v 不可执行状态:%+v "+
				"可执行状态:%+v "+
				"已被消费者读取未得到消费者响应:%+v "+
				"已被消费完成或者已被删除:%+v ",
				time.Now().Format("2006-01-02 15:04:05"),
				dq.DelayCount,
				dq.ReadyCount,
				dq.ReservedCount,
				dq.DeletedCount,
			)
			dq.mux.RUnlock()
			time.Sleep(time.Second)

			select {
			case <-dq.exit:
				log.Warn("monitoring exit")
				return
			default:
			}
		}
	}()
}

func (dq *DQ) getBucketName(i int) string {
	return fmt.Sprintf("%s%d", _bucketPrefix, i)
}

func (dq *DQ) getBucketNameByJobID(jobID int64) string {
	return fmt.Sprintf("%s%d", _bucketPrefix, dq.shardingBucketIndex(jobID))
}

func (dq *DQ) shardingBucketIndex(jobID int64) int64 {
	//return crc32.ChecksumIEEE([]byte(jobID)) % uint32(dq.bucketCount)
	return jobID % int64(dq.bucketCount)

}

//消息存储
//在选择存储介质之前，先来确定下具体的数据结构：
//
//Job Poll存放的Job元信息，只需要K/V形式的结构即可。key为job id，value为job struct。
//Delay Bucket是一个有序队列。
//Ready Queue是一个普通list或者队列都行。
//能够同时满足以上需求的，非redis莫属了。
//bucket的数据结构就是redis的zset，将其分为多个bucket是为了提高扫描速度，降低消息延迟。

//举例说明一个Job的生命周期
//用户对某个商品下单，系统创建订单成功，同时往延迟队列里put一个job。
//job结构为：{'topic':'orderclose', 'id':'ordercloseorderNoXXX', 'delay':1800 ,'TTR':60 , 'body':'XXXXXXX'}
func (dq *DQ) EnQueue(job Job) (err error) {
	//延迟队列收到该job后，先往job redisPool中存入job信息，
	//然后根据delay计算出绝对执行时间，并以轮询(round-robbin)的方式将job id放入某个bucket。
	job.TimeStamp = job.Delay + time.Now().Unix()

	jobByte, err := json.Marshal(job)
	if err != nil {
		return
	}

	jobStr := string(jobByte)
	jobIDStr := strconv.FormatInt(job.ID, 10)

	_, err = dq.addJob(jobIDStr, jobStr)
	if err != nil {
		log.Error("EnQueue dq.addJob fail job(%+v) error(%v)", job, err)
		return err
	}

	bucket := dq.getBucketNameByJobID(job.ID)
	log.Info("EnQueue job-%v: 加入延时队列的桶中-%s, 当前时间：%s, 过期时间是：%v",
		job.ID,
		bucket,
		time.Now().Format("2006-01-02 15:04:05"),
		time.Unix(job.TimeStamp, 0).Format("2006-01-02 15:04:05"),
	)

	_, err = dq.pushToBucket(bucket, job.TimeStamp, jobIDStr)
	if err != nil {
		log.Error("EnQueue pushToBucket fail job(%+v) error(%v)", job, err)
		return err
	}

	dq.mux.Lock()
	dq.DelayCount++
	dq.mux.Unlock()
	return
}

func (dq *DQ) DeQueue(topic string) (job *Job, err error) {
	jobId, err := dq.popFromReadyQueue(topic)
	if err != nil {
		return
	}
	if jobId == "" { // 队列为空
		return
	}

	job, err = dq.getJob(jobId) // 获取job元信息
	if err != nil {
		return
	}
	if job == nil { // 消息不存在, 可能已被删除
		return
	}

	_, err = dq.delJob(strconv.FormatInt(job.ID, 10))
	if err != nil {
		log.Error("DeQueue dq.delJob jobID(%v) error(%v)", job.ID, err)
	}

	log.Info("DeQueue jobID(%v) delete success", job.ID)

	//timestamp := time.Now().Unix() + job.TTR
	//_, err = pushToBucket(getBucketName(job.ID, bucketCount), timestamp, job.ID)
	dq.mux.Lock()
	dq.DeletedCount++
	dq.ReadyCount--
	dq.mux.Unlock()
	return
}

// 添加JobId到bucket中
func (dq *DQ) pushToBucket(key string, timestamp int64, jobId string) (res int, err error) {
	conn := dq.redisPool.Get()
	defer func() {
		err = conn.Close()
		return
	}()

	if err = conn.Send("ZADD", key, timestamp, jobId); err != nil {
		log.Error("pushToBucket ZADD key(%v) timestamp(%d) jobId(%v) error(%v)", key, timestamp, jobId, err)
		return
	}

	if err = conn.Flush(); err != nil {
		log.Error("pushToBucket Flush key(%v) timestamp(%d) jobId(%v) error(%v)", key, timestamp, jobId, err)
		return
	}

	if res, err = redis.Int(conn.Receive()); err != nil {
		log.Error("pushToBucket Receive key(%v) timestamp(%d) jobId(%v) error(%v)", key, timestamp, jobId, err)
	}

	return
}

func (dq *DQ) getBucket(key string) (res *Bucket, err error) { // 从bucket中获取延迟时间最小的jobID
	conn := dq.redisPool.Get()
	defer func() {
		err = conn.Close()
		return
	}()

	if err = conn.Send("ZRANGE", key, 0, 0, "WITHSCORES"); err != nil {
		log.Error("getBucket ZREM key(%v) error(%v)", key, err)
		return
	}

	if err = conn.Flush(); err != nil {
		log.Error("getBucket Flush key(%v) error(%v)", key, err)
		return
	}

	var resp []string

	if resp, err = redis.Strings(conn.Receive()); err != nil {
		log.Error("getBucket Receive key(%v) error(%v)", key, err)
		return
	}

	if len(resp) == 0 {
		//log.Info("getBucket key(%v) no resp", key)
		return
	}

	res = &Bucket{}
	res.jobID = resp[0]

	timeStampStr := resp[1]
	res.timestamp, _ = strconv.ParseInt(timeStampStr, 10, 64)
	return
}

func (dq *DQ) removeFromBucket(bucket string, jobID string) (res int, err error) {
	conn := dq.redisPool.Get()
	defer func() {
		err = conn.Close()
		return
	}()

	if err = conn.Send("ZREM", bucket, jobID); err != nil {
		log.Error("removeFromBucket ZREM bucket(%v) jobID(%v) error(%v)", bucket, jobID, err)
		return
	}

	if err = conn.Flush(); err != nil {
		log.Error("removeFromBucket Flush bucket(%v) jobID(%v) error(%v)", bucket, jobID, err)
		return
	}

	if res, err = redis.Int(conn.Receive()); err != nil {
		log.Error("removeFromBucket Receive bucket(%v) jobID(%v) error(%v)", bucket, jobID, err)
	}

	return
}

func (dq *DQ) pushToReadyQueue(topic string, jobID string) (res int, err error) {
	conn := dq.redisPool.Get()
	defer func() {
		err = conn.Close()
		return
	}()

	if err = conn.Send("RPUSH", topic, jobID); err != nil {
		log.Error("pushToReadyQueue LPUSH topic(%v) jobID(%v) error(%v)", topic, jobID, err)
		return
	}

	//if err = conn.Send("EXPIRE", queueName, 5); err != nil {
	//	log.Error("TaskLPUSH EXPIRE queueName(%d) jobID(%v) error(%v)", queueName, jobID, err)
	//	return
	//}

	if err = conn.Flush(); err != nil {
		log.Error("pushToReadyQueue Flush topic(%v) jobID(%v) error(%v)", topic, jobID, err)
		return
	}

	if res, err = redis.Int(conn.Receive()); err != nil {
		log.Error("pushToReadyQueue Receive topic(%v) jobID(%v) error(%v)", topic, jobID, err)
	}

	return
}

func (dq *DQ) popFromReadyQueue(topic string) (res string, err error) {
	conn := dq.redisPool.Get()
	defer func() {
		err = conn.Close()
		return
	}()

	if res, err = redis.String(conn.Do("LPOP", topic)); err != nil { //TODO BLPOP Error: command not support
		log.Error("popFromReadyQueue Receive topic(%+v) error(%v)", topic, err)
		if err == redis.ErrNil {
			err = nil
		}
		return
	}

	return
}

func (dq *DQ) addJob(key, val string) (res int, err error) {
	conn := dq.redisPool.Get()
	defer func() {
		err = conn.Close()
		return
	}()

	if err = conn.Send("SET", key, val); err != nil {
		log.Error("addJob SET key(%v) error(%v)", key, err)
		return
	}

	if err = conn.Flush(); err != nil {
		log.Error("addJob Flush key(%v) error(%v)", key, err)
		return
	}

	var resp string
	if resp, err = redis.String(conn.Receive()); err != nil {
		log.Error("addJob Receive key(%v) error(%v)", key, err)
	}

	if resp == "OK" {
		res = 1
	}

	return
}

func (dq *DQ) getJob(key string) (res *Job, err error) {
	conn := dq.redisPool.Get()
	defer func() {
		err = conn.Close()
		return
	}()

	if err = conn.Send("GET", key); err != nil {
		log.Error("getJob GET key(%v) error(%v)", key, err)
		return
	}

	if err = conn.Flush(); err != nil {
		log.Error("getJob Flush key(%v) error(%v)", key, err)
		return
	}

	var resp []byte
	if resp, err = redis.Bytes(conn.Receive()); err != nil {
		log.Error("getJob Receive key(%v) error(%v)", key, err)
		return
	}

	if len(resp) == 0 {
		log.Error("getJob Receive key(%v) no data", key)
		return
	}

	res = &Job{}
	err = json.Unmarshal(resp, res)
	if err != nil {
		log.Error("getJob json.Unmarshal key(%v) error(%v)", key, err)
	}

	return
}

func (dq *DQ) delJob(key string) (res int, err error) {
	conn := dq.redisPool.Get()
	defer func() {
		err = conn.Close()
		return
	}()

	if err = conn.Send("DEL", key); err != nil {
		log.Error("delJob GET key(%v) error(%v)", key, err)
		return
	}

	if err = conn.Flush(); err != nil {
		log.Error("delJob DEL key(%v) error(%v)", key, err)
		return
	}

	if res, err = redis.Int(conn.Receive()); err != nil {
		log.Error("delJob Receive key(%v) error(%v)", key, err)
	}

	return
}
