package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	count = 100
)

type sender struct {
	myChan chan int
	wg     sync.WaitGroup
}

var s *sender

func init() {
	// runtime.GOMAXPROCS(runtime.NumCPU())
	s = &sender{
		myChan: make(chan int, 1024),
	}
	for i := 0; i < count; i++ {
		go func(i int) {
			s.myChan <- i
		}(i)
	}
}

func main() {
	s.wg.Add(1)
	go s.send()
	s.wg.Wait()
}

func (s *sender) send() {
	defer s.wg.Done()
	tick := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-tick.C:
			sends := []int{}
			var done bool
			for {
				if done {
					break
				} else if len(sends) == 10 {
					fmt.Println("=10...")
					break
				}
				select {
				case data := <-s.myChan:
					sends = append(sends, data)
				default:
					fmt.Println("done = true")
					done = true
				}
			}
			if len(sends) > 0 {
				fmt.Println(sends)
			}
		}
	}
}

func tick() {
	ch := make(chan int, 1024)
	go func(ch chan int) {
		for {
			val := <-ch
			fmt.Printf("val:%d\n", val)
		}
	}(ch)

	tick := time.NewTicker(1 * time.Second)
	for i := 0; i < 20; i++ {
		select {
		case ch <- i:
		}
		select {
		case <-tick.C:
			fmt.Printf("%d: case <-tick.C\n", i)
		default:
		}
		time.Sleep(200 * time.Millisecond)
	}
	close(ch)
	tick.Stop()
}
