package main

import (
	"fmt"
	"log"
	"net/rpc"
)

// Args param.
type Args struct {
	A, B int
}

func sync() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:7777")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// rpc 客户端同步调用程序执行
	args := &Args{7, 8}
	var reply int
	err = client.Call("Arith.Mutiply", args, &reply)
	fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)
}

func async() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:7777")
	if err != nil {
		log.Fatal("arith error:", err)
	}
	// rpc 客户端异步调用程序执行
	args := &Args{6, 3}
	var reply int
	divCall := client.Go("Arith.Divide", args, &reply, nil)
	replyCall := <-divCall.Done
	println(replyCall)
}

func main() {
	//sync()
	async()
}
