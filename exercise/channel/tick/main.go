package main

import (
	"fmt"
<<<<<<< HEAD
	"sync"
=======
	"os"
	"os/signal"
	"sync"
	"syscall"
>>>>>>> 2247aafc530d89948b78988019120a3a9381a16f
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

func test3() {
	var wg sync.WaitGroup
	ch := make(chan int, 100)
	chSend := make(chan int)
	chConsume := make(chan int)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		os.Kill,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	go func(ch, quit chan int) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("send to ch panic.===", err)
			}
		}()
		i := 0
		for {
			select {
			case ch <- i:
				fmt.Println("send", i)
				time.Sleep(time.Second)
				i++
			case <-quit:
				fmt.Println("send quit.")
				return
			}
		}
	}(ch, chSend)
	go func(ch, quit chan int) {
		wg.Add(1)
		for {
			select {
			case i, ok := <-ch:
				if ok {
					fmt.Println("read1", i)
					time.Sleep(time.Second * 2)
				} else {
					fmt.Println("close ch1.")
				}
			case <-quit:
				for {
					select {
					case i, ok := <-ch:
						if ok {
							fmt.Println("read2", i)
							time.Sleep(time.Second * 2)
						} else {
							fmt.Println("close ch2.")
							goto L
						}
					}
				}
			L:
				fmt.Println("consume quit.")
				wg.Done()
				return

			}
		}
	}(ch, chConsume)
	<-sc
	close(ch)
	fmt.Println("close ch ")
	close(chSend)
	close(chConsume)
	wg.Wait()
}

func test2() {
	jobs := make(chan int)
	timeout := make(chan bool)
	var wg sync.WaitGroup
	go func() {
		time.Sleep(time.Second * 3)
		timeout <- true
	}()
	go func() {
		for i := 0; ; i++ {
			select {
			case <-timeout:
				close(jobs)
				return
			default:
				jobs <- i
				fmt.Println("produce:", i)
			}
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range jobs {
			fmt.Println("consume:", i)
		}
	}()
	wg.Wait()
}
