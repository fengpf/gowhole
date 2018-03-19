package main

import (
	"fmt"
	"time"
)

type atomicInt int

func (a *atomicInt) increment() {
	*a++
}
func (a *atomicInt) get() int {
	return int(*a)
}

func main() {
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
