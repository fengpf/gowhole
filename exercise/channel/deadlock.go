package main

import (
	"fmt"
)

var ch = make(chan int)
var ch1 = make(chan int)
var quit chan int // 只开一个信道

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", i)
	}
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n\r", i)
	}
	ch <- 100 // 向ch中加数据，如果没有其他goroutine来取走这个数据，那么挂起foo, 直到main函数把0这个数据拿走
}

func say(s string) {
	fmt.Println(s)
	ch <- <-ch1 //ch 等待ch1流出的数据
}

func foo2(id int) {
	quit <- id // ok, finished
}

func main() {
	// go loop()
	// loop()
	// time.Sleep(time.Second)

	// msg := make(chan string)
	// go func(m string) {
	// 	msg <- m //存消息
	// }("ping")
	// fmt.Println(<-msg)

	//如果不用信道来阻塞主线的话，主线就会过早跑完，loop线都没有机会执行
	// go foo()
	// go func() {
	// ch <- 1 //单线死锁
	// }()
	// fmt.Printf("直到线程跑完, 取到消息. main在此阻塞住!!! %v\n", <-ch)

	//多线死锁
	// go say("Hello")
	// <-ch // 堵塞主线

	//缓冲信道死锁
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// ch <- 3
	// fmt.Println(<-ch) // 1
	// fmt.Println(<-ch) // 2

	//如果不显式地关闭信道，下面的代码，会报死锁错误的，原因是range不等到信道关闭是不会结束读取的。
	close(ch)
	//也就是如果 缓冲信道干涸了，那么range就会阻塞当前goroutine, 所以死锁。
	// for v := range ch {
	// 	fmt.Println(v)
	// 	// if len(ch) <= 0 { // 如果现有数据量为0，跳出循环
	// 	// 	break
	// 	// }
	// }

	// 只使用单个无缓冲信道阻塞主线
	count := 1000
	// quit = make(chan int) // 无缓冲
	// for i := 0; i < count; i++ {
	// 	go foo2(i)
	// }
	// for i := 0; i < count; i++ {
	// 	println(<-quit)
	// }

	//使用容量为goroutines数量的缓冲信道
	var quit2 chan int
	quit2 = make(chan int, 1000) // 无缓冲
	for i := 0; i < count; i++ {
		quit2 <- i
	}
	close(quit2)
	for v := range quit2 {
		println(v)
		// if len(quit2) <= 0 {
		// 	break
		// }
	}
}
