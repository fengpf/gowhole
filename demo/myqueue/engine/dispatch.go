package engine

import (
	"sync"
)

//DispatchEngine for dispatch engine.
type DispatchEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	Wg          sync.WaitGroup
}

//Scheduler for scheduler interface.
type Scheduler interface {
	InitQueue([]chan int, *sync.WaitGroup)
	SubmitMsg(int)
	Dispatch()
	Start()
	Close()
}

//Run run an engine.
func (de *DispatchEngine) Run(msg []int) {
	defer de.Wg.Done()
	workerQueue := make([]chan int, de.WorkerCount)
	for i := 0; i < de.WorkerCount; i++ {
		workerQueue[i] = make(chan int)
	}
	de.Scheduler.InitQueue(workerQueue, &de.Wg)
	de.Wg.Add(1)
	de.Scheduler.Dispatch()
	de.Wg.Add(1)
	de.Scheduler.Start()
	for _, m := range msg {
		if m > 0 {
			de.Scheduler.SubmitMsg(m)
		}
	}
}

//Stop stop an engine.
func (de *DispatchEngine) Stop() {
	de.Scheduler.Close()
}
