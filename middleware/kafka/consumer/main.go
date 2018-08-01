package main

import (
	"fmt"
	"gowhole/lib/kafka/consumer/service"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	s := service.New()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT, syscall.SIGHUP)
	for {
		sg := <-c
		fmt.Printf("main exit by signal(%s)\n", sg.String())
		switch sg {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT:
			s.Close()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
