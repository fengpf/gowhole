package scheduler

import (
	"gowhole/exercise/actualdemo/queuespider/model"
)

type QueueScheduler struct {
	requestChan chan model.Request
	workerChan  chan chan model.Request
}

func (qs *QueueScheduler) WorkerChan() chan model.Request {
	return make(chan model.Request)
}

func (qs *QueueScheduler) Submit(r model.Request) {
	qs.requestChan <- r
}

func (qs *QueueScheduler) WorkReady(w chan model.Request) {
	qs.workerChan <- w
}

func (qs *QueueScheduler) Run() {
	qs.requestChan = make(chan model.Request)
	qs.workerChan = make(chan chan model.Request)
	go func() {
		var (
			requestQ []model.Request
			workerQ  []chan model.Request
		)
		for {
			var (
				activeRequest model.Request
				activeWorker  chan model.Request
			)
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			select {
			case r := <-qs.requestChan:
				//send request to a worker
				requestQ = append(requestQ, r)
			case w := <-qs.workerChan:
				//send next_request to w
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}
