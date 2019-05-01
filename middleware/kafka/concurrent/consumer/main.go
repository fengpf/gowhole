package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"time"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	diskqueue "github.com/nsqio/go-diskqueue"
)

var (
	err error

	//queue
	chanSize    = 1024
	queueLen    = 1000
	workerQueue []chan []byte

	//kafak
	config   *cluster.Config
	consumer *cluster.Consumer
	brokers  = []string{"10.23.39.129:9092", "10.23.39.129:9093", "10.23.39.129:9094"}
	topics   = []string{"topic001"}
	groupID  = "group-1"

	signals = make(chan os.Signal, 1)

	//持久化
	dq BackendQueue
)

func main() {
	if consumer == nil {
		log.Fatalln("kafka cluster consumer instance is nil")
		return
	}
	go stasticGroutine()

	dqLogf := func(lvl diskqueue.LogLevel, f string, args ...interface{}) {
		log.Println((fmt.Sprintf(lvl.String()+": "+f, args...)))
	}

	dqName := "test_disk_queue" + strconv.Itoa(int(time.Now().Unix()))
	tmpDir, err := ioutil.TempDir("", fmt.Sprintf("test-%d", time.Now().UnixNano()))
	if err != nil {
		panic(err)
	}
	dq = diskqueue.New(dqName, tmpDir, 1024, 4, 1<<10, 2500, 2*time.Second, dqLogf)
	// msgOut := <-dq.ReadChan()

	defer func() {
		for i := 0; i < queueLen; i++ { //关闭队列中的所有chan
			close(workerQueue[i])
			// if v, ok := <-workerQueue[i]; ok && len(v) > 0 {
			// dq.Put(v)//需要实现对应的存储接口
			// }
		}

		consumer.Close()
		// os.RemoveAll(tmpDir)
		dq.Close()
	}()

	//定义chan队列并启动消费
	workerQueue = make([]chan []byte, queueLen)
	for i := 0; i < queueLen; i++ {
		ch := make(chan []byte, chanSize)
		workerQueue[i] = ch
		go func(m chan []byte) { //播放
			for v := range m {
				//todo 业务逻辑，耗时用sleep代替
				println(string(v))
				time.Sleep(time.Millisecond * 50)
			}
		}(ch)
	}

	signal.Notify(signals, os.Interrupt)
	consume() //消费kf
}

func consume() {
	go func() { // consume errors
		for err := range consumer.Errors() {
			log.Printf("consumer-%s:Error: %s\n", groupID, err.Error())
		}
	}()
	go func() { // consume notifications
		for ntf := range consumer.Notifications() {
			log.Printf("consumer-%s:Rebalanced: %+v \n", groupID, ntf)
		}
	}()

	for {
		select {
		case msg, ok := <-consumer.Messages():
			if !ok {
				log.Fatalf("consumer.Messages: %+v \n", ok)
				return
			}
			fmt.Fprintf(os.Stdout, "GroupID-(%s):Topic(%s)\tPartition(%d)\tOffset(%d)\tKey(%s)\tValue(%s)\n",
				groupID, msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value)) //打印消费日志

			workerQueue[bytesToInt(msg.Key)%queueLen] <- msg.Value //消费到数据放到chan队列

			consumer.MarkOffset(msg, "") // mark message as processed
		case <-signals:
			fmt.Fprintf(os.Stdout, "exit kafka consume...")
			return
		}
	}
}

func init() {
	config = cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	// init consumer
	c, err := cluster.NewConsumer(brokers, groupID, topics, config)
	if err != nil {
		panic(err)

	}
	if c == nil {
		log.Printf("cluster.NewConsumer instance is nil")
		return
	}
	consumer = c

	// consume errors
	go func() {
		for err := range consumer.Errors() {
			log.Printf("consumer%s:Error: %s\n", groupID, err.Error())
		}
	}()
	// consume notifications
	go func() {
		for ntf := range consumer.Notifications() {
			log.Printf("consumer%s:Rebalanced: %+v \n", groupID, ntf)
		}
	}()
}

func bytesToInt(b []byte) int { //字节转换成整形
	bytesBuffer := bytes.NewBuffer(b)
	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return int(x)
}

func stasticGroutine() {
	for {
		time.Sleep(time.Second)
		total := runtime.NumGoroutine()
		fmt.Println("NumGoroutine:", total)
	}
}

//BackendQueue for save disk
type BackendQueue interface {
	Put([]byte) error
	ReadChan() chan []byte // this is expected to be an *unbuffered* channel
	Close() error
	Delete() error
	Depth() int64
	Empty() error
}

//参考：
// https://www.jianshu.com/p/1a746f57cdd6
// https://github.com/nsqio/go-diskqueue
