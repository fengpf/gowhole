package main

import (
	"fmt"
	"io"
	"os"
	"strings"
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

	// 	Type switch与Type assertions
	// 在Go语言中，我们可以使用type switch语句查询接口变量的真实数据类型，语法如下：

	var x interface{} = "1" //x必须是接口类型
	switch x.(type) {
	// cases
	case string:
		fmt.Printf("x(%s) is string\n", x)
	}

	// var value interface{} = "aaa" // Value provided by caller.
	var value interface{}
	// value = "hello"
	value = Stringer(&val{str: "hello"})
	switch str := value.(type) {
	case string:
		fmt.Printf("str(%v) is string\n", str) //type of str is string
	case Stringer:
		fmt.Printf("str(%v) is Stringer\n", str.String()) //type of str is Stringer
	}

	// str := value.(string)
	// 上面的转换有一个问题，如果该值不包含一个字符串，则程序会产生一个运行时错误。为了避免这个问题，可以使用“comma, ok”的习惯用法来安全地测试值是否为一个字符串：
	str, ok := value.(string)
	if ok {
		fmt.Printf("string value is: %q\n", str)
	} else {
		fmt.Printf("value is not a string\n")
	}
	// 如果类型断言失败，则str将依然存在，并且类型为字符串，不过其为零值，即一个空字符串。
	// 我们可以使用类型断言来实现type switch的中例子：
	if str, ok := value.(string); ok {
		fmt.Printf("string value is: %v\n", str)
	} else if str, ok := value.(Stringer); ok {
		fmt.Printf("string value is: %v\n", str.String())
	}
}

type val struct {
	str string
}

//Stringer test
type Stringer interface {
	String() string
}

func (v *val) String() string {
	// println(v.String()) //递归
	return strings.ToUpper(v.str)
}
