package main

import (
	"errors"
	"log"
	"sync"
	"time"
)

var (
	task = []Task{}
)

func main() {
	var err error

	tw := New(15)
	//添加任务
	err = tw.AddTask(
		time.Now().Add(time.Second*1), "task-1",
		func(args ...interface{}) {
			log.Println(args...)
		},
		[]interface{}{1, 2, 3})

	err = tw.AddTask(time.Now().Add(time.Second*5), "task-2", func(args ...interface{}) {
		log.Println(args...)
	}, []interface{}{4, 5, 6})

	err = tw.AddTask(time.Now().Add(time.Second*10), "task-3", func(args ...interface{}) {
		log.Println(args...)
	}, []interface{}{"hello", "world"})

	err = tw.AddTask(time.Now().Add(time.Second*15), "task-4", func(args ...interface{}) {
		sum := 0
		for arg := range args {
			sum += arg
		}
		log.Println("sum:", sum)
	}, []interface{}{1, 2, 3})

	if err != nil {
		log.Fatal(err)
	}

	time.AfterFunc(time.Second*20, func() {
		tw.Close()
	})

	tw.Start()
}

type TaskFunc func(args ...interface{})

type Task struct {
	id         int
	cycleTimes int
	run        TaskFunc
	params     []interface{}
}

type TimeWheel struct {
	degree    int
	curIndex  int
	buckets   []map[string]*Task
	close     chan struct{}
	taskClose chan struct{}
	timeClose chan struct{}
	startTime time.Time

	wg sync.WaitGroup
}

//degree 精度，15表示15s
func New(degree int) *TimeWheel {
	tw := &TimeWheel{
		degree:    degree,
		buckets:   make([]map[string]*Task, degree),
		close:     make(chan struct{}),
		taskClose: make(chan struct{}),
		timeClose: make(chan struct{}),
	}

	for i := 0; i < degree; i++ {
		tw.buckets[i] = make(map[string]*Task)
	}
	return tw
}

func (tw *TimeWheel) Start() {
	tw.wg.Add(2)
	go tw.taskRun()
	go tw.timeRun()

	select {
	case <-tw.close:
		tw.taskClose <- struct{}{}
		tw.timeClose <- struct{}{}
		log.Println("Start close")
		return
	}
}

func (tw *TimeWheel) AddTask(t time.Time, key string, run TaskFunc, params []interface{}) (err error) {
	if tw.startTime.After(t) {
		err = errors.New("AddTask time expire")
		log.Printf("AddTask error(%v)", err)
		return
	}

	subSecond := t.Unix() - tw.startTime.Unix() //当前时间与指定时间相差秒数

	tasks := tw.buckets[int(subSecond)%tw.degree]
	if _, ok := tasks[key]; ok {
		err = errors.New("该buckets中已存在key为" + key + "的任务")
		log.Printf("AddTask error(%v)", err)
		return
	}
	tasks[key] = &Task{
		cycleTimes: int(subSecond / 3600), //计算循环次数
		run:        run,
		params:     params,
	}

	return
}

func (tw *TimeWheel) taskRun() {
	defer tw.wg.Done()

	for {
		select {
		case <-tw.taskClose:
			log.Println("taskRun close")

		default:
			tasks := tw.buckets[tw.curIndex]
			if len(tasks) > 0 {
				for k, v := range tasks {
					if v.cycleTimes == 0 {

						go v.run(v.params)

						delete(tasks, k)
					} else {
						v.cycleTimes--
					}
				}
			}
		}
	}
}

func (tw *TimeWheel) timeRun() {
	defer tw.wg.Done()

	tick := time.Tick(time.Second)
	for {
		select {
		case <-tw.timeClose:
			log.Println("timeRun close")
			return

		case <-tick:
			//log.Println("当前时间", time.Now().Format("2006-01-02 15:04:05"));
			if tw.curIndex == tw.degree-1 {
				tw.curIndex = 0 //重置
			} else {
				tw.curIndex++
			}
		}
	}
}

func (tw *TimeWheel) Close() {
	tw.wg.Wait()
	tw.close <- struct{}{}
}
