package main

import (
	"fmt"
	"gowhole/exercise/actualdemo/spider/engine"
	"gowhole/exercise/actualdemo/spider/model"
	"gowhole/exercise/actualdemo/spider/parser"
	"sync"
)

const (
	zhenaiURL = "http://www.zhenai.com/zhenghun"
)

func main() {
	engine.Run(model.Request{
		URL:       zhenaiURL,
		ParseFunc: parser.ParseCityList,
	})
	return
}

type worker struct {
	job  chan int
	done func()
}

func doWork(i int, w worker) {
	for n := range w.job {
		fmt.Printf("Worker %d received %d\n", i, n)
		w.done()
	}
}

func genWorker(n int, wg *sync.WaitGroup) (wk worker) {
	wk = worker{
		job: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWork(n, wk)
	return
}

func dispatch() {
	var workers [10]worker
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		workers[i] = genWorker(i, &wg)
	}
	wg.Add(20)
	for i, w := range workers {
		w.job <- i //发完,stop发的要有人收
	}
	for i, w := range workers {
		w.job <- i + 1000 //stop 还没人就发，卡住
	}
	wg.Wait()
}
