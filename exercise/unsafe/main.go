package main

import (
	"fmt"
	"unsafe"
)

//	- A pointer value of any type can be converted to a Pointer.
//	- A Pointer can be converted to a pointer value of any type.
//	- A uintptr can be converted to a Pointer.
//	- A Pointer can be converted to a uintptr.

func main() {
	ss := make([]byte, 2)
	ss[0] = 1
	fmt.Println(ss)

	ptr := unsafe.Pointer(&ss[0])
	fmt.Println(&ss[0], ptr, (*byte)(ptr), *(*byte)(ptr))

	var a float64
	a = 5.0
	fmt.Println(a, Float64bits(a))

	fmt.Println("=====================")

	// If p points into an allocated object, it can be advanced through the object
	// by conversion to uintptr, addition of an offset, and conversion back to Pointer.
	//
	//	p = unsafe.Pointer(uintptr(p) + offset)
	//
	// The most common use of this pattern is to access fields in a struct
	// or elements of an array:
	//
	//	// equivalent to f := unsafe.Pointer(&s.f)

	type t struct {
		a int
		b string
	}

	var s t
	s.a = 1
	s.b = "b"

	// f := unsafe.Pointer(&s.a)
	// f := unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + unsafe.Offsetof(s.b))

	f := unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + unsafe.Offsetof(s.b))

	fmt.Println(
		unsafe.Pointer(&s),
		uintptr(unsafe.Pointer(&s)),
		unsafe.Offsetof(s.b),
		uintptr(unsafe.Pointer(&s))+unsafe.Offsetof(s.b),
		f,
		*(*string)(f),
	)

	//
	//	// equivalent to e := unsafe.Pointer(&x[i])
	//	e := unsafe.Pointer(uintptr(unsafe.Pointer(&x[0])) + i*unsafe.Sizeof(x[0]))
}

//Float64bits from *T1 to *T2
func Float64bits(f float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&f))
}
