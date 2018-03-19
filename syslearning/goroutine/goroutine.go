package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var a [10]int

	// for i := 0; i < 10; i++ {
	// 	go func() { //race condition
	// 		for {
	// 			a[i]++
	// 			// fmt.Printf("Hello from "+"goroutine %d\n", i)//io可以交出cpu控制权
	// 			runtime.Gosched() //主动交出控制权
	// 		}
	// 	}()
	// }

	for i := 0; i < 10; i++ {
		go func(ii int) { //race condition
			for {
				a[ii]++
				// fmt.Printf("Hello from "+"goroutine %d\n", ii)//io可以交出cpu控制权
				runtime.Gosched() //主动交出控制权
			}
		}(i)
	}

	for i := 0; i < 1000; i++ {
		go func(ii int) { //race condition
			for {
				fmt.Printf("Hello from "+"goroutine %d\n", ii) //io可以交出cpu控制权
			}
		}(i)
	}
	time.Sleep(time.Minute)
	fmt.Println(a) //a[10] panic: runtime error: index out of range
}
