package main

import (
	"fmt"
	"time"
)

var (
	consumeLimit = int64(10)
	tokenChan    = make(chan int, consumeLimit) //速率缓冲大小控制
	consumeRate  = int64(1e9 / consumeLimit)
)

func main() {
	a := []int{11, 22, 33, 44, 55, 66, 77, 88, 99}
	go generateToken(time.Duration(consumeRate))
	for _, v := range a {
		start := time.Now()
		<-tokenChan //消费速率控制
		elapsed := time.Since(start)
		fmt.Println(elapsed)
		fmt.Println(v)
	}

	time.Sleep(10 * time.Second)
}

func generateToken(duration time.Duration) {
	var (
		timer = time.NewTicker(duration)
		token = 0
	)
	for range timer.C {
		token++
		tokenChan <- token
	}
}
