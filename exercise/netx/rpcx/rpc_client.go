package main

import (
	"gostudy/lib/net/rpcx"
)

func main() {
	rpcx.SynchronousCall()
	rpcx.AsynchronousCall()
	//rpcx.PrintWrr()
	//rpcx.PrintWrrNgx()
}
