package main

import rpcx "gowhole/exercise/netx/rpcx/client"

func main() {
	rpcx.SynchronousCall()
	rpcx.AsynchronousCall()
	//rpcx.PrintWrr()
	//rpcx.PrintWrrNgx()
}
