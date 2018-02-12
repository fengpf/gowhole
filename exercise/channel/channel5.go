package main

import (
	"fmt"
	"time"
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

	//我们不想等到通道被关闭之后再推出循环，我们创建并初始化一个辅助的通道，利用它模拟出操作超时行为
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(time.Millisecond) //休息1ms
		timeout <- false
	}()

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
			case ok = <-timeout:
				//向timeout通道发送元素false后，该case几乎马上就会被执行, ok = false
				fmt.Println("Timeout.")
				break
			}
			//终止for循环
			if !ok {
				sign <- 0
				break
			}
		}

	}()

	//惯用手法，读取sign通道数据，为了等待select的Goroutine执行。
	<-sign
}
