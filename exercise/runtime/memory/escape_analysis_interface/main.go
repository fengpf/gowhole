package main

import (
	"fmt"
	"unsafe"
)

// LEA：操作地址；
// MOV：操作数据

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
	var (
		c = new(struct{})
		d = new([0]int)
	)
	// go 不支持指针运算和转型，只要有个合法的地址就ok，反正0长度不能赋值(忽略)和越界。unsafe除外
	//如果是堆上分配，所有0长度都是指向一个全局变量，共享同一个内存地址，就算类型相同也一样

	// fmt.Printf("%p\t%p\n", &a, &b) //使用fmt.Printf,地址不同，原因是参数是interface,会导致堆上分配，所以会导致ab地址不同
	fmt.Printf("%p\t%p\n", &c, &d) //确保在堆上分配
}

//go build  -o  main -gcflags '-l -m' main.go
//go tool objdump -s "main\.main" main
