package main

import (
	"fmt"
)

func badCall() { //定义一个让程序运行时崩溃的函数
	panic("bad end")
}

func test() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s\n", e) //我们知道这个程序已经抛出了panic的错误了，但是我们用recover函数是可以处理这个错误的，我们这里的做法就是打印这个错误的输出并且不让程序崩溃。
		}
	}()
	badCall()                        //调用这个运行时崩溃的函数，因此下面的一行代码是不会被执行的，而是直接结束当前函数，而结束函数之后就会触发defer关键字，因此会被recover函数捕捉。
	fmt.Printf("After bad call\r\n") // <-- wordt niet bereikt
}

func main() {
	fmt.Printf("Calling test\r\n")
	test() //调用我们定义的函数，发现程序并没有崩溃，而是可以继续执行下一行代码的哟！
	fmt.Printf("Test completed\r\n")
}
