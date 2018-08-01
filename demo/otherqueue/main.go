package main

import (
	"fmt"
	"sync"

	"gowhole/exercise/actualdemo/otherqueue/queue"
)

func main() {
	q := queue.NewQueue()
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			q.Put(i)
		}(i)
	}
	wg.Wait()
	fmt.Println(1111, q.Show()) //查看队列
	for {
		i, ok := q.Get() //出队
		if !ok {
			break
		}
		fmt.Println(i)
	}
	fmt.Println(222, q.Show())
	q.Delete(f) //清空队列
	fmt.Println(333, q.Show())
}

func f(i interface{}) bool {
	if i == nil {
		return false
	}
	return true
}
