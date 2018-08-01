package sort

import (
	"fmt"
	"testing"
	"time"
)

func quickSort(arr []int, l, r int) {
	c := len(arr)
	if c == 0 || c == 1 {
		return
	}
	k := arr[l]
	p := l
	i, j := l, r
	for i <= j {
		for j >= p && arr[j] >= k {
			j--
		}
		if j >= p {
			arr[p] = arr[j]
			p = j
		}
		if arr[i] <= k && i <= p {
			i++
		}
		if i < p {
			arr[p] = arr[i]
			p = i
		}
		arr[p] = k
		// spew.Dump(arr)
		if p-l > 1 {
			quickSort(arr, l, p-1)
		}
		if r-p > 1 {
			quickSort(arr, p+1, r)
		}
	}
}

func Test_quickSort(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1, 9, 6}
	start := time.Now()
	quickSort(arr, 0, len(arr)-1)
	elapsed := time.Since(start)
	fmt.Println(elapsed)
	fmt.Println(arr)
}
