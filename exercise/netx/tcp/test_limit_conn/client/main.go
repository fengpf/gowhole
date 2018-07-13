package main

import (
	"fmt"
	"net"
	"runtime"
	"strconv"
	"time"
)

//test tcp最大连接数限制

func Connect(host string, port int) {
	_, err := net.Dial("tcp", host+":"+strconv.Itoa(port))
	if err != nil {
		fmt.Printf("Dial to %s:%d failed\n", host, port)
		return
	}

	for {
		time.Sleep(30 * 1000 * time.Millisecond)
	}
}

func main() {
	count := 0
	for {
		go Connect("127.0.0.1", 8080)
		count++
		fmt.Printf("Gorutue num:%d\n", runtime.NumGoroutine())
		time.Sleep(100 * time.Millisecond)
	}
}
