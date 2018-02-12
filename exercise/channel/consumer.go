package main

import (
	"fmt"
	"time"
)

func consumer(queue <-chan int) {
	for i := 0; i < 10; i++ {
		v := <-queue
		fmt.Printf("consumer:%d\n", v)
	}
}

func producer(queue chan<- int) {
	for i := 0; i < 10; i++ {
		println("start produce:")
		queue <- i
	}
}

func main() {
	queue := make(chan int, 1)
	go consumer(queue)
	go producer(queue)
	time.Sleep(1e9)
}
