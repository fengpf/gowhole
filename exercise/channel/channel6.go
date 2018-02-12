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
	//close(ch11),为了看效果先注释掉

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
			case ok = <-func() chan bool {
				//经过大约1ms后，该接收语句会从timeout通道接收到一个新元素并赋值给ok,从而恰当地执行了针对单个操作的超时子流程，恰当地结束当前for循环
				timeout := make(chan bool, 1)
				go func() {
					time.Sleep(time.Millisecond) //休息1ms
					timeout <- false
				}()
				return timeout
			}():
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
