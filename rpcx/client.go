package rpcx

import (
	"fmt"
	"log"
	"net/rpc"
)

// SynchronousCall 同步调用.
func SynchronousCall() {
	var reply int
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:9001")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := &Args{7, 8}
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("同步调用Arith: %d*%d=%d\n", args.A, args.B, reply)
}

// AsynchronousCall 异步调用.
func AsynchronousCall() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:9001")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := &Args{16, 8}
	quotient := new(Quotient)
	divCall := client.Go("Arith.Divide", args, quotient, nil)
	replyCall := <-divCall.Done // will be equal to divCall
	// check errors, print, etc.
	fmt.Printf("异步调用Divide: %v\n", replyCall.Reply)
}
