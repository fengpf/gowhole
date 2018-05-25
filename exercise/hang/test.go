package main

import "runtime"

// GODEBUG="schedtrace=300,scheddetail=1" ./test

func main() {
	var ch = make(chan int, 2)
	go func() {
		for i := 0; i < 2; i++ {
			ch <- 1
			if i == 88 {
				runtime.GC()
			}
		}
	}()

	for {
		// the wrong part
		if len(ch) == 2 {
			sum := 0
			itemNum := len(ch)
			for i := 0; i < itemNum; i++ {
				sum += <-ch
			}
			if sum == itemNum {
				return
			}
		}
	}
}
