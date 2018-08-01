package scheduler

import "gowhole/exercise/actualdemo/queuespider/model"

type SimpleScheduler struct {
	workerChan chan model.Request
}

func (s *SimpleScheduler) WorkerChan() chan model.Request {
	return s.workerChan
}

func (s *SimpleScheduler) Submit(r model.Request) {
	go func() {
		s.workerChan <- r
	}()
}

func (s *SimpleScheduler) Run(r model.Request) {
	s.workerChan = make(chan model.Request)
}

func (s *SimpleScheduler) WorkReady(r model.Request) {}
