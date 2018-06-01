package signalcall

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Handle() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		fmt.Printf("get signal s(%v)|s.String(%s)\n", s, s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT:
			return
		case syscall.SIGHUP:
		// TODO reload
		default:
			return
		}
	}
}
