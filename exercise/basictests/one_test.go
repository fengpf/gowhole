package basictests

import (
	"fmt"
	"testing"
	"unsafe"
)

const (
	_        = iota             //iota=0
	KB int64 = 1 << (10 * iota) //iota =1
	MB                          //与KB表达式相同，但iota=2
	GB
	TB
)
const (
	A, B = iota, iota << 10 //0, 0<<10
	C, D                    //1,1<<10
)
const (
	AA = iota //0
	BB        //1
	CC = "c"  // c
	DD        //c，与上一行相同
	EE
)

func Test_Def(t *testing.T) {
	// var x int
	// var f float32 = 1.6
	// var s = "abc"
	// println(x, f, s)
	// test()
	// var i, j, k int
	// var m, n = "abc", 123
	// var (
	// 	a int
	// 	b float32
	// )
	// q, w := 0x1234, "hello world"
	// println(i, j, k, a, b, m, n, q, w)
	data, i := [3]int{0, 1, 2}, 0
	i, data[i] = 2, 100
	const x, y int = 1, 2
	const s = "Hello World"
	const (
		a = "abc"
		b = len(a)
		c = unsafe.Sizeof(b)
	)
	const (
		d byte = 1 // int to byte
		//e int  = 1e20 // float64 to int, overflows
	)
	const (
		Sunday    = iota //0
		Monday           // 1 通常省略后续行表达式
		Tuesday          //2
		Wednesday        //3
		Thursday         //4
		Friday           //5
		Saturday         //6
	)
}

type Color int

const (
	Black Color = iota
	Red
	Blue
)

func test(c Color) {
	c = Black
	test(c)
	//x := 1
	//test(x)
	test(1)
}
func Test_type(t *testing.T) {
	a := []int{0, 0, 0}
	a[1] = 10
	b := make([]int, 3)
	b[1] = 10
	//c := new([]int)
	//c(1) = 10
	var m byte = 100
	var n = int(m)
	println(n)
	// *Point(p) //相当于*(Point(p))
	//(*Point)(p)
	//<-chan int(c) //相当于 <-(chan int(c))
	// (<-chan int)(c)
	s := "abc"
	println(s[0] == '\x61', s[1] == 'b', s[2] == 0x63)
	var s0, s1 rune = '\u6211', '们'
	println(s0 == '我', string(s1) == "\xe4\xbb\xac")
	str := "Hello" +
		"World!"
	println(str)
	s2 := "Hello, World!"
	s3 := s2[:5]
	s4 := s2[7:]
	s5 := s2[1:7]
	println(s3)
	println(s4)
	println(s5)
	s6 := "abcd"
	bs := []byte(s6)
	bs[1] = 'B'
	println(string(bs))
	u := "电脑"
	us := []rune(u)
	us[1] = '话'
	println(string(us))
	s7 := "abc汉字"
	for i := 0; i < len(s7); i++ {
		fmt.Printf("%c,", s7[i])
	}
	fmt.Println()
	for _, r := range s7 {
		fmt.Printf("%c,", r)
	}
	fmt.Println()
}

func Test_Ptr(t *testing.T) {
	type data struct{ a int }
	var d = data{1234}
	var p *data
	p = &d
	fmt.Printf("%p,%v\n", p, p.a)
	x := 0x12345678
	p1 := unsafe.Pointer(&x) //*int->Pointer
	n := (*[4]byte)(p1)
	for i := 0; i < len(n); i++ {
		fmt.Printf("%X ", n[i])
	}
	fmt.Println()
	testHeap()
	d2 := struct {
		s string
		x int
		y string
	}{"abc", 100, "cf"}
	p2 := uintptr(unsafe.Pointer(&d2)) //*struct->Pointer->uintptr
	fmt.Printf("%x\n", p2)
	fmt.Printf("%x\n", unsafe.Offsetof(d2.s))
	fmt.Printf("%x\n", unsafe.Offsetof(d2.x))
	fmt.Printf("%x\n", unsafe.Offsetof(d2.y))
	p2 += unsafe.Offsetof(d2.s) //uintptr + offset
	fmt.Printf("%x\n", p2)
	p2 += unsafe.Offsetof(d2.x) //uintptr + offset
	fmt.Printf("%x\n", p2)
	p2 += unsafe.Offsetof(d2.y) //uintptr + offset
	fmt.Printf("%x\n", p2)
	p3 := unsafe.Pointer(&p2)
	px := (*int)(p3)
	*px = 200
	fmt.Printf("%#v\n", d)
}

func testHeap() *int {
	x := 10
	return &x // 在堆上分配 x 内存。但在内联时，也可能直接分配在⺫标栈。
}
