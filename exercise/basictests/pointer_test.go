package basictests

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

type MyStruct struct {
	A int
	B int
}

var sizeOfMyStruct = int(unsafe.Sizeof(MyStruct{}))

func MyStructToBytes(s *MyStruct) []byte {
	var x reflect.SliceHeader
	x.Len = sizeOfMyStruct
	x.Cap = sizeOfMyStruct
	x.Data = uintptr(unsafe.Pointer(s))
	return *(*[]byte)(unsafe.Pointer(&x))
}

func BytesToMyStruct(b []byte) *MyStruct {
	return (*MyStruct)(unsafe.Pointer(
		(*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
	))
}

func Test_bytesToMyStruct(t *testing.T) {
	b := MyStructToBytes(&MyStruct{A: 256, B: 258})
	fmt.Println(b, len(b))

	ms := BytesToMyStruct(b)
	fmt.Println(ms)
}

// returns &s[0], which is not allowed in go
func stringPointer(s string) unsafe.Pointer {
	p := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return unsafe.Pointer(p.Data)
}

// returns &b[0], which is not allowed in go
func bytePointer(b []byte) unsafe.Pointer {
	p := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	return unsafe.Pointer(p.Data)
}

func Test_ByteString(t *testing.T) {
	var x = []byte("Hello World!")
	var y = *(*string)(unsafe.Pointer(&x))
	var z = string(x)

	if y != z {
		t.Fail()
	}
	fmt.Println(y, z, reflect.TypeOf(y), reflect.TypeOf(z))
}

func Benchmark_Normal(b *testing.B) {
	var x = []byte("Hello World!")
	for i := 0; i < b.N; i++ {
		_ = string(x)
	}
}

func Benchmark_ByteString(b *testing.B) {
	var x = []byte("Hello World!")
	for i := 0; i < b.N; i++ {
		_ = *(*string)(unsafe.Pointer(&x))
	}
}
