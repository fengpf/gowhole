package service

import (
	"fmt"
)

type stu struct {
	id int
}

type chain struct {
	next *chain
	data *stu
	done bool
}

//Scheduler for schedule.
type Scheduler interface {
	InitQueue(*Service)
	SubmitMsg(*chain)
	Dispatch()
	Start()
	Close()
}

//DispatchEngine for dispatch engine.
type DispatchEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

//DataScheduler for queue schedule.
type DataScheduler struct {
	s                            *Service
	msgChan                      chan *chain
	workerQueue                  []chan *chain
	remainQueue                  []*chain
	lenOfQueue, remainLenOfQueue int
	stopCh                       chan struct{}
	toStop                       chan struct{}
}

//InitQueue for init queue.
func (d *DataScheduler) InitQueue(s *Service) {
	d.s = s
	d.msgChan = make(chan *chain)
	d.lenOfQueue = len(s.workerQueue)
	d.workerQueue = s.workerQueue
	d.stopCh = make(chan struct{})
	d.toStop = make(chan struct{})
}

//SubmitMsg for submit msg
func (d *DataScheduler) SubmitMsg(m *chain) {
	d.msgChan <- m
}

//Dispatch for dispatch
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

//Start for performance the queue.
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
					fmt.Println("Start default stopCh........")
				}
				select {
				case n, ok := <-d.workerQueue[iWorker]:
					if !ok {
						fmt.Printf("Start toStop at num(%d)\n", iWorker)
						select {
						case d.toStop <- struct{}{}:
							fmt.Println("Start toStop........")
						default:
							fmt.Println("Start toStop default........")
						}
						return
					}
					fmt.Printf("worker Start at num(%d) data(%d)\n", iWorker, n.data.id)
				}
			}
		}(i)
	}
}

//Run for performance the schdule.
func (de *DispatchEngine) Run(s *Service) {
	defer s.wg.Done()
	s.workerQueue = make([]chan *chain, de.WorkerCount)
	for i := 0; i < de.WorkerCount; i++ {
		s.workerQueue[i] = make(chan *chain, chanSize)
	}
	de.Scheduler.InitQueue(s)
	de.Scheduler.Dispatch()
	de.Scheduler.Start()
	for m := range s.myData {
		de.Scheduler.SubmitMsg(m)
	}
	fmt.Printf("Run.....over\n")
}

//ShardingQueueIndex for splitting.
func (d *DataScheduler) ShardingQueueIndex(m *chain) (i int) {
	i = m.data.id % d.lenOfQueue
	return
}

//Stop for stop goroutine
func (d *DataScheduler) Stop() {
	go func() {
		<-d.toStop
		close(d.stopCh)
	}()
}

//Close for close chan
func (d *DataScheduler) Close() {
	d.Stop()
	close(d.msgChan)
	for i := 0; i < d.lenOfQueue; i++ {
		close(d.workerQueue[i])
	}
}
