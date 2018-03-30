package sort

import (
	"fmt"
	"testing"
	"time"
)

func insertion(arr []int) { //从后往前插入
	l := len(arr)
	if l <= 1 {
		return
	}
	for i := 1; i < l; i++ { //从第一个元素开始，该元素可以认为已经被排序
		cur := arr[i]
		prev := i - 1 //在已经排序的元素序列中从后向前扫描，取出下一个元素
		// fmt.Println("1111", prev, cur)
		for prev >= 0 && cur < arr[prev] { //如果该元素（已排序）小于新元素，将该元素往前移到下一位置
			// fmt.Println(prev, arr[prev+1], arr[prev])
			arr[prev+1] = arr[prev] //如果cur比arr[prev]小，则arr[prev]往后不停移动
			prev--
		}
		//fmt.Println("2222", prev, cur)
		//直到找到已排序的元素大于或者等于新元素的位置,将新元素插入到该位置后
		arr[prev+1] = cur
	}
}

func insertion2(arr []int) { //从前往后插入
	l := len(arr)
	if l <= 1 {
		return
	}
	for i := l - 2; i >= 0; i-- {
		cur := arr[i]
		prev := i + 1
		for prev <= l-1 && cur < arr[prev] {
			arr[prev-1] = arr[prev]
			prev++
		}
		arr[prev-1] = cur
	}
}

func Test_insertionSort(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1, 9, 6}
	start := time.Now()
	insertion2(arr)
	elapsed := time.Since(start)
	fmt.Println(elapsed)
	fmt.Println(arr)
}
