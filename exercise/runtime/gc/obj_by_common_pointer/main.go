package main

import (
	"runtime"
	"time"
)

var ptrs []*data

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
	ptrs = append(ptrs, &d) //每次创建100MB的对象，然后将其指针保存到全局对象ptrs
}

// go build -o main -gcflags "-N -l" && GODEBUG=gctrace=1   ./main

// gc 1 @0.001s 0%: 0.004+48+0.025 ms clock, 0.016+0/0.011/48+0.10 ms cpu, 100->100->100 MB, 101 MB goal, 4 P (forced)
// gc 2 @0.050s 0%: 0.012+0.045+0.013 ms clock, 0.048+0/0.035/0.047+0.054 ms cpu, 100->100->100 MB, 200 MB goal, 4 P (forced)
// gc 3 @1.052s 0%: 0.003+50+0.061 ms clock, 0.014+0/0/50+0.24 ms cpu, 200->200->200 MB, 201 MB goal, 4 P
// gc 4 @1.103s 0%: 0.012+0.074+0.017 ms clock, 0.049+0/0.036/0.049+0.070 ms cpu, 200->200->200 MB, 400 MB goal, 4 P (forced)
// gc 5 @2.105s 0%: 0.004+67+0.027 ms clock, 0.017+0/0.005/67+0.11 ms cpu, 300->300->300 MB, 400 MB goal, 4 P

//结论：
//普通指针让每次创建的对象可达，无法被回收，内存膨胀
