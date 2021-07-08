package engine

import (
	"gowhole/project/spider/model"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan interface{}
}

type Scheduler interface {
	WorkerChan() chan model.Request
	Submit(model.Request)
	Run()
	ReadyNotifier
}

type ReadyNotifier interface {
	WorkReady(chan model.Request)
}

func (ce *ConcurrentEngine) Run(seeds ...model.Request) {
	out := make(chan model.ParseResult)
	ce.Scheduler.Run()
	for i := 0; i < ce.WorkerCount; i++ {
		ce.createWorker(ce.Scheduler.WorkerChan(), out, ce.Scheduler)
	}
	for _, r := range seeds {
		// if isDuplicate(r.URL) {
		// 	continue
		// }
		ce.Scheduler.Submit(r)
	}
	for {
		result := <-out
		for _, item := range result.Items {
			go func() {
				ce.ItemChan <- item
			}()
		}
		for _, r := range result.Requests {
			// if isDuplicate(r.URL) {
			// 	continue
			// }
			ce.Scheduler.Submit(r)
		}
	}
}

func (ce *ConcurrentEngine) createWorker(in chan model.Request, out chan model.ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkReady(in)
			request := <-in
			result, err := Worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
