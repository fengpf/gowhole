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

//test
func foo() *int {
	var x int
	return &x
}

func bar() int {
	x := new(int)
	*x = 1
	return *x
}

// ./main.go:61:9: &x escapes to heap
// ./main.go:60:6: moved to heap: x
// ./main.go:65:10: bar new(int) does not escape
// 上面的意思是 foo() 中的 x 最后在堆上分配，而 bar() 中的 x 最后分配在了栈上
