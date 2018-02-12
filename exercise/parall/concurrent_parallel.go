package main

import (
	"fmt"
	"runtime"
	"time"
)

var quit = make(chan int)

func loop() {
	for i := 0; i < 10; i++ {
		runtime.Gosched() // 显式地让出CPU时间给其他goroutine
		fmt.Printf("%d  ", i)
	}
	quit <- 0
}

func foo(id int) {
	fmt.Println(id)
	time.Sleep(time.Second)
	quit <- 0
}

func say(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

// Gosched 让出cpu
// NumCPU 返回当前系统的CPU核数量
// GOMAXPROCS 设置最大的可同时使用的CPU核数
// Goexit 退出当前goroutine(但是defer语句会照常执行)

func main() {
	// 并发
	// go loop()
	// go loop()
	// for i := 0; i < 2; i++ {
	// 	<-quit
	// }

	// count := 1000
	// quit = make(chan int, count) // 缓冲1000个数据
	// for i := 0; i < count; i++ { //开1000个goroutine
	// 	go foo(i)
	// }
	// for i := 0; i < count; i++ { // 等待所有完成消息发送完毕
	// 	<-quit
	// }

	//并行
	// runtime.GOMAXPROCS(2) //最多使用2个核
	// go loop()
	// go loop()
	// for i := 0; i < 2; i++ {
	// 	<-quit
	// }

	go say("world") // 显式地让出CPU时间给其他goroutine
	for {

	}
}
