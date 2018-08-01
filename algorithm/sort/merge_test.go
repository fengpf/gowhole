package sort

import (
	"fmt"
	"testing"
	"time"
)

func merge(left, right []int) (res []int) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			res = append(res, left[l])
			l++
		} else {
			res = append(res, right[r])
			r++
		}
	}
	res = append(res, left[l:]...)
	res = append(res, right[r:]...)
	return
}

func mergeSort(A []int) []int { //对A[left..right]进行归并排序
	l := len(A)
	if l < 2 {
		return A
	}
	middle := l / 2 //将A[left..right]分成两个子序列进行递归归并排序
	left := mergeSort(A[:middle])
	right := mergeSort(A[middle:])
	return merge(left, right) //将已排序的两个子序列进行合并
}

func Test_mergeSort(t *testing.T) {
	A := []int{1, 2, 3, 854, 9, 55, 46, 4, 5, 45, 4, 41, 4, 5, 555}
	start := time.Now()
	B := mergeSort(A)
	elapsed := time.Since(start)
	fmt.Println(elapsed)
	fmt.Println(B)
}
