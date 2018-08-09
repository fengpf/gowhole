package main

import (
	"fmt"
	"unsafe"
)

func main() {
	test := struct{}{}
	fmt.Println(unsafe.Sizeof(test)) //0

	var in interface{}
	fmt.Println(unsafe.Sizeof(in)) //16 由于接口自身数据结构是由两个指针构成

	//有个例外，就是0长度作为结构体的最后一个成员
	var x struct {
		a int
		b struct{}
	}
	fmt.Println(unsafe.Sizeof(x)) //16

	// go 不支持指针运算和转型，只要有个合法的地址就ok，反正0长度不能赋值(忽略)和越界。unsafe除外
	//如果是堆上分配，所有0长度都是指向一个全局变量，共享同一个内存地址，就算类型相同也一样

	var (
		d = new([0]int)
		a struct{}
		c = new(struct{})
		b = 1
	)
	_ = a
	_ = b
	_ = c
	_ = d
	println(&a, &b) //0xc42004bef0 0xc42004bef0  相同地址
	println(&c, &d) //0xc42004bf00 0xc42004bef8  不同地址
	fmt.Println(
		unsafe.Sizeof(a), //0
		unsafe.Sizeof(b), //8
		unsafe.Sizeof(c), //8
		unsafe.Sizeof(d), //8
	)

	// fmt.Printf("%p\t%p\n", &a, &b) //使用fmt.Printf,地址不同，原因是参数是interface,会导致堆上分配，所以会导致ab地址不同
	// fmt.Printf("%p\t%p\n", &c, &d) //确保在堆上分配
}

//go build -gcflags '-l' -o  main main.go
//go tool objdump -s "main\.main" main
