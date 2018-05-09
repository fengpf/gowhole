package main

import (
	"gowhole/exercise/actualdemo/concurrentspider/engine"
	"gowhole/exercise/actualdemo/concurrentspider/model"
	"gowhole/exercise/actualdemo/concurrentspider/parser"
	"gowhole/exercise/actualdemo/concurrentspider/scheduler"
)

const (
	zhenaiURL = "http://www.zhenai.com/zhenghun"
)

func main() {
	ce := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
	}
	ce.Run(model.Request{
		URL:       zhenaiURL,
		ParseFunc: parser.ParseCityList,
	})
	return
}
