package channel

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup

func Test_run(t *testing.T) {
	//模拟4x4接力比赛
	baton := make(chan int)
	//等待最后一位跑者跑完
	wg.Add(1)
	go run(baton)
	fmt.Printf("%d runner with baton\n", 1)
	baton <- 1
	wg.Wait()
}

func run(baton chan int) {
	var newrunner int
	runner := <-baton
	fmt.Printf("runner %d run with baton\n", runner)
	//不是最后一位，传递给新的跑者
	if runner != 4 {
		newrunner = runner + 1
		go run(baton)
	}
	//模拟跑了100ms就传递接力棒
	time.Sleep(1 * time.Second)
	fmt.Printf("runner %d use 1s finish the round\n\n", runner)
	//如果是最后一位跑着，则知主线程结束
	if runner == 4 {
		wg.Done()
		return
	}
	fmt.Printf("runner %d exchange baton with newrunner %d \n", runner, newrunner)
	//传递接力棒
	baton <- newrunner
}
