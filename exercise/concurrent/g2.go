package main

import "fmt"

func Afuntion(ch chan int) {
	var i int
	fmt.Println("finish")
	i = <-ch
	println(i)
}

func main() {
	ch := make(chan int) //无缓冲的channel
	go Afuntion(ch)
	ch <- 100

	// 输出结果：
	// finish
}
