package main

import "fmt"

func main() {
	var ch = make(chan int, 20)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	//ch <- 11 //panic: runtime error: send on closed channel
	for i := range ch {
		fmt.Println(i)
	}
}
