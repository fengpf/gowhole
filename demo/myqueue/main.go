package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"gowhole/exercise/actualdemo/myqueue/engine"
	"gowhole/exercise/actualdemo/myqueue/scheduler"
)

func main() {
	count := 50
	msg := make([]int, 0, count)
	for i := 0; i < count; i++ {
		if i%5 == 0 {
			msg = append(msg, i)
		}
	}

	de := engine.DispatchEngine{
		Scheduler:   &scheduler.DataScheduler{},
		WorkerCount: 5,
		Wg:          sync.WaitGroup{},
	}

	de.Wg.Add(1)
	t := time.Now()
	de.Run(msg)
	elapsed := time.Since(t)
	fmt.Printf("app elapsed(%v)\n", elapsed)
	de.Wg.Wait()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT, syscall.SIGHUP)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT:
			de.Stop()
			fmt.Println("engine dispatch exit...")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

func testChan() {
	c := make(chan int, 10)
	c <- 1
	c <- 2
	c <- 3
	close(c)

	// for {
	// 	fmt.Println(<-c)
	// }

	for {
		i, isOpen := <-c
		if !isOpen {
			fmt.Println("channel closed!")
			break
		}
		fmt.Println(i)
	}
}
