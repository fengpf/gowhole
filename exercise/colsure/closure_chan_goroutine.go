package main

import (
	"fmt"
	"sync"
)

// 非缓冲channel 结合 groutine用法

// func main() {
// 	var waitNum sync.WaitGroup
// 	waitNum.Add(1)
// 	var c = make(chan []int)
// 	go func(c chan []int) {
// 		arr := make([]int, 20)
// 		for i := 0; i <= 19; i++ {
// 			arr[i] = i
// 			fmt.Println("i=", i)
// 		}
// 		c <- arr
// 		close(c)
// 		// for v := range c {
// 		// 	fmt.Println("get c value=", v)
// 		// }
// 		waitNum.Done()
// 		return
// 	}(c)
// 	var d = make([]int, 20)
// 	d = <-c
// 	fmt.Printf("d--->%+v\n", d)
// 	for _, v := range d {
// 		fmt.Println("v=", v)
// 	}
// }

// 缓冲channel 结合 groutine用法
func main() {
	var waitNum sync.WaitGroup
	waitNum.Add(1)
	var c = make(chan int, 20)
	go func(c chan int) {
		for i := 0; i <= 19; i++ {
			fmt.Println("set i=", i)
			c <- i
		}
		close(c)
		waitNum.Done()
		return
	}(c)
	var d = make([]int, 0, 20)
	for v := range c {
		d = append(d, v)
	}
	for _, v := range d {
		fmt.Println("val=", v)
	}
}
