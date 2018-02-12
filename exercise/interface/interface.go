package main

import (
	"fmt"
	"strconv"
)

type Stringer interface {
	String() string
}

type Binary uint64

func (i Binary) String() string {
	return strconv.FormatUint(i.Get(), 2)
}

func (i Binary) Get() uint64 {
	return uint64(i)
}

func main() {
	var b Binary = 32
	//首次遇见s := Stringer(b)这样的语句时，golang会生成Stringer接口对应于Binary类型的虚表，并将其缓存
	s := Stringer(b)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", b.Get())
	emptyInterface()

	c := 100
	d := &c
	*d++
	fmt.Printf("c pointer(%p), value(%v) \n\r", &c, c)
	fmt.Printf("d pointer(%p), value(%v) pointerValue(%d)\n\r", &d, d, *d)
}

func emptyInterface() {
	//接口类型的一个极端重要的例子是空接口：interface{},它表示空的方法集合，
	//由于任何值都有零个或者多个方法，所以任何值都可以满足它。 注意，[]T不能直接赋值给[]interface{}
	t := []int{1, 2, 3, 4}
	// var s []interface{} = t  //wrong
	s := make([]interface{}, len(t))
	for i, v := range t {
		s[i] = v
	}

	var value interface{}
	fmt.Printf("value is: %+v\n", value)
	switch str := value.(type) {
	case string:
		fmt.Printf("string value is: %q\n", str)
	case Stringer:
		fmt.Printf("value is not a string\n")
	}
}
