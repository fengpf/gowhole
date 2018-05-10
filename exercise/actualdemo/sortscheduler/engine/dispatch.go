package engine

type DispatchEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	SetQueue([]chan int)
	Dispatch()
	Start()
	SubmitMsg(int)
}

func (de *DispatchEngine) Run(msg []int) {
	workerQueue := make([]chan int, de.WorkerCount)
	for i := 0; i < de.WorkerCount; i++ {
		workerQueue[i] = make(chan int)
	}
	de.Scheduler.SetQueue(workerQueue)
	de.Scheduler.Dispatch()
	de.Scheduler.Start()
	for m := range msg {
		de.Scheduler.SubmitMsg(m)
	}
}
