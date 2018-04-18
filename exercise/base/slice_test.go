package base

import (
	"fmt"
	"testing"
)

func Test_Slice(t *testing.T) {
	a := make([]int, 0, 3)
	for i := 0; i < 5; i++ {
		a = append(a, i)
		if len(a) == 3 {
			break
		}
	}
	fmt.Println(a)
}
