package basictests

import (
	"fmt"
	"testing"
)

//golang 的返回值是通过栈空间，不是通过寄存器，这点最重要。
//调用函数前，首先分配的是返回值空间，然后是入参地址，再是其他临时变量地址。
//return 操作  1、将返回值拷贝到栈空间第一块区域 2、判断defer函数是否修改栈空间的返回值 3、空的return（ret 跳转）
func Test_defer(t *testing.T) {
	fmt.Println(f(), f2(), f3())
}

func f() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}
