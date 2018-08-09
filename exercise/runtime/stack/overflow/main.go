package main

import "fmt"

type val struct{}

//Stringer test
type Stringer interface {
	String() string
}

var i int

func (v val) String() string {
	// fmt.Println(v.String()) //递归 栈溢出
	i++
	println(i)
	return "aaa"
}

func main() {
	// var value interface{} = "aaa" // Value provided by caller.
	vv := val{}
	println(vv.String())
	var value = Stringer(vv)
	switch str := value.(type) {
	// case string:
	// 	fmt.Printf("str(%v) is string\n", str) //type of str is string
	case Stringer:
		fmt.Printf("str(%v) is Stringer\n", str.String()) //type of str is Stringer
	}
}
