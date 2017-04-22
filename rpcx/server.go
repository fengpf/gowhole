package rpcx

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

// Args a,b.
type Args struct {
	A, B int
}

// Quotient Quo,Rem.
type Quotient struct {
	Quo, Rem int
}

// Arith type.
type Arith int

// Multiply args.A * args.B.
func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	println("客户端同步请求执行Multiply方法")
	return nil
}

// Divide args.A / args.B.
func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	println("客户端异步请求执行Divide方法")
	return nil
}

// Start start a server.
func Start() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":9001")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
	//go http.Serve(l, nil)
}
