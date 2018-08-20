package main

// go build -o main -gcflags -N
// go tool objdump -S -s "main\.main" main
// go tool compile -N -l -S main.go > main.s

type Xer interface {
	A()
	B()
}

type X int

// type X struct {
// 	a int
// 	_ [100]byte
// }

//go:noinline
func test() int {
	c := 0
	for i := 0; i < 10000; i++ {
		c++
	}
	return c
}

func (x X) A() {
	println("a")
	// test()
}

func (x X) B() {
	// test()
	println("b")
}

//go:noinline
//go:nosplit
func makeiface() Xer {
	var x X = 100
	return x
}

func main() {
	x := makeiface()
	x.B() //(*X).B(x.data)

	//var x X=100
	//x.B()
}
