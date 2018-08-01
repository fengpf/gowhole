package schedule

import (
	"fmt"
)

type Dispatcher struct {
	// A pool of workers channels that are registered with the dispatcher
	WorkerPool chan chan Job
	MaxWorkers int
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{
		WorkerPool: pool,
		MaxWorkers: maxWorkers,
	}
}

func (d *Dispatcher) Run() {
	// starting n number of workers
	for i := 0; i < d.MaxWorkers; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}
	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	fmt.Println("Worker que dispatcher started...")
	for {
		select {
		case job := <-JobQueue:
			fmt.Println("get a job from JobQueue")
			// a job request has been received
			go func(job Job) {
				// try to obtain a worker job channel that is available.
				// this will block until a worker is idle
				jobChannel := <-d.WorkerPool
				fmt.Println("get a jobChannel  from WorkerPool")

				// dispatch the job to the worker job channel
				jobChannel <- job
				fmt.Println("dispatch job to jobChannel")
			}(job)
		}
	}
}
