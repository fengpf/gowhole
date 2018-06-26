package basictests

import (
	"fmt"
	"testing"
)

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
