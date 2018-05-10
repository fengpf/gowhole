package main

import (
	"fmt"
	"gowhole/exercise/actualdemo/sortscheduler/engine"
	"gowhole/exercise/actualdemo/sortscheduler/scheduler"
	"time"
)

func main() {
	//mock 消费数据
	count := 100
	msg := make([]int, count)
	for i := 0; i < count; i++ {
		msg[i] = i
	}
	de := engine.DispatchEngine{
		Scheduler:   &scheduler.DataScheduler{},
		WorkerCount: 100,
	}

	t := time.Now()
	de.Run(msg)
	elapsed := time.Since(t)

	time.Sleep(time.Second * 1)
	fmt.Printf("app elapsed(%v)\n", elapsed)
}
