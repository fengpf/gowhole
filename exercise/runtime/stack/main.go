package main

import (
	"fmt"
	"runtime"
	"sync"
)

var buf = make([]byte, 1<<20)

func test() {
	runtime.Stack(buf, true)
	println(222)
}

func main() {

	// fmt.Printf("%s", debug.Stack())
	// debug.PrintStack()

	runtime.Stack(buf, true)

	test()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		runtime.Stack(buf, true)
		println(111)
	}()
	wg.Wait()

	fmt.Printf("\n%s", buf)
}
