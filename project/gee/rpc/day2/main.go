package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"gowhole/project/gee/rpc"
)

func startServer(addr chan string) {
	// pick a free port
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("network error:", err)
	}
	log.Println("start rpc server on", l.Addr().String())
	addr <- l.Addr().String()
	rpc.Accept(l)
}

func main() {
	//启动rpc server
	log.SetFlags(0)
	addr := make(chan string)
	go startServer(addr)

	//启动rpc client
	conn, _ := rpc.Dial("tcp", <-addr)
	defer func() { _ = conn.Close() }()
	time.Sleep(time.Second)

	// send request & receive response
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			args := fmt.Sprintf("rpc req %d", i)
			var reply string
			if err := conn.Call("Foo.Sum", args, &reply); err != nil {
				log.Fatal("call Foo.Sum error:", err)
			}

			log.Println("reply:", reply)
		}(i)
	}
	wg.Wait()
}
