package main

import (
	"fmt"
	"sync"
)

type worker struct {
	c chan job
}

type job chan int

var (
	count = 2
)

func send(i int, wk worker, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		i++
		j := make(chan int)
		go func(i int) {
			j <- i
		}(i)
		wk.c <- j
		fmt.Printf("send i->%d\n", i)
		// time.Sleep(time.Second)
	}
}

func receive(wk worker, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		if j, ok := <-wk.c; ok {
			fmt.Printf("recevie i->%d\n", <-j)
		}
		// time.Sleep(time.Second)
	}
}
func main() {
	var wg sync.WaitGroup
	wk := worker{
		c: make(chan job),
	}
	wg.Add(count * 3)
	for i := 0; i < count; i++ {
		go send(i, wk, &wg)
	}
	for i := 0; i < count; i++ {
		go receive(wk, &wg)
	}
	wg.Wait()
}
