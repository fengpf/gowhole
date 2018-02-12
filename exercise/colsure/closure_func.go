package main

import (
	"sync"
)

func myFunc() func() int {
	foo := 0
	return func() int {
		foo++
		return foo
	}
}

func foo() {
	var a = 1
	var wg sync.WaitGroup
	wg.Add(1)
	go func(a int) {
		a++
		println(a)
		wg.Done()
	}(a)
	wg.Wait()
	println(a)
}

func foo2() {
	var a = 1
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		a++
		println(a)
		wg.Done()
	}()
	wg.Wait()
	println(a)
}

func foo3() {
	var a = 1
	var wg sync.WaitGroup
	wg.Add(1)
	go func() (a int) {
		a++
		println(a)
		wg.Done()
		return
	}()
	wg.Wait()
	println(a)
}

func main() {
	bar := myFunc()
	val := bar()
	val2 := bar()
	val3 := bar()
	println(val)  // 1
	println(val2) // 2
	println(val3) // 3

	foo()
	foo2()
	foo3()
}
