package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var x int = 100

	var a interface{} = x
	var b interface{} = &x

	fmt.Println(unsafe.Sizeof(a), unsafe.Sizeof(b))
}
