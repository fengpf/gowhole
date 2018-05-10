package main

import (
	"fmt"
	"gowhole/exercise/actualdemo/demoscheduler/model"
	"gowhole/exercise/actualdemo/demoscheduler/scheduler"
	"time"
)

//Scheduler for Scheduler interface
type Scheduler interface {
	WorkerChan() chan *model.Msg
	Submit(*model.Msg)
	Run()
	ReadyNotifier
}

//ReadyNotifier for worker ready interface
type ReadyNotifier interface {
	WorkReady(chan *model.Msg)
}

//ConcurrentEngine for concurrent engine
type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

//Run for engine run.
func (ce *ConcurrentEngine) Run(contens []int32) {
	ce.Scheduler.Run()
	for i := 0; i < ce.WorkerCount; i++ {
		ce.createWorker(ce.Scheduler.WorkerChan(), ce.Scheduler)
	}
	for _, c := range contens {
		ce.Scheduler.Submit(&model.Msg{
			ID: c,
		})
	}
}

func (ce *ConcurrentEngine) createWorker(in chan *model.Msg, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkReady(in) //set a worker
			fmt.Println("ret-----------", in)
			m := <-in //block to get msg
			echo(m.ID)
			fmt.Println("msg-----------", m.ID)
			fmt.Println("get-----------", in)
		}
	}()
}

func echo(id int32) int32 {
	id = id + 1
	// fmt.Println(id)
	return id
}

func main() {
	ce := ConcurrentEngine{
		Scheduler:   &scheduler.DataScheduler{},
		WorkerCount: 1,
	}

	count := 10
	contents := make([]int32, 0, count)
	for index := 0; index < count; index++ {
		contents = append(contents, int32(index))
	}

	// fmt.Println(contents)

	t := time.Now()
	ce.Run(contents)
	elapsed := time.Since(t)

	time.Sleep(time.Second * 3)

	fmt.Printf("app elapsed(%v)\n", elapsed)

}
