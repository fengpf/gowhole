package main

import (
	"fmt"
)

func main() {
	chanCap := 5
	ch7 := make(chan int, chanCap)

	for i := 0; i < chanCap; i++ {
		select {
		case ch7 <- 1:
		case ch7 <- 2:
		case ch7 <- 3:
		}
	}

	for i := 0; i < chanCap; i++ {
		fmt.Printf("%v\n", <-ch7)
	}
}
