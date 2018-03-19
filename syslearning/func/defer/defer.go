package main

import "fmt"

func closure(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func main() {
	var fs [4]func()
	{
	}
	for i := 0; i < 4; i++ {
		// i := i
		defer fmt.Println("defer i = ", i, &i)
		defer func() {
			fmt.Println("defer_closure i = ", i, &i)
		}()
		// fmt.Println("before i = ", i)
		fs[i] = func() {
			fmt.Println("closure i = ", i, &i)
		}
		// fmt.Println("after i = ", i)
	}

	for _, f := range fs {
		f()
	}
}
