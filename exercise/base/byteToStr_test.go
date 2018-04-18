package base

import (
	"fmt"
	"testing"
	"unsafe"
)

func Test_byte(t *testing.T) {
	str2 := "aAhello"
	data2 := []byte(str2)
	fmt.Println(data2)
	str2 = string(data2[:])
	fmt.Println(str2)

	upBytes()
	upBytes2()
}

func upBytes() {
	bytes := []byte("I am byte array !")
	println(&bytes)
	str := string(bytes)
	println(&str)
	bytes[0] = 'i' //注意这一行，bytes在这里修改了数据，但是str打印出来的依然没变化，
	fmt.Println(str)
}

func upBytes2() {
	bytes := []byte("I am byte array !")
	//这样做的意义在于，在网络通信中，大多数的接受方式都是[]byte，如果[]byte的数据比较大，内存拷贝的话会影响系统的性能。
	str := (*string)(unsafe.Pointer(&bytes))
	bytes[0] = 'i' //str和bytes共用一片内存
	println(&bytes)
	println(&str)
	fmt.Println(*str)
}
