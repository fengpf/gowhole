package main

import (
	"fmt"
)

func main() {
	//初始化通道
	ch11 := make(chan int, 1000)
	sign := make(chan int, 1)

	//给ch11通道写入数据
	for i := 0; i < 1000; i++ {
		ch11 <- i
	}
	//关闭ch11通道
	close(ch11)

	//单独起一个Goroutine执行select
	go func() {
		var e int
		ok := true
		for {
			select {
			case e, ok = <-ch11:
				if !ok {
					fmt.Println("End.")
					break
				}
				fmt.Printf("ch11 -> %d\n", e)
			}
			//通道关闭后退出for循环
			if !ok {
				sign <- 0
				println("tui")
				break
			}
		}

	}()

	//惯用手法，读取sign通道数据，为了等待select的Goroutine执行。
	<-sign
}
