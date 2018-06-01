package rpcx

import (
	"fmt"
	"log"
	"net/rpc"

	"gowhole/exercise/netx/rpcx/model"
)

// SynchronousCall 同步调用.
func SynchronousCall() {
	var reply int
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:9002")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := &model.Args{A: 7, B: 8}
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("同步调用Arith: %d*%d=%d\n", args.A, args.B, reply)
}

// AsynchronousCall 异步调用.
func AsynchronousCall() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:9002")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := &model.Args{A: 16, B: 8}
	quotient := new(model.Quotient)
	divCall := client.Go("Arith.Divide", args, quotient, nil)
	replyCall := <-divCall.Done // will be equal to divCall
	// check errors, print, etc.
	fmt.Printf("异步调用Divide: %v\n", replyCall.Reply)
}
