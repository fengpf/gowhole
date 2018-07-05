package main

import (
	"fmt"
	"time"
)

// func init() {
// 	runtime.GOMAXPROCS(runtime.NumCPU())
// }

func main() {
	ch := make(chan int, 1024)
	go func(ch chan int) {
		for {
			val := <-ch
			fmt.Printf("val:%d\n", val)
		}
	}(ch)

	tick := time.NewTicker(1 * time.Second)
	for i := 0; i < 20; i++ {
		select {
		case ch <- i:
		}
		select {
		case <-tick.C:
			fmt.Printf("%d: case <-tick.C\n", i)
		default:
		}
		time.Sleep(200 * time.Millisecond)
	}
	close(ch)
	tick.Stop()
}
