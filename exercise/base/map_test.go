package base

import (
	"fmt"
	"testing"
)

func Test_map(t *testing.T) {
	var (
		a, res map[int]int
		i      int
		j      int
	)
	fmt.Println(len(a), a == nil, a)
	a = make(map[int]int)
	a[1] = 1
	fmt.Println(len(a), a != nil)

	res = rank(a)
	fmt.Printf("res %p\n", res)
	fmt.Println(len(res), res != nil, res)

	i = 1
	j = rank2(i)
	fmt.Printf("j %p\n", &j)
}

func rank(b map[int]int) (c map[int]int) {
	c = make(map[int]int)
	fmt.Printf("c %p\n", c)
	// for k, v := range b {
	// 	c[k] = v
	// }
	return
}

func rank2(b int) (c int) {
	c = b
	fmt.Printf("c %p\n", &c)
	return
}
