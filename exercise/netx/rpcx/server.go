package main

import (
	rpcx "gowhole/exercise/netx/rpcx/server"

	"gowhole/exercise/os/signalcall"
)

func main() {
	rpcx.Start()
	signalcall.Handle()
}
