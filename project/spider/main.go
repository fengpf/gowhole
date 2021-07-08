package main

import (
	"gowhole/project/spider/engine"
	"gowhole/project/spider/model"
	"gowhole/project/spider/parser"
	"gowhole/project/spider/persist"
	"gowhole/project/spider/scheduler"
)

const (
	zhenaiURL = "http://www.zhenai.com/zhenghun"
)

func main() {
	ce := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 10,
		ItemChan:    persist.ItemSaver(),
	}
	ce.Run(model.Request{
		URL:       zhenaiURL,
		ParseFunc: parser.ParseCityList,
	})
	// ce.Run(model.Request{
	// 	URL:       "http://www.zhenai.com/zhenghun/shanghai",
	// 	ParseFunc: parser.ParseCity,
	// })
	return
}
