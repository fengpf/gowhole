package main

import (
	"log"
	"net"
	"sync"
	"time"

	"gowhole/project/gee/rpc"
)

type Foo int

type Args struct{ Num1, Num2 int }

func (f Foo) Sum(args Args, reply *int) error {
	*reply = args.Num1 + args.Num2
	return nil
}

func startServer(addr chan string) {
	var foo Foo
	if err := rpc.Register(&foo); err != nil {//实例注册，注册 Foo 到 Server 中，并启动 RPC 服务
		log.Fatal("register error:", err)
	}

	// pick a free port
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("network error:", err)
	}

	log.Println("start rpc server on", l.Addr())
	addr <- l.Addr().String()
	rpc.Accept(l)
}

func main() {
	//启动rpc server
	log.SetFlags(0)
	addr := make(chan string)
	go startServer(addr)

	//启动rpc client
	client, _ := rpc.Dial("tcp", <-addr)
	defer func() { _ = client.Close() }()

	time.Sleep(time.Second)

	// send request & receive response
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			args := &Args{Num1: i, Num2: i * i}
			var reply int
			if err := client.Call("Foo.Sum", args, &reply); err != nil {
				log.Fatal("call Foo.Sum error:", err)
			}
			log.Printf("%d + %d = %d", args.Num1, args.Num2, reply)
		}(i)
	}
	wg.Wait()
}
