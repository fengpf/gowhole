package main

import (
	"gowhole/exercise/actualdemo/queuespider/engine"
	"gowhole/exercise/actualdemo/queuespider/model"
	"gowhole/exercise/actualdemo/queuespider/parser"
	"gowhole/exercise/actualdemo/queuespider/scheduler"
)

const (
	zhenaiURL = "http://www.zhenai.com/zhenghun"
)

func main() {
	ce := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 100,
	}
	ce.Run(model.Request{
		URL:       zhenaiURL,
		ParseFunc: parser.ParseCityList,
	})
	return
}
