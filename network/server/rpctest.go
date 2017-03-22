package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

// Args param.
type Args struct {
	A, B int
}

// Arith 算术.
type Arith int

// Mutiply 乘法.
func (t *Arith) Mutiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

// Divide 除法
func (t *Arith) Divide(args *Args, reply *int) error {
	if args.B == 0 {
		return errors.New("Divide by zero")
	}
	*reply = args.A / args.B

	return nil
}

// Remainder 取余.
func (t *Arith) Remainder(args *Args, reply *int) error {
	if args.B == 0 {
		return errors.New("Divide by zero")
	}
	*reply = args.A % args.B
	return nil
}

// registerRPC 注册服务对象并开启该RPC服务.
func registerRPC() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", "127.0.0.1:7777")
	println(e)
	if e != nil {
		log.Fatal("Listen error:", e)
	}
	go http.Serve(l, nil)
}

func main() {
	registerRPC()
}
