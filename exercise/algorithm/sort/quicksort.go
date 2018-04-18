package sort

import (
	"sync"
)

func main() {
	// var sa SortArr
	// sa.Comparate()
}

// func QuickSort2(xs []int) {
// 	wg := &sync.WaitGroup{}
// 	wg.Add(1)
// 	go QuickSortHelp(xs, 0, len(xs)-1, wg)
// 	wg.Wait()
// }

// func QuickSortHelp(xs []int, p, r int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	if len(xs) < 2 {
// 		return
// 	}
// 	q := partition(xs, p, r)
// 	wg.Add(2)
// 	go QuickSortHelp(xs, p, q-1, wg)
// 	go QuickSortHelp(xs, q+1, r, wg)
// }

// func partition(xs []int, p, r int) int {
// 	e := xs[r]
// 	i := p - 1
// 	for j := p; j < r; j++ {
// 		if xs[j] <= e {
// 			i++
// 			swap(xs[i], xs[j])
// 		}
// 	}
// 	swap(xs[i+1], xs[r])
// 	return i + 1
// }

// func swap(a, b int) {
// 	temp := a
// 	a = b
// 	b = temp
// }

func QuickSort3(src []int, first, last int, wg *sync.WaitGroup) {
	defer wg.Done()
	flag := first
	left := first
	right := last
	if first >= last {
		return
	}
	for first < last {
		for first < last {
			if src[last] >= src[flag] {
				last--
				continue
			} else {
				tmp := src[last]
				src[last] = src[flag]
				src[flag] = tmp
				flag = last
				break
			}
		}
		for first < last {
			if src[first] <= src[flag] {
				first++
				continue
			} else {
				tmp := src[first]
				src[first] = src[flag]
				src[flag] = tmp
				flag = first
				break
			}
		}
	}
	wg.Add(2)
	go QuickSort3(src, left, flag-1, wg)
	go QuickSort3(src, flag+1, right, wg)
}
