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
	queueLen    = 10
	workerQueue []chan []byte

	//kafak
	config   *cluster.Config
	consumer *cluster.Consumer
	brokers  = []string{"192.168.36.194:9092", "192.168.36.194:9093", "192.168.36.194:9094"}
	topics   = []string{"topic001"}
	groupID  = "group-1"

	signals = make(chan os.Signal, 1)

	//持久化
	dq diskqueue.Interface
)

func main() {
	if consumer == nil {
		log.Fatalln("kafka cluster consumer instance is nil")
		return
	}
	go stasticGroutine()

	dqName := "mychan" + strconv.Itoa(int(time.Now().Unix()))
	tmpDir, err := ioutil.TempDir("/data/app/go/src/gowhole/middleware/kafka/concurrent/consumer", fmt.Sprintf("test-%d", time.Now().Unix()))
	if err != nil {
		panic(err)
	}
	dq = diskqueue.New(dqName, tmpDir, 1024, 4, 1<<10, 2500, 2*time.Second,
		func(lvl diskqueue.LogLevel, f string, args ...interface{}) {
			log.Println((fmt.Sprintf(lvl.String()+": "+f, args...)))
		})

	defer func() {
		consumer.Close()
		for i := 0; i < queueLen; i++ { //关闭队列中的所有chan
			close(workerQueue[i]) //先关闭chan，防止关闭时清空阻塞chan阻塞

			if v, ok := <-workerQueue[i]; ok && len(v) > 0 {
				log.Println("chan关闭,保存数据", string(v))
				dq.Put(v)
			}
		}
		dq.Close()

		os.RemoveAll(tmpDir)
	}()

	//定义chan队列并启动消费
	workerQueue = make([]chan []byte, queueLen)
	for i := 0; i < queueLen; i++ {
		ch := make(chan []byte, chanSize)
		workerQueue[i] = ch
		go func(m chan []byte) {
			for v := range m {
				//todo 业务逻辑，耗时用sleep代替
				log.Println("业务逻辑处理", string(v))
				time.Sleep(time.Second * 1)
			}
		}(ch)
	}

	signal.Notify(signals, os.Interrupt)
	go func() {
		for {
			select {
			case v, ok := <-dq.ReadChan():
				if !ok {
					log.Fatalf("dq.ReadChan: %+v \n", ok)
					return
				}
				log.Println("读取本地磁盘数据", string(v))
			case <-signals:
				fmt.Fprintf(os.Stdout, "退出读取本地磁盘协程")
				return
			}
		}
	}()

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

			log.Println("commit offest:", msg.Offset)
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
		log.Println("NumGoroutine:", total)
	}
}

//参考：
// https://www.jianshu.com/p/1a746f57cdd6
// https://github.com/nsqio/go-diskqueue
