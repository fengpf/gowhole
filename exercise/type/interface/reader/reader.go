package main

import (
	"fmt"
)

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

type Seeker interface {
	Seek(offset int64, whence int) (int64, error)
}

type File struct {
	Reader
	Writer
	Closer
	Seeker
}

func (f *File) Read(buf []byte) (n int, err error) {
	return
}

func (f *File) Write(buf []byte) (n int, err error) {
	return
}
func (f *File) Seek(off int64, whence int) (pos int64, err error) {
	return
}

func (f *File) Close() (err error) {
	return
}

//接口嵌套
// ReadWriter is the interface that groups the basic Read and Write methods.
type ReadWriter interface {
	Reader
	Writer
}

// 该接口嵌套了io.Reader和io.Writer两个接口，实际上，它等同于下面的写法：
type ReadWriter2 interface {
	Read(p []byte) (n int, err error)
	Write(p []byte) (n int, err error)
}

// 注意，Go语言中的接口不能递归嵌套，
// illegal: Bad cannot embed itself
// type Bad interface {
// 	Bad
// }

// illegal: Bad1 cannot embed itself using Bad2
// type Bad1 interface {
// 	Bad2
// }

// type Bad2 interface {
// 	Bad1
// }

func main() {
	var file1 Reader = new(File)
	var file2 Writer = new(File)
	var file3 Closer = new(File)
	var file4 Seeker = new(File)

	fmt.Println(file1, file2, file3, file4, 1111)

	// 空接口（empty interface）
	// 空接口比较特殊，它不包含任何方法：

	// interface{}
	// 在Go语言中，所有其它数据类型都实现了空接口。

	var v1 interface{} = 1
	var v2 interface{} = "abc"
	var v3 interface{} = struct{ X int }{1}
	fmt.Println(v1, v2, v3, 1111)

	// []T不能直接赋值给[]interface{}
	t := []int{1, 2, 3, 4}
	// cannot use t (type []int) as type []interface {} in assignment
	// var s []interface{} = t
	s := make([]interface{}, len(t))
	for i, v := range t {
		s[i] = v
	}
	fmt.Println(s, t)

	// 	类型转换（Conversions）
	// 类型转换的语法：

	// Conversion = Type "(" Expression [ "," ] ")" .
	// 当以运算符*或者<-开始，有必要加上括号避免模糊：
	// *Point(p)        // same as *(Point(p))
	// (*Point)(p)      // p is converted to *Point
	// <-chan int(c)    // same as <-(chan int(c))
	// (<-chan int)(c)  // c is converted to <-chan int
	// func()(x)        // function signature func() x
	// (func())(x)      // x is converted to func()
	// (func() int)(x)  // x is converted to func() int
	// func() int(x)    // x is converted to func() int (unambiguous-不含糊的)

	// 	Type switch与Type assertions
	// 在Go语言中，我们可以使用type switch语句查询接口变量的真实数据类型，语法如下：

	var x interface{} = "1" //x必须是接口类型
	switch x.(type) {
	// cases
	case string:
		fmt.Printf("x(%s) is string\n", x)
	}

	// var value interface{} = "aaa" // Value provided by caller.
	vv := &val{}
	var value = Stringer(vv)
	switch str := value.(type) {
	// case string:
	// 	fmt.Printf("str(%v) is string\n", str) //type of str is string
	case Stringer:
		fmt.Printf("str(%v) is Stringer\n", str.String()) //type of str is Stringer
	}
}

type val struct{}

//Stringer test
type Stringer interface {
	String() string
}

func (v *val) String() string {
	println(v.String() + "aaa")
	return v.String()
}
