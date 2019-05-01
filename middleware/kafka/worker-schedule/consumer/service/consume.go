package service

import (
	"log"
	"strconv"
	"sync"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
)

var (
	err error
	//queue
	chanSize    = 1024
	dataLen     = 200
	workerCount = 10
	//kafak
	config   *cluster.Config
	consumer *cluster.Consumer
	brokers  = []string{"localhost:9092"}
	groupID  = "group-1"
	topics   = []string{"test"}
)

//Service for kafka consume service.
type Service struct {
	//queue
	wg          sync.WaitGroup
	myData      chan *chain
	workerQueue []chan *chain
	// doneChan    chan []*chain
	de DispatchEngine
	//kafka
	config   *cluster.Config
	consumer *cluster.Consumer
}

//New for a kafka service.
func New() (s *Service) {
	s = &Service{
		//queue
		de: DispatchEngine{
			Scheduler:   &DataScheduler{},
			WorkerCount: workerCount,
		},
		myData: make(chan *chain, chanSize),
		//kafka
		config: cluster.NewConfig(),
	}
	s.config.Consumer.Return.Errors = true
	s.config.Group.Return.Notifications = true
	s.config.Consumer.Offsets.Initial = sarama.OffsetNewest
	// init consumer
	c, err := cluster.NewConsumer(brokers, groupID, topics, config)
	if err != nil {
		log.Printf("%s: sarama.NewSyncProducer err, message=%s \n", groupID, err)
		return
	}
	s.consumer = c
	//run queue
	s.wg.Add(1)
	go s.sub()
	s.wg.Add(1)
	go s.de.Run(s)
	return s
}

func (s *Service) sub() {
	defer s.wg.Done()
	// consume errors
	go func() {
		for err := range s.consumer.Errors() {
			log.Printf("%s:Error: %s\n", groupID, err.Error())
		}
	}()
	// consume notifications
	go func() {
		for ntf := range s.consumer.Notifications() {
			log.Printf("%s:Rebalanced: %+v \n", groupID, ntf)
		}
	}()

	// consume messages, watch signals
	var successes int
	for {
		select {
		case msg, ok := <-s.consumer.Messages():
			if !ok {
				log.Fatalf("consumer.Messages: %+v \n", ok)
				return
			}

			i, _ := strconv.Atoi(string(msg.Value))
			m := &chain{data: &stu{id: i}}
			s.myData <- m

			s.consumer.MarkOffset(msg, "") // mark message as processed
			successes++
			// fmt.Printf("GroupID-(%s):Topic(%s)\tPartition(%d)\tOffset(%d)\tKey(%s)\tValue(%s)\n", groupID, msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
		}
	}
}

//Close for release resource.
func (s *Service) Close() {
	defer s.wg.Wait()
	s.consumer.Close()
	close(s.myData)
	s.de.Scheduler.Close()
}
