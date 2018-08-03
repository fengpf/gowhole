package lib

import (
	"fmt"
	"sync"
	"time"
)

type MyStruct struct {
	A int
	B string
	C map[int]string
	D []string
}

func DBGTestRun(var1 int, var2 string, var3 []int, var4 MyStruct) {
	fmt.Println("DBGTestRun Begin!\n")
	waiter := &sync.WaitGroup{}

	waiter.Add(1)
	go RunFunc1(var1, waiter)

	waiter.Add(1)
	go RunFunc2(var2, waiter)

	waiter.Add(1)
	go RunFunc3(&var3, waiter)

	waiter.Add(1)
	go RunFunc4(&var4, waiter)

	waiter.Wait()
	fmt.Println("DBGTestRun Finished!\n")
}

func RunFunc1(variable int, waiter *sync.WaitGroup) {
	fmt.Printf("var1:%v\n", variable)
	for {
		if variable != 123456 {
			continue
		} else {
			break
		}
	}
	time.Sleep(10 * time.Second)
	waiter.Done()
}

func RunFunc2(variable string, waiter *sync.WaitGroup) {
	fmt.Printf("var2:%v\n", variable)
	time.Sleep(10 * time.Second)
	waiter.Done()
}

func RunFunc3(pVariable *[]int, waiter *sync.WaitGroup) {
	fmt.Printf("*pVar3:%v\n", *pVariable)
	time.Sleep(10 * time.Second)
	waiter.Done()
}

func RunFunc4(pVariable *MyStruct, waiter *sync.WaitGroup) {
	fmt.Printf("*pVar4:%v\n", *pVariable)
	time.Sleep(10 * time.Second)
	waiter.Done()
}
