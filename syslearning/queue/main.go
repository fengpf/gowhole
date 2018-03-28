package main

import (
	"fmt"
	"gowhole/syslearning/queue/schedule"
	"time"
)

var (
	MaxWorker = 2
	MaxQueue  = 2
)

func test() {
	for i := 0; i < 1000; i++ {
		fmt.Println(i)
	}

}

func main() {
	schedule.JobQueue = make(chan schedule.Job, MaxQueue)

	dp := schedule.NewDispatcher(MaxWorker)
	dp.Run()

	for i := 0; i < 10; i++ {
		work := schedule.Job{Consume: test}
		schedule.JobQueue <- work
		fmt.Println("sending payload  to workque")
	}

	time.Sleep(time.Minute)
}
