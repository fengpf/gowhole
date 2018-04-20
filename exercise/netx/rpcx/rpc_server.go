package main

import (
	"gostudy/lib/net/rpcx"
)

func main() {
	rpcx.Start()
	signalHandler()
}
