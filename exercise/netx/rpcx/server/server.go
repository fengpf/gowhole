package rpcx

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"

	"gowhole/exercise/netx/rpcx/model"
)

// Arith type.
type Arith int

// Multiply args.A * args.B.
func (t *Arith) Multiply(args *model.Args, reply *int) error {
	*reply = args.A * args.B
	println("客户端同步请求执行Multiply方法")
	return nil
}

// Divide args.A / args.B.
func (t *Arith) Divide(args *model.Args, quo *model.Quotient) error {
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
	l, e := net.Listen("tcp", ":9002")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	println("服务端开始监听...")
	//http.Serve(l, nil)
	go http.Serve(l, nil)
}
