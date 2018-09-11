package worker

import (
	"log"
	"sync"
	"testing"
	"time"
)

var names = []string{
	"dsd",
	"bob",
	"jon",
	"tom",
	"jack",
}

type namPrinter struct {
	name string
}

func (m *namPrinter) Task() {
	log.Println(m.name)
	time.Sleep(time.Second)
}

func TestWorker(t *testing.T) {
	//使用两个goroutine 来创建工作池
	p := New(2)

	var wg sync.WaitGroup
	wg.Add(10 * len(names))

	for i := 0; i < 10; i++ {
		//迭代names 切片
		for _, name := range names {
			//创建一个 namPrinter 并且提供指定的名字
			np := namPrinter{
				name: name,
			}

			go func() {
				//将任务提交执行，当run返回的时候，则表示任务已经完成
				p.Run(&np)
				wg.Done()
			}()
		}

		wg.Wait()
	}

	//让工作池停止工作， 等待所有现有的工作完成
	p.Close()
}
