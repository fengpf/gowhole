package main

import (
	"fmt"
	"time"
)

type Job struct {
	Data interface{}
	Proc func(interface{})
}

//Job队列，存储要做的Job
var JobQueue chan Job

//Woker,用来从Job队列中取出Job执行
type Worker struct {
	WokerPool  chan chan Job //表示属于哪个Worker池,同时接收JobChannel注册
	JobChannel chan Job      //任务管道，通过这个管道获取任务执行
	Quit       chan bool     //用来停止Worker
}

//新建一个Worker,需要传入Worker池参数
func NewWorker(wokerPool chan chan Job) Worker {
	return Worker{
		WokerPool:  wokerPool,
		JobChannel: make(chan Job),
		Quit:       make(chan bool),
	}
}

//Worker的启动：包含：(1) 把该worker的JobChannel注册到WorkerPool中去  (2) 监听JobChannel上有没有新的任务到来 (3) 监听是否受到关闭的请求
func (worker Worker) Start() {
	go func() {
		for {
			worker.WokerPool <- worker.JobChannel //每次做完任务后就重新注册上去通知本worker又处于可用状态了
			select {
			case job := <-worker.JobChannel:
				job.Proc(job.Data)
			case quit := <-worker.Quit: //接收到关闭信息，直接退出即可
				if quit {
					return
				}
			}
		}
	}()
}

//Worker的关闭：只要发送一个关闭信号即可
func (worker Worker) Stop() {
	go func() {
		worker.Quit <- true
	}()
}

//管理Worker的调度器,包含最大worker数量和workerpool
type Dispatcher struct {
	MaxWorker  int
	WorkerPool chan chan Job
}

//启动一个调度器
func (dispatcher *Dispatcher) Run() {
	//启动maxworker个worker
	for i := 0; i < dispatcher.MaxWorker; i++ {
		worker := NewWorker(dispatcher.WorkerPool)
		worker.Start()
	}
	//接下来启动调度服务
	go dispatcher.dispatch()
}

func (dispatcher *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			go func(job Job) {
				jobChannel := <-dispatcher.WorkerPool //获取一个可用的worker
				jobChannel <- job                     //将该job发送给该worker
			}(job)
		}
	}
}

//新建一个调度器
func NewDispatcher(maxWorker int) *Dispatcher {
	workerPool := make(chan chan Job, maxWorker)
	return &Dispatcher{
		WorkerPool: workerPool,
		MaxWorker:  maxWorker,
	}
}

func handle(i interface{}) {
	j, ok := i.(int)
	if !ok {
		return
	}
	fmt.Println(j % 10)
	// fmt.Println(fib(j))
}

func fib(i int) (j int) {
	if i <= 1 {
		return 1
	}
	j = fib(i-1) + fib(i-2)
	return
}

func main() {
	dispatcher := NewDispatcher(10)
	dispatcher.Run()
	JobQueue = make(chan Job, 10)
	for i := 0; i < 10; i++ {
		JobQueue <- Job{
			Data: i,
			Proc: handle,
		}
	}

	time.Sleep(time.Second * 10)
}
