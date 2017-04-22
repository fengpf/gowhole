package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"gostudy/rpcx"
)

func main() {
	// err := imgprocessing.WaterMark("golang水印")
	// if err != nil {
	// 	fmt.Printf("%v\n", err)
	// 	return
	// }
	//imgprocessing.HTTPPrint()
	//fmt.Printf("%s\n", "create watermark success ...")
	rpcx.Start()
	signalHandler()
}

func signalHandler() {
	// signal handler
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGSTOP)
	for {
		s := <-c
		println("get a signal")
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT:
			time.Sleep(time.Second * 2)
			println("exit")
			return
		case syscall.SIGHUP:
		// TODO reload
		default:
			return
		}
	}
}
