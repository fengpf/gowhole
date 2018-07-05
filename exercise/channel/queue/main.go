package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"runtime/pprof"
	"sync"
	"syscall"
)

type stu struct {
	id int
}

type msg struct {
	next *msg
	data *stu
	done bool
}

type service struct {
	wg        sync.WaitGroup
	soureData chan *stu
	myData    chan *msg
	mergeChan []chan *msg
	doneChan  chan []*msg
	de        DispatchEngine
}

var (
	chanSize    = 1024
	dataLen     = 1000
	workerCount = 10
	workerQueue []chan *msg
	s           *service
	cpuprofile  = flag.String("cpuprofile", "", "write cpu profile to `file`")
	memprofile  = flag.String("memprofile", "", "write memory profile to `file`")
)

func pprofStart() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}
	// ... rest of the program ...
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}
}
func init() {
	pprofStart()
	s = &service{
		de: DispatchEngine{
			Scheduler:   &DataScheduler{},
			WorkerCount: workerCount,
		},
		soureData: make(chan *stu, chanSize),
		myData:    make(chan *msg, chanSize),
		mergeChan: make([]chan *msg, workerCount),
		doneChan:  make(chan []*msg, chanSize),
	}
	for i := 0; i < dataLen; i++ {
		s.soureData <- &stu{id: i}
	}
}

func main() {
	s.wg.Add(1)
	go s.consume()
	s.wg.Add(1)
	go s.de.Run(s)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT, syscall.SIGHUP)
	for {
		sg := <-c
		fmt.Printf("main exit by signal(%s)\n", sg.String())
		switch sg {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT:
			s.Close()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

func (s *service) consume() {
	defer s.wg.Done()
	var head, last *msg
	for {
		select {
		case stu, ok := <-s.soureData:
			if !ok {
				return
			}
			m := &msg{data: stu}
			if head == nil {
				head = m
				last = m
			} else {
				last.next = m
				last = m
			}
			s.myData <- m
		case done := <-s.doneChan:
			fin := make(map[int]*stu)
			for _, d := range done {
				d.done = true
			}
			for ; head != nil && head.done; head = head.next {
				fin[head.data.id] = head.data
			}
			for _, f := range fin {
				fmt.Printf("consume at id(%d)\n", f.id)
			}
		}
	}
}

func (s *service) Close() {
	defer s.wg.Wait()
	close(s.soureData)
	close(s.myData)
	s.de.Scheduler.Close()
}

type Scheduler interface {
	InitQueue([]chan *msg, *service)
	SubmitMsg(*msg)
	Dispatch()
	Start()
	Close()
}

type DispatchEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type DataScheduler struct {
	s                            *service
	msgChan                      chan *msg
	workerQueue, remainQueue     []chan *msg
	lenOfQueue, remainLenOfQueue int
	stopCh                       chan struct{}
	toStop                       chan struct{}
}

func (d *DataScheduler) InitQueue(q []chan *msg, s *service) {
	d.s = s
	d.msgChan = make(chan *msg)
	d.lenOfQueue = len(q)
	d.workerQueue = q
	d.stopCh = make(chan struct{})
	d.toStop = make(chan struct{})
}

func (d *DataScheduler) SubmitMsg(m *msg) {
	d.msgChan <- m
}

func (d *DataScheduler) Dispatch() { //队列分发
	defer d.s.wg.Done()
	d.s.wg.Add(1)
	go func() {
		for {
			select {
			case <-d.stopCh:
				fmt.Printf("Dispatch stopCh \n")
			default:
			}
			select {
			case m, ok := <-d.msgChan:
				if !ok {
					select {
					case d.toStop <- struct{}{}:
						fmt.Printf("Dispatch toStop ...\n")
					default:
					}
					return
				}
				d.workerQueue[d.ShardingQueueIndex(m)] <- m
			}
		}
	}()
}

func (d *DataScheduler) Start() { //队列消费
	for i := 0; i < d.lenOfQueue; i++ {
		d.s.wg.Add(1)
		go func(iWorker int) {
			defer d.s.wg.Done()
			for {
				select {
				case <-d.stopCh:
					fmt.Printf("Start stopCh\n")
				default:
				}
				select {
				case n, ok := <-d.workerQueue[iWorker]:
					if !ok {
						select {
						case d.toStop <- struct{}{}:
						default:
						}
						fmt.Printf("Start toStop at num(%d)\n", iWorker)
						return
					}
					fmt.Printf("worker Start at num(%d) data(%d)\n", iWorker, n.data.id)
				}
			}
		}(i)
	}
}

func (de *DispatchEngine) Run(s *service) {
	defer s.wg.Done()
	workerQueue := make([]chan *msg, de.WorkerCount)
	for i := 0; i < de.WorkerCount; i++ {
		m := make(chan *msg, chanSize)
		// s.mergeChan[i] = m
		workerQueue[i] = m
	}
	de.Scheduler.InitQueue(workerQueue, s)
	de.Scheduler.Dispatch()
	de.Scheduler.Start()
	for m := range s.myData {
		de.Scheduler.SubmitMsg(m)
	}
	fmt.Printf("Run.....over\n")
}

func (d *DataScheduler) ShardingQueueIndex(m *msg) (i int) {
	i = m.data.id % d.lenOfQueue
	return
}

func (d *DataScheduler) Stop() {
	go func() {
		<-d.toStop
		close(d.stopCh)
	}()
}

func (d *DataScheduler) Close() {
	if d.msgChan != nil {
		close(d.msgChan)
	}
	for i := 0; i < d.lenOfQueue; i++ {
		close(d.workerQueue[i])
	}
	for m := range d.msgChan {
		fmt.Printf("after close msgChan get m(%d)\n", m.data.id)
	}
	d.Stop()
}
