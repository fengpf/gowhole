package scheduler

import "gowhole/exercise/actualdemo/concurrentspider/model"

type SimpleScheduler struct {
	workerChan chan model.Request
}

func (s *SimpleScheduler) ConfigMasterWorkerChan(c chan model.Request) {
	s.workerChan = c
}

func (s *SimpleScheduler) Submit(r model.Request) {
	go func() {
		s.workerChan <- r
	}()
}
