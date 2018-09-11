package runner

import (
	"log"
	"os"
	"testing"
	"time"
)

const (
	timeout = 1 * time.Millisecond //设置超时
	count   = 10                   //设置任务数量
)

func TestRunner(t *testing.T) {
	t.Log("Starting work.")

	//创建一个runner
	r := New(timeout)

	//添加要执行的任务
	for index := 0; index < count; index++ {
		r.Add(createTask())
	}

	//执行任务并且处理结果
	if err := r.Start(); err != nil {
		switch err {
		case ErrTimeout:
			t.Log("Terminating due to timeout")
			os.Exit(1)
		case ErrInterrupt:
			t.Log("Terminating due to interrupt")
			os.Exit(2)
		}
	}
}

//createTask 返回一个根据id休眠指定秒数的事例任务
func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d\n", id)
		// time.Sleep(time.Duration(id) * time.Second)
	}
}
