package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test_chan(t *testing.T) {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	select {
	case e1 := <-ch1:
		//如果ch1通道成功读取数据，则执行该case处理语句
		fmt.Printf("1th case is selected. e1=%v", e1)
	case e2 := <-ch2:
		//如果ch2通道成功读取数据，则执行该case处理语句
		fmt.Printf("2th case is selected. e2=%v", e2)
	default:
		//如果上面case都没有成功，则进入default处理流程
		fmt.Println("default!.")
	}
}
func Test_select(t *testing.T) {
	//定义几个变量，其中chs和numbers分别代表了包含了有限元素的通道列表和整数列表
	var ch1 chan int
	var ch2 chan int
	var chs = []chan int{ch1, ch2}
	var numbers = []int{1, 2, 3, 4, 5}
	getNumber := func(i int) int {
		fmt.Printf("numbers[%d]\n", i)
		return numbers[i]
	}
	getChan := func(i int) chan int {
		fmt.Printf("chs[%d]\n", i)
		return chs[i]
	}
	select {
	case getChan(0) <- getNumber(2):
		fmt.Println("1th case is selected.")
	case getChan(1) <- getNumber(3):
		fmt.Println("2th case is selected.")
	default:
		fmt.Println("default!.")
	}
}

func Test_cap(t *testing.T) {
	chanCap := 5
	ch7 := make(chan int, chanCap)
	for i := 0; i < chanCap; i++ {
		select {
		case ch7 <- 1:
		case ch7 <- 2:
		case ch7 <- 3:
		}
	}
	for i := 0; i < chanCap; i++ {
		fmt.Printf("%v\n", <-ch7)
	}
}
func Test_gosel(t *testing.T) {
	//初始化通道
	ch11 := make(chan int, 1000)
	sign := make(chan int, 1)
	//给ch11通道写入数据
	for i := 0; i < 1000; i++ {
		ch11 <- i
	}
	//关闭ch11通道
	close(ch11)
	//单独起一个Goroutine执行select
	go func() {
		var e int
		ok := true
		for {
			select {
			case e, ok = <-ch11:
				if !ok {
					fmt.Println("End.")
					break
				}
				fmt.Printf("ch11 -> %d\n", e)
			}
			//通道关闭后退出for循环
			if !ok {
				sign <- 0
				println("tui")
				break
			}
		}

	}()
	//惯用手法，读取sign通道数据，为了等待select的Goroutine执行。
	<-sign
}

func Test_timeout(t *testing.T) {
	//初始化通道
	ch11 := make(chan int, 1000)
	sign := make(chan int, 1)
	//给ch11通道写入数据
	for i := 0; i < 1000; i++ {
		ch11 <- i
	}
	//关闭ch11通道
	close(ch11)
	//我们不想等到通道被关闭之后再推出循环，我们创建并初始化一个辅助的通道，利用它模拟出操作超时行为
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(time.Millisecond) //休息1ms
		timeout <- false
	}()
	//单独起一个Goroutine执行select
	go func() {
		var e int
		ok := true

		for {
			select {
			case e, ok = <-ch11:
				if !ok {
					fmt.Println("End.")
					break
				}
				fmt.Printf("ch11 -> %d\n", e)
			case ok = <-timeout:
				//向timeout通道发送元素false后，该case几乎马上就会被执行, ok = false
				fmt.Println("Timeout.")
				break
			}
			//终止for循环
			if !ok {
				sign <- 0
				break
			}
		}

	}()
	//惯用手法，读取sign通道数据，为了等待select的Goroutine执行。
	<-sign
}

func Test_timeout2(t *testing.T) {
	//初始化通道
	ch11 := make(chan int, 1000)
	sign := make(chan int, 1)
	//给ch11通道写入数据
	for i := 0; i < 1000; i++ {
		ch11 <- i
	}
	//关闭ch11通道
	//close(ch11),为了看效果先注释掉
	//单独起一个Goroutine执行select
	go func() {
		var e int
		ok := true
		for {
			select {
			case e, ok = <-ch11:
				if !ok {
					fmt.Println("End.")
					break
				}
				fmt.Printf("ch11 -> %d\n", e)
			case ok = <-func() chan bool {
				//经过大约1ms后，该接收语句会从timeout通道接收到一个新元素并赋值给ok,从而恰当地执行了针对单个操作的超时子流程，恰当地结束当前for循环
				timeout := make(chan bool, 1)
				go func() {
					time.Sleep(time.Millisecond) //休息1ms
					timeout <- false
				}()
				return timeout
			}():
				fmt.Println("Timeout.")
				break
			}
			//终止for循环
			if !ok {
				sign <- 0
				break
			}
		}
	}()
	//惯用手法，读取sign通道数据，为了等待select的Goroutine执行。
	<-sign
}

func Test_unbufChan(t *testing.T) {
	unbufChan := make(chan int)
	sign := make(chan byte, 2)
	go func() {
		for i := 0; i < 10; i++ {
			select {
			case unbufChan <- i:
			case unbufChan <- i + 10:
			default:
				fmt.Println("default!")
			}
			time.Sleep(time.Second)
		}
		close(unbufChan)
		fmt.Println("The channel is closed.")
		sign <- 0
	}()
	go func() {
	loop:
		for {
			select {
			case e, ok := <-unbufChan:
				if !ok {
					fmt.Println("Closed channel.")
					break loop
				}
				fmt.Printf("e: %d\n", e)
				time.Sleep(2 * time.Second)
			}
		}
		sign <- 1
	}()
	<-sign
	<-sign
}

func Test_sync(t *testing.T) {
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

func consumer(queue <-chan int) {
	for i := 0; i < 10; i++ {
		v := <-queue
		fmt.Printf("consumer:%d\n", v)
	}
}

func producer(queue chan<- int) {
	for i := 0; i < 10; i++ {
		println("start produce:")
		queue <- i
	}
}
func Test_produce_consume(t *testing.T) {
	queue := make(chan int, 1)
	go consumer(queue)
	go producer(queue)
	time.Sleep(1e9)
}
