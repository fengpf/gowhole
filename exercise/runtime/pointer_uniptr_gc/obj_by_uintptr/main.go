package main

import (
	"runtime"
	"time"
	"unsafe"
)

var ptrs []uintptr

type data struct {
	x [100 << 20]byte
}

func main() {
	for i := 0; i < 100; i++ {
		test()
		runtime.GC()
		time.Sleep(time.Second)
	}
}

func test() {
	var d data
	ptrs = append(ptrs, uintptr(unsafe.Pointer(&d))) //每次创建100MB的对象，然后将其指针保存到全局对象ptrs
}

// go build -o main -gcflags "-N -l" && GODEBUG=gctrace=1   ./main

// gc 1 @0.001s 0%: 0.004+93+0.032 ms clock, 0.016+0/0.006/93+0.12 ms cpu, 100->100->0 MB, 101 MB goal, 4 P (forced)
// gc 2 @0.095s 0%: 0.004+0.098+0.014 ms clock, 0.018+0/0.086/0+0.059 ms cpu, 0->0->0 MB, 8 MB goal, 4 P (forced)
// gc 3 @1.108s 0%: 0.004+6.0+0.026 ms clock, 0.018+0/0.019/6.0+0.10 ms cpu, 100->100->0 MB, 101 MB goal, 4 P
// gc 4 @1.114s 0%: 0.009+0.054+0.012 ms clock, 0.037+0/0.045/0.040+0.049 ms cpu, 0->0->0 MB, 8 MB goal, 4 P (forced)
// gc 5 @2.122s 0%: 0.003+5.0+0.024 ms clock, 0.014+0/0.011/5.0+0.098 ms cpu, 100->100->0 MB, 101 MB goal, 4 P
// gc 6 @2.127s 0%: 0.008+0.098+0.014 ms clock, 0.032+0/0.037/0.052+0.056 ms cpu, 0->0->0 MB, 8 MB goal, 4 P (forced)

//结论：
//uintptr 让每次创建的对象不可达，对象被回收
//在runtime里面，uintptr被当做类似弱引用来使用
