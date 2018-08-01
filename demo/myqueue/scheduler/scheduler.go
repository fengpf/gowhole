package scheduler

import (
	"fmt"
	"hash/crc32"
	"strconv"
	"sync"
)

//DataScheduler for data scheduler.
type DataScheduler struct {
	msgChan       chan int
	workerQueue   []chan int
	lengthOfQueue int
	quit          chan bool
	wg            *sync.WaitGroup
}

//InitQueue init queue.
func (ds *DataScheduler) InitQueue(q []chan int, wg *sync.WaitGroup) {
	ds.msgChan = make(chan int)
	ds.lengthOfQueue = len(q)
	ds.workerQueue = q
	ds.wg = wg
	ds.quit = make(chan bool)
}

//SubmitMsg submit msg to  job chan.
func (ds *DataScheduler) SubmitMsg(m int) {
	ds.msgChan <- m
}

//Dispatch dispatch job for worker.
func (ds *DataScheduler) Dispatch() {
	ds.wg.Done()
	ds.wg.Add(1)
	go func() {
		ds.wg.Done()
		for {
			select {
			case m, ok := <-ds.msgChan:
				if !ok {
					return
				}
				// index := m % ds.lengthOfQueue//直接取模，缺点是当m是队列长度的整数倍时，会生成相同的队列编号.

				mi := strconv.Itoa(m)
				ch := crc32.ChecksumIEEE([]byte(mi)) //校验和取模能够均匀散列队列编号
				index := int(ch) % ds.lengthOfQueue

				ds.workerQueue[index] <- m
				fmt.Printf("Dispatch----key(%d)----val(%d)---\n", index, m)
			case quit := <-ds.quit:
				if quit {
					return
				}
			}
		}
	}()
}

// Start a worker.
func (ds *DataScheduler) Start() {
	ds.wg.Done()
	var (
		l     sync.Mutex
		count = make(map[int]int)
	)
	for i := 0; i < ds.lengthOfQueue; i++ {
		ds.wg.Add(1)
		go func(i int) {
			ds.wg.Done()
			for {
				select {
				case m, ok := <-ds.workerQueue[i]:
					if !ok {
						return
					}
					if m < 0 {
						continue
					}
					fmt.Printf("Start----key(%d)----val(%d)---\n", i, m)
					l.Lock()
					count[i]++
					// fmt.Printf("Count----count(%+v)\n", count)
					l.Unlock()
				case quit := <-ds.quit:
					if quit {
						return
					}
				}
			}
		}(i)
	}
}

//Stop exit for cycle.
func (ds *DataScheduler) Stop() {
	go func() {
		ds.quit <- true
	}()
}

//Close close chan.
func (ds *DataScheduler) Close() {
	close(ds.msgChan)
	for i := 0; i < ds.lengthOfQueue; i++ {
		close(ds.workerQueue[i])
	}
	ds.Stop()
}
