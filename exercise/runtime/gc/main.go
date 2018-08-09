package main

import (
	"runtime"
)

func main() {
	var x [100 << 20]byte

	runtime.SetFinalizer(&x, func(o interface{}) {
		println("x dead")
	})

	p := &x //普通指针可达
	// p := unsafe.Pointer(&x) //unsafe.Pointer指针可达
	// p := uintptr(unsafe.Pointer(&x)) //uintptr不可达，x dead

	_ = p

	runtime.GC()

	runtime.KeepAlive(p) //确保p指针活着,这样gc完成后，可以观察p指针是否可以让x保持可达状态
}

//备注：使用最新版本的go，以前的版本可能无法回收stack frame内的变量

// go build -o main -gcflags "-N -l" &&  ./main
// x dead

//结论：
//普通指针和unsafe.Pointer指针对于GC而言都是直接引用，如果parent是黑色，那么所指向的目标也是黑色；
//而uintptr实际上不是指针，只是一种保存地址的整数类型，所以不构成引用
