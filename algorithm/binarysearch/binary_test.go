package search

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

// 普通二分查找
func binNormalSearch(data []int, key int) (mid int) {
	var (
		start = 0
		end   = len(data) - 1
	)
	for start <= end {
		mid = (start + end) >> 1
		if data[mid] > key {
			end = mid - 1
		} else if data[mid] < key {
			start = mid + 1
		} else {
			return
		}
	}
	return
}

//查找关键字第一次出现的位置 TODO check wrong
func binFirstSearch(data []int, key int) (mid int) {
	var (
		start = 0
		end   = len(data) - 1
	)
	for start < end {
		mid = (start + end) >> 1
		if data[mid] > key {
			end = mid - 1
		} else if data[mid] < key {
			start = mid + 1
		} else {
			end = mid
		}
	}
	if data[start] == key {
		return start
	}
	return
}

// 查找关键字最后一次出现的位置
func binLastSearch(data []int, key int) (mid int) {
	var (
		start = 0
		end   = len(data) - 1
	)
	for start < end-1 {
		mid = (start + end) >> 1
		if data[mid] > key {
			end = mid - 1
		} else if data[mid] < key {
			start = mid + 1
		} else {
			start = mid
		}
	}
	if data[start] == key {
		return start
	}
	if data[end] == key {
		return end
	}
	return
}

// 查找小于关键字的最大数字出现的位置
func binMaxLessSearch(data []int, key int) (mid int) {
	var (
		start = 0
		end   = len(data) - 1
	)
	for start < end-1 {
		mid = (start + end) >> 1
		if data[mid] > key {
			end = mid - 1
		} else if data[mid] < key {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	if data[start] < key {
		return start
	}
	if data[end] < key {
		return end
	}
	return
}

// 查找大于关键字的最小数字出现的位置
func binMinMoreSearch(data []int, key int) (mid int) {
	var (
		start = 0
		end   = len(data) - 1
	)
	for start < end-1 {
		mid = (start + end) >> 1
		if data[mid] > key {
			end = mid
		} else if data[mid] < key {
			start = mid + 1
		} else {
			end = mid + 1
		}
	}
	if data[start] > key {
		return start
	}
	return
}

func Test_binarySearch(t *testing.T) {
	data := []int{18, 16, 17, 1, 3, 5, 7, 9, 8, 8, 15}

	ni := binNormalSearch(data, 8)
	fmt.Printf("normal index(%d) digit(%d) \n", ni, data[ni])

	fi := binFirstSearch(data, 8)
	fmt.Printf("first index(%d) digit(%d)\n", fi, data[fi])

	li := binLastSearch(data, 8)
	fmt.Printf("last index(%d) digit(%d)\n", li, data[li])

	mi := binMaxLessSearch(data, 8)
	fmt.Printf("maxless index(%d) digit(%d)\n", mi, data[mi])

	mmi := binMinMoreSearch(data, 8)
	fmt.Printf("minmore index(%d) digit(%d)\n", mmi, data[mmi])
}

var mids = "268104,389088,1724598,4931952,8820267"

func TestFindArr(t *testing.T) {
	mid := "8820267"
	midArr := strings.Split(mids, ",")
	now := time.Now()
	// res := orderSearch(midArr, mid)
	res := binarySearch(midArr, mid)
	println(res)
	elapsed := time.Since(now)
	fmt.Println(elapsed)
}

func orderSearch(s []string, e string) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func binarySearch(s []string, e string) bool {
	var low, mid, high int
	high = len(s) - 1
	for low <= high {
		mid = (low + high) / 2
		if s[mid] == e {
			return true
		}
		if s[mid] < e {
			low = mid + 1
		}
		if s[mid] > e {
			high = mid - 1
		}
		println("mid", mid, s[mid], "low", low, "high", high)
	}
	return false
}
