package scheduler

import (
	"fmt"
	"hash/crc32"
	"strconv"
)

type DataScheduler struct {
	msgChan       chan int
	workerQueue   []chan int
	lengthOfQueue int
}

func (ds *DataScheduler) SetQueue(q []chan int) {
	ds.msgChan = make(chan int)
	ds.lengthOfQueue = len(q)
	ds.workerQueue = q
}

func (ds *DataScheduler) SubmitMsg(m int) {
	ds.msgChan <- m
}

func (ds *DataScheduler) Dispatch() {
	go func() {
		for {
			m, ok := <-ds.msgChan
			if !ok {
				fmt.Println("dispatch wait")
				return
			}
			mi := strconv.Itoa(m)
			ch := crc32.ChecksumIEEE([]byte(mi))
			index := int(ch) % ds.lengthOfQueue
			ds.workerQueue[index] <- m
		}
	}()
}

func (ds *DataScheduler) Start() {
	for i := 0; i < ds.lengthOfQueue; i++ {
		go func(i int) {
			for {
				m, ok := <-ds.workerQueue[i]
				if !ok {
					fmt.Printf("does not exist the queue(%d)\n", i)
				}
				fmt.Printf("queue(%d) msg(%d)\n", i, m)
			}
		}(i)
	}
}
