package main

import (
	"fmt"
)

func server1(ch chan string) {
	// time.Sleep(time.Second * 3)
	ch <- "server1 start"
}

func server2(ch chan string) {
	// time.Sleep(time.Second * 1)
	ch <- "server2 start"
}

func process(ch chan string) {
	// time.Sleep(time.Millisecond * 2500)
	ch <- "process successful"
}

func main() {
	select {}
	// var wg sync.WaitGroup
	// var ch chan string
	ch := make(chan string)
	ch2 := make(chan string)
	// wg.Add(1)
	go server1(ch)
	go server2(ch2)
	// go process(ch)

	// time.Sleep(time.Millisecond * 1000)
	// for {
	select {
	case n := <-ch:
		fmt.Println(n)
	case n := <-ch2:
		fmt.Println(n)
		// default:
		// 	fmt.Printf("no value\n")
	}
	// }
	// wg.Wait()
}
