package main

import (
	"gostudy/net/rpcx"
)

func main() {
	rpcx.SynchronousCall()
	rpcx.AsynchronousCall()
	//rpcx.PrintWrr()
	//rpcx.PrintWrrNgx()
}
