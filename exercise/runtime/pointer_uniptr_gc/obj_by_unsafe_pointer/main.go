package main

import (
	"runtime"
	"time"
	"unsafe"
)

var ptrs []unsafe.Pointer

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
	ptrs = append(ptrs, unsafe.Pointer(&d)) //每次创建100MB的对象，然后将其指针保存到全局对象ptrs
}

// go build -o main -gcflags "-N -l" && GODEBUG=gctrace=1   ./main

// gc 1 @0.001s 0%: 0.003+106+0.037 ms clock, 0.014+0/0.003/106+0.14 ms cpu, 100->100->100 MB, 101 MB goal, 4 P
// gc 2 @0.108s 0%: 0.003+0.11+0.015 ms clock, 0.014+0/0.097/0+0.060 ms cpu, 100->100->100 MB, 200 MB goal, 4 P (forced)
// gc 3 @1.110s 0%: 0.005+128+0.029 ms clock, 0.022+0/0/128+0.11 ms cpu, 200->200->200 MB, 201 MB goal, 4 P
// gc 4 @1.239s 0%: 0.092+0.10+0.031 ms clock, 0.36+0/0.084/0+0.12 ms cpu, 200->200->200 MB, 400 MB goal, 4 P (forced)

//结论：
//unsafe.Pointer 指针让每次创建的对象可达，无法被回收，内存膨胀
