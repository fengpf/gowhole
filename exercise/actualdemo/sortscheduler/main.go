package main

import (
	"fmt"
	"sync"
)

type queue []chan int

func main() {
	count := 5
	ids := make([]int, 0)

	for index := 0; index < count; index++ {
		ids = append(ids, index)
	}

	queue := make([]chan int, 0)
	for _, id := range ids {
		var w chan int
		if id%2 == 0 {
			w = make(chan int, 5)
		}
		w <- id
		queue = append(queue, w)
	}

	var wg sync.WaitGroup
	for _, w := range queue {
		fmt.Println(w)
		j, ok := <-w
		if !ok {
			return
		}
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			worker(i)
		}(j)

	}
	wg.Wait()
}

func worker(i int) int {
	i = i * 2
	fmt.Println(i)
	return i
}
