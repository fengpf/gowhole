package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/Shopify/sarama"
)

var (
	err error

	brokers  = []string{"192.168.36.194:9092", "192.168.36.194:9093", "192.168.36.194:9094"}
	topic    = "topic001"
	producer sarama.SyncProducer

	mockNum = 10 //模拟每秒并发数量
)

func main() {
	if producer == nil {
		log.Fatalln("kafka producer instance is nil")
		return
	}

	defer producer.Close()
	go stasticGroutine()

	var wg sync.WaitGroup
	for {
		wg.Add(mockNum)
		for i := 0; i < mockNum; i++ {
			go syncProducer(i, &wg)
		}
		wg.Wait()
		time.Sleep(time.Second)
	}
}

func syncProducer(i int, wg *sync.WaitGroup) {
	defer wg.Done()

	value := fmt.Sprintf("sync: this is a kafka message. index=%d", i)
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.ByteEncoder(strconv.Itoa(i)),
		Value: sarama.ByteEncoder(value),
	}

	part, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Printf("send message(%s) err=%s \n", value, err)
	} else {
		fmt.Fprintf(os.Stdout, value+"发送成功，partition=%d, offset=%d \n", part, offset)
	}
}

func init() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll //等待服务器所有副本都保存成功后的响应
	config.Producer.Return.Successes = true          //是否等待成功和失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用.
	config.Producer.Return.Errors = true
	config.Producer.Timeout = 5 * time.Second
	producer, err = sarama.NewSyncProducer(brokers, config)
	if err != nil {
		panic(err)
	}
}

func stasticGroutine() {
	for {
		time.Sleep(time.Second)
		total := runtime.NumGoroutine()
		fmt.Println("NumGoroutine:", total)
	}
}
