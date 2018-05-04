package main

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

var waitgroup sync.WaitGroup

func Afunction(shownum int) {
	fmt.Println(shownum)
	waitgroup.Done() //任务完成，将任务队列中的任务数量-1，其实.Done就是.Add(-1)
}

func Test_wg(t *testing.T) {
	for i := 0; i < 10; i++ {
		waitgroup.Add(1) //每创建一个goroutine，就把任务队列中任务的数量+1
		go Afunction(i)
	}
	waitgroup.Wait() //.Wait()这里会发生阻塞，直到队列中所有的任务结束就会解除阻塞
	println(111)
}

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			// println(k)
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func Test_gen(t *testing.T) {
	// Set up the pipeline.
	c := gen(1, 2, 3, 4)
	out := sq(c)
	// Consume the output.
	for o := range out {
		fmt.Println(o)
	}
}

func Test_timeout(t *testing.T) {
	var out bool
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case i := <-c:
				fmt.Println(i)
			case <-time.After(time.Duration(3) * time.Second): //设置超时时间为３ｓ，如果channel　3s钟没有响应，一直阻塞，则报告超时，进行超时处理．
				o <- true
				break
			}
		}
	}()
	go func() {
		c <- 2
	}()
	println(1111)
	out = <-o
	println(2222)
	if out {
		fmt.Println("timeout")
	}
}

func Test_goSched(t *testing.T) {
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
