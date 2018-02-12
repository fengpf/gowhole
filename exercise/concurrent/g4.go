package main

import (
	"fmt"
	"time"
)

func main() {
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
