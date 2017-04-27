package main

import (
	"gostudy/net/rpcx"
)

func main() {
	rpcx.Start()
	signalHandler()
}
