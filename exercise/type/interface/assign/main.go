package main

import (
	"io"
	"os"
)

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}
func (a *Integer) Add(b Integer) {
	*a += b
}

type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
}

// （1）通过对象实例赋值
var a Integer = 1
var b1 LessAdder = &a //OK
// The method set of any other named type T consists of all methods with receiver type T.
// The method set of the corresponding pointer type T is the set of all methods with receiver T or T (that is, it also contains the method set of T).
// 也就是说*Integer实现了接口LessAdder的所有方法，而Integer只实现了Less方法，所以不能赋值。
// var b2 LessAdder = a  //not OK

// （2）通过接口赋值
var r io.Reader = new(os.File)

// 因为r没有Write方法，所以不能赋值给rw
// var rw io.ReadWriter = r //not ok

var rw2 io.ReadWriter = new(os.File)
var r2 io.Reader = rw2 //ok

func main() {

}
