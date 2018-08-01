package sort

import (
	"fmt"
	"testing"
	"time"
)

func bubble(arr []int) {
	c := len(arr)
	if c == 0 || c == 1 {
		return
	}
	for i := 0; i < len(arr); i++ { //控制a[n]
		// fmt.Printf("%d trip\n", i)
		for j := 0; j < len(arr)-i-1; j++ { //a[0] 和 a[n-1]比较
			// fmt.Printf(" %d round\n ", j)
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
			// spew.Dump(arr)
		}
	}
}

func Test_bubbleSort(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1, 9, 6}
	start := time.Now()
	bubble(arr)
	elapsed := time.Since(start)
	fmt.Println(elapsed)
	fmt.Println(arr)
}
