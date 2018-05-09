package scheduler

import "gowhole/exercise/actualdemo/demoscheduler/model"

type DataScheduler struct {
	msgChan    chan *model.Msg
	workerChan chan chan *model.Msg
}

func (ds *DataScheduler) Submit(r *model.Msg) {
	ds.msgChan <- r
}

func (ds *DataScheduler) WorkerChan() chan *model.Msg {
	return make(chan *model.Msg)
}

func (ds *DataScheduler) WorkReady(m chan *model.Msg) {
	ds.workerChan <- m
}

func (ds *DataScheduler) Run() {
	ds.msgChan = make(chan *model.Msg)
	ds.workerChan = make(chan chan *model.Msg)
	go func() {
		var (
			msgQ    []*model.Msg
			workerQ []chan *model.Msg
		)
		for {
			var (
				activeMsg    *model.Msg
				activeWorker chan *model.Msg
			)
			if len(msgQ) > 0 && len(workerQ) > 0 {
				activeMsg = msgQ[0]
				activeWorker = workerQ[0]
			}
			select {
			case in := <-ds.msgChan:
				msgQ = append(msgQ, in)
			case out := <-ds.workerChan:
				workerQ = append(workerQ, out)
			case activeWorker <- activeMsg:
				msgQ = msgQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}
