package main

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

var pv PoolVar

type PoolVar struct {
	maxWorker  int
	queueLen   int
	wg         sync.WaitGroup
	jobQueue   chan Job
	workerPool chan chan Job
	limiter    chan bool //限制goroutine个数
}

type ActionListener interface {
	DoAction() (err error)
}

// Job represents the job to be run
type Job struct {
	ActionListener
}

// Worker represents the worker that executes the job
type Worker struct {
	Id         int
	JobChannel chan Job
	quit       chan bool
}

func NewWorker(i int) Worker {
	return Worker{
		Id:         i,
		JobChannel: make(chan Job),
		quit:       make(chan bool),
	}
}

// Start method starts the run loop for the worker, listening for a quit channel in
// case we need to stop it
func (w Worker) Start() {
	for {
		// register the current worker into the worker queue.
		//w.WorkerPool <- w.JobChannel
		pv.workerPool <- w.JobChannel
		select {
		case job := <-w.JobChannel:
			pv.limiter <- true
			fmt.Printf("======worker:%d, get job:%d======\n", w.Id, job.ActionListener)
			// we have received a work request.
			go func() {
				job.DoAction()
				<-pv.limiter
				pv.wg.Done()
			}()
		case <-w.quit:
			// we have received a signal to stop
			return
		}
	}
}

// Stop signals the worker to stop listening for work requests.
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

func AddJob(l ActionListener) {
	pv.wg.Add(1)
	work := Job{ActionListener: l}
	pv.jobQueue <- work
}

func WaitDone() {
	pv.wg.Wait()
	close(pv.jobQueue)
	close(pv.workerPool)
}

func init() {
	pv = PoolVar{
		wg:        sync.WaitGroup{},
		maxWorker: runtime.NumCPU(),
		queueLen:  10,
		limiter:   make(chan bool, 10),
	}
	pv.jobQueue = make(chan Job, pv.queueLen)
	pv.workerPool = make(chan chan Job, pv.maxWorker)

	for i := 1; i <= pv.maxWorker; i++ {
		worker := NewWorker(i)
		go worker.Start()
	}
	go dispatch()

}

func dispatch() {
	for job := range pv.jobQueue {
		// try to obtain a worker job channel that is available.
		// this will block until a worker is idle
		// then dispatch the job to the worker job channel
		<-pv.workerPool <- job
	}
}

//eg:

type A struct {
	id int
}

func (a *A) DoAction() (err error) {
	time.Sleep(time.Second)
	return
}

func Test_work(t *testing.T) {
	for i := 1; i <= 100; i++ {
		a := &A{i}
		AddJob(a)
	}
	WaitDone()
}
