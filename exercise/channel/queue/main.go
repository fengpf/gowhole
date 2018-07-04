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

type service struct {
	wg                sync.WaitGroup
	soureData, myData chan int
	de                DispatchEngine
}

var (
	dataLen     = 100
	workerCount = 10
	workerQueue []chan int
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
		soureData: make(chan int, dataLen),
		myData:    make(chan int),
	}
	for i := 0; i < dataLen; i++ {
		s.soureData <- i
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
	for {
		i, ok := <-s.soureData
		if !ok {
			fmt.Printf("consume exit\n")
			return
		}
		s.myData <- i
	}
}

func (s *service) Close() {
	close(s.soureData)
	close(s.myData)
	s.de.Scheduler.Close()
	s.wg.Wait()
}

type Scheduler interface {
	InitQueue([]chan int, *service)
	SubmitMsg(int)
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
	msgChan                      chan int
	workerQueue, remainQueue     []chan int
	lenOfQueue, remainLenOfQueue int
	quit, quit2                  chan struct{}
}

func (d *DataScheduler) InitQueue(q []chan int, s *service) {
	d.s = s
	d.msgChan = make(chan int)
	d.lenOfQueue = len(q)
	d.workerQueue = q
	d.quit = make(chan struct{})
	d.quit2 = make(chan struct{})
}

func (d *DataScheduler) SubmitMsg(m int) {
	d.msgChan <- m
}

func (d *DataScheduler) Dispatch() { //队列分发
	defer d.s.wg.Done()
	d.s.wg.Add(1)
	go func() {
		for {
			select {
			case m, ok := <-d.msgChan:
				if !ok {
					fmt.Printf("close Dispatch msgChan \n")
					return
				}
				d.workerQueue[d.ShardingQueueIndex(m)] <- m
			case <-d.quit:
				fmt.Printf("Dispatch quit \n")
				return
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
				case n, ok := <-d.workerQueue[iWorker]:
					if !ok {
						fmt.Printf("close Start workerQueue(%d) chan \n", iWorker)
						return
					}
					fmt.Printf("worker num(%d) data(%d)\n", iWorker, n)
				case <-d.quit2:
					fmt.Printf("Start quit\n")
					return
				}
			}
		}(i)
	}
}

func (de *DispatchEngine) Run(s *service) {
	defer s.wg.Done()
	workerQueue := make([]chan int, de.WorkerCount)
	for i := 0; i < de.WorkerCount; i++ {
		workerQueue[i] = make(chan int)
	}
	de.Scheduler.InitQueue(workerQueue, s)
	de.Scheduler.Dispatch()
	de.Scheduler.Start()
	for m := range s.myData {
		de.Scheduler.SubmitMsg(m)
	}
	fmt.Printf("Run.....over\n")
}

func (d *DataScheduler) ShardingQueueIndex(m int) (i int) {
	i = m % d.lenOfQueue
	return
}

func (d *DataScheduler) Stop() {
	go func() {
		d.quit <- struct{}{}
		d.quit2 <- struct{}{}
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
		fmt.Printf("after close msgChan get m(%d)\n", m)
	}
	d.Stop()
}
