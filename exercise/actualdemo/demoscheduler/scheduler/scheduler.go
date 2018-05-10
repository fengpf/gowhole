package scheduler

import (
	"gowhole/exercise/actualdemo/demoscheduler/model"
)

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
			ids     []int32
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
			case msg := <-ds.msgChan:
				ids = append(ids, msg.ID)
				// fmt.Println("ids", ids)
				msgQ = append(msgQ, msg)
			case worker := <-ds.workerChan:
				workerQ = append(workerQ, worker)
			case activeWorker <- activeMsg:
				// for _, m := range msgQ {
				// 	fmt.Println("m.ID", m.ID)
				// }
				msgQ = msgQ[1:]
				// for _, m := range msgQ {
				// 	fmt.Println("m.ID---------", m.ID)
				// }
				workerQ = workerQ[1:]
			}
		}
	}()
}
