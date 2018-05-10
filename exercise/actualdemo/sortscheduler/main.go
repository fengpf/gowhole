package main

import (
	"fmt"
	"sync"
)

type queue []chan int

func main() {
	ids := []int{1, 2, 3, 4, 5}
	queue := make([]chan int, 0)
	for _, id := range ids {
		w := make(chan int, 10)
		w <- id
		queue = append(queue, w)
	}

	var wg sync.WaitGroup
	for _, w := range queue {
		select {
		case j := <-w:
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				worker(i)
			}(j)
		default:
			fmt.Println("no")
		}
	}
	wg.Wait()
}

func worker(i int) int {
	i = i + 1
	return i
}
