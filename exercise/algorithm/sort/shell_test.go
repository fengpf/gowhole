package sort

import (
	"fmt"
	"testing"
	"time"
)

func shell(arr []int) {

}

func Test_shellSort(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1, 9, 6}
	start := time.Now()
	shell(arr)
	elapsed := time.Since(start)
	fmt.Println(elapsed)
	fmt.Println(arr)
}
