package tests

import (
	"fmt"
	"math"
	"math/cmplx"
	"testing"
)

var (
	aa = 3
	ss = "ss"
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s = "abc"
	fmt.Println(a, b, s)

}

func variableTypeDdeuction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 6
	fmt.Println(a, b, c, s)
}

func euler() {
	// c := 3 + 4i
	// fmt.Println(cmplx.Abs(c))
	b := cmplx.Pow(math.E, 1i*math.Pi) + 1
	d := cmplx.Exp(1i*math.Pi) + 1
	fmt.Println(b, d)
	fmt.Printf("%.3f\n", b)

}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func consts() {
	const filename = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	const (
		cpp = iota
		_
		java
		python
		golang
		javascript
	)
	//b,kb,mb,gb,td,pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, java, python, golang, javascript)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func increByVal(a int) int {
	println(&a)
	a = a + 1
	return a
}

func increByRef(a *int) int {
	println(a)
	*a = *a + 1
	return *a
}

func Test_all(t *testing.T) {
	fmt.Println("Hello World")
	variableZeroValue()
	variableInitialValue()
	variableTypeDdeuction()
	variableShorter()
	fmt.Println(aa, ss)

	euler()
	triangle()
	consts()
	enums()

	a := 1
	b := 1
	println(increByVal(a), a)
	println(increByRef(&b), b)
}
