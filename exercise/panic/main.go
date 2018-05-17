package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"
)

var (
	// exeName       = os.Args[0]  //获取程序名称
	pid       = os.Getpid() //获取进程ID
	curDir, _ = os.Getwd()  //获取当前目录
	// _, file, _, _ = runtime.Caller(1)
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
	defer panicHandler()
	fmt.Printf("Calling test\r\n")
	test() //调用我们定义的函数，发现程序并没有崩溃，而是可以继续执行下一行代码的哟！
	fmt.Printf("Test completed\r\n")
}

func panicHandler() {
	filename := fmt.Sprintf(curDir+"/"+"%s-%d.log", "panicHandler", pid) //保存错误信息文件名:程序名-进程ID-当前时间（年月日时分秒）
	fmt.Printf("save  to path(%s)\n", filename)
	f, err := os.Create(filename)
	if err != nil {
		fmt.Printf("save to error(%s)\n", err)
		return
	}
	defer f.Close()
	if err := recover(); err != nil {
		f.WriteString(fmt.Sprintf("%v\r\n", err)) //输出panic信息
		f.WriteString("========\r\n")
	}
	f.WriteString(string(debug.Stack())) //输出堆栈信息
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
