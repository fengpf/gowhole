package basictests

import (
	"fmt"
	"testing"
	"unsafe"
)

func Float64bits(f float64) uint64 { return *(*uint64)(unsafe.Pointer(&f)) }

func Test_ptr(t *testing.T) {
	fmt.Printf("%#016x\n", Float64bits(1.0)) // "0x3ff0000000000000"

	var x struct {
		a bool
		b int16
		c []int
	}
	// 和 pb := &x.b 等价
	pb := (*int16)(unsafe.Pointer(
		uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	*pb = 42
	fmt.Println(x.b) // "42

	// NOTE: subtly incorrect!
	tmp := uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)
	pb2 := (*int16)(unsafe.Pointer(tmp))
	*pb2 = 43
	fmt.Println(x.b) // "43
}
