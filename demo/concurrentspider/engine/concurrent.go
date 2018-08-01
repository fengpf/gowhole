package engine

import (
	"log"

	"gowhole/exercise/actualdemo/concurrentspider/model"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	ConfigMasterWorkerChan(chan model.Request)
	Submit(model.Request)
}

func (ce *ConcurrentEngine) Run(seeds ...model.Request) {
	in := make(chan model.Request)
	out := make(chan model.ParseResult)
	ce.Scheduler.ConfigMasterWorkerChan(in)
	for i := 0; i < ce.WorkerCount; i++ {
		log.Printf("create Worker(%d)\n", i)
		ce.createWorker(in, out)
	}
	for _, req := range seeds {
		log.Printf("first Submit req.URL(%s)\n", req.URL)
		ce.Scheduler.Submit(req)
	}
	itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			itemCount++
			log.Printf("get item(%+v)|itemCount(%d)\n", item, itemCount)
		}
		for _, req := range result.Requests {
			log.Printf("second Submit req.URL(%+v)\n", req.URL)
			ce.Scheduler.Submit(req)
		}
	}
}

func (ce *ConcurrentEngine) createWorker(in chan model.Request, out chan model.ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := Worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
