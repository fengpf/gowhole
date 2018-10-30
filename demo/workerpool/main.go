package main

import (
	"fmt"
)

func handle(i interface{}) {
	j, ok := i.(int)
	if !ok {
		return
	}

	fmt.Println(j)

	// fmt.Println(j % 10)
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
	dispatcher := NewDispatcher(2)
	dispatcher.Run()

	JobQueue = make(chan Job, 100)

	for i := 0; i < 100; i++ {
		JobQueue <- Job{
			Data: i,
			Proc: handle,
		}
	}

	// time.Sleep(time.Second * 10)
}
