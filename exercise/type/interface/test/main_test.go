package main

import "testing"

var x X
var xr Xer = x

func BenchmarkCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x.B()
	}
}

func BenchmarkIface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xr.B()
	}
}

type I1 interface {
	Method1()
}

type I2 interface {
	Method1()
	Method2()
}

type TS uint16
type TM uintptr
type TL [2]uintptr

var (
	e  interface{}
	e_ interface{}
	i1 I1
	i2 I2
	ts TS
	tm TM
	tl TL
	ok bool
)

type T8 uint8
type T16 uint16
type T32 uint32
type T64 uint64
type Tstr string
type Tslice []byte

func (T8) Method1()     {}
func (T16) Method1()    {}
func (T32) Method1()    {}
func (T64) Method1()    {}
func (Tstr) Method1()   {}
func (Tslice) Method1() {}

func TestZeroConvT2x(t *testing.T) {
	tests := []struct {
		name string
		fn   func()
	}{
		{name: "E8", fn: func() { e = eight8 }},  // any byte-sized value does not allocate
		{name: "E16", fn: func() { e = zero16 }}, // zero values do not allocate
		{name: "E32", fn: func() { e = zero32 }},
		{name: "E64", fn: func() { e = zero64 }},
		{name: "Estr", fn: func() { e = zerostr }},
		{name: "Eslice", fn: func() { e = zeroslice }},
		{name: "Econstflt", fn: func() { e = 99.0 }}, // constants do not allocate
		{name: "Econststr", fn: func() { e = "change" }},
		{name: "I8", fn: func() { i1 = eight8I }},
		{name: "I16", fn: func() { i1 = zero16I }},
		{name: "I32", fn: func() { i1 = zero32I }},
		{name: "I64", fn: func() { i1 = zero64I }},
		{name: "Istr", fn: func() { i1 = zerostrI }},
		{name: "Islice", fn: func() { i1 = zerosliceI }},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			n := testing.AllocsPerRun(1000, test.fn)
			if n != 0 {
				t.Errorf("want zero allocs, got %v", n)
			}
		})
	}
}

var (
	eight8  uint8 = 8
	eight8I T8    = 8

	zero16  uint16 = 0
	zero16I T16    = 0
	one16   uint16 = 1

	zero32  uint32 = 0
	zero32I T32    = 0
	one32   uint32 = 1

	zero64  uint64 = 0
	zero64I T64    = 0
	one64   uint64 = 1

	zerostr  string = ""
	zerostrI Tstr   = ""
	nzstr    string = "abc"

	zeroslice  []byte = nil
	zerosliceI Tslice = nil
	nzslice    []byte = []byte("abc")

	zerobig [512]byte
	nzbig   [512]byte = [512]byte{511: 1}
)
