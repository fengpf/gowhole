package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"unsafe"
)

// 指针类型：

// *类型：普通指针，用于传递对象地址，不能进行指针运算。

// unsafe.Pointer：unsafe.Pointer其实就是类似C的void*，用于各种指针相互转换的桥梁， 通用指针类型，用于转换不同类型的指针，不能进行指针运算。

// uintptr：是golang的内置类型，是能存储指针的整型，uintptr的底层类型是int，参与指针运算，GC 不把 uintptr 当指针，uintptr 无法持有对象。uintptr 类型的目标会被回收。

// 　　unsafe.Pointer 可以和 普通指针 进行相互转换。
// 　　unsafe.Pointer 可以和 uintptr 进行相互转换。

// 　　也就是说 unsafe.Pointer 是桥梁，可以让任意类型的指针实现相互转换，也可以将任意类型的指针转换为 uintptr 进行指针运算。

// ------------------------------

// 示例：通过指针修改结构体字段

func Test_unsafe(t *testing.T) {
	s := struct {
		a byte
		b byte
		c byte
		d int64
	}{0, 0, 0, 0}

	// 将结构体指针转换为通用指针
	p := unsafe.Pointer(&s)
	// 保存结构体的地址备用（偏移量为 0）
	up0 := uintptr(p)
	// 将通用指针转换为 byte 型指针
	pb := (*byte)(p)
	// 给转换后的指针赋值
	*pb = 10
	// 结构体内容跟着改变
	fmt.Println(s)

	// 偏移到第 2 个字段
	up := up0 + unsafe.Offsetof(s.b)
	// 将偏移后的地址转换为通用指针
	p = unsafe.Pointer(up)
	// 将通用指针转换为 byte 型指针
	pb = (*byte)(p)
	// 给转换后的指针赋值
	*pb = 20
	// 结构体内容跟着改变
	fmt.Println(s)

	// 偏移到第 3 个字段
	up = up0 + unsafe.Offsetof(s.c)
	// 将偏移后的地址转换为通用指针
	p = unsafe.Pointer(up)
	// 将通用指针转换为 byte 型指针
	pb = (*byte)(p)
	// 给转换后的指针赋值
	*pb = 30
	// 结构体内容跟着改变
	fmt.Println(s)

	// 偏移到第 4 个字段
	up = up0 + unsafe.Offsetof(s.d)
	// 将偏移后的地址转换为通用指针
	p = unsafe.Pointer(up)
	// 将通用指针转换为 int64 型指针
	pi := (*int64)(p)
	// 给转换后的指针赋值
	*pi = 40
	// 结构体内容跟着改变
	fmt.Println(s)
}

// ------------------------------

// 修改其它包中的结构体私有字段：

// ------------------------------

func Test_modify(t *testing.T) {
	// 创建一个 strings 包中的 Reader 对象
	// 它有三个私有字段：s string、i int64、prevRune int
	sr := strings.NewReader("abcdef")
	// 此时 sr 中的成员是无法修改的
	fmt.Println(sr)
	// 但是我们可以通过 unsafe 来进行修改
	// 先将其转换为通用指针
	p := unsafe.Pointer(sr)
	// 获取结构体地址
	up0 := uintptr(p)
	// 确定要修改的字段（这里不能用 unsafe.Offsetof 获取偏移量，因为是私有字段）
	if sf, ok := reflect.TypeOf(*sr).FieldByName("i"); ok {
		// 偏移到指定字段的地址
		up := up0 + sf.Offset
		// 转换为通用指针
		p = unsafe.Pointer(up)
		// 转换为相应类型的指针
		pi := (*int64)(p)
		// 对指针所指向的内容进行修改
		*pi = 3 // 修改索引
	}
	// 看看修改结果
	fmt.Println(sr)
	// 看看读出的是什么
	b, err := sr.ReadByte()
	fmt.Printf("%c, %v\n", b, err)
}

// ------------------------------

// 　　另外还有一种简单的方法，不用考虑偏移量的问题：

// ------------------------------

// 定义一个和 strings 包中的 Reader 相同的本地结构体
type Reader struct {
	s        string
	i        int64
	prevRune int
}

func Test_Reader(t *testing.T) {
	// 创建一个 strings 包中的 Reader 对象
	sr := strings.NewReader("abcdef")
	// 此时 sr 中的成员是无法修改的
	fmt.Println(sr)
	// 我们可以通过 unsafe 来进行修改
	// 先将其转换为通用指针
	p := unsafe.Pointer(sr)
	// 再转换为本地 Reader 结构体
	pR := (*Reader)(p)
	// 这样就可以自由修改 sr 中的私有成员了
	(*pR).i = 3 // 修改索引
	// 看看修改结果
	fmt.Println(sr)
	// 看看读出的是什么
	b, err := sr.ReadByte()
	fmt.Printf("%c, %v\n", b, err)
}

func Test_ptr(t *testing.T) {
	fmt.Printf("%#016x\n", Float64bits(1.0)) // "0x3ff0000000000000"

	var x struct {
		a bool
		b int16
		c []int
	}
	// 和 pb := &x.b 等价
	pb := (*int16)(unsafe.Pointer(
		uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	*pb = 42
	fmt.Println(x.b) // "42

	// NOTE: subtly incorrect!
	tmp := uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)
	pb2 := (*int16)(unsafe.Pointer(tmp))
	*pb2 = 43
	fmt.Println(x.b) // "43
}
