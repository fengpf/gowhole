package sort

import (
	"fmt"
	"testing"
	"time"
)

func selection(arr []int) {
	l := len(arr)
	if l <= 1 {
		return
	}
	for i := 0; i < l; i++ { //有序区域 R[0..i]
		min := i
		for j := i + 1; j < l-1; j++ { //无序区域 R[i+1..n-1] 查找
			if arr[min] > arr[j] { //选出最小的
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i] //有序区域增加最小元素
	}
}

func Test_selectionSort(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1, 9, 6}
	start := time.Now()
	selection(arr)
	elapsed := time.Since(start)
	fmt.Println(elapsed)
	fmt.Println(arr)
}
