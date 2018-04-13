package main

import "sync"

func QuickSort2(xs []int, wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(1)
	go QuickSortHelp(xs, 0, len(xs)-1, wg)
}

func QuickSortHelp(xs []int, p, r int, wg *sync.WaitGroup) {
	defer wg.Done()
	if p >= r {
		return
	}
	q := partition(xs, p, r)
	wg.Add(2)
	go QuickSortHelp(xs, p, q-1, wg)
	go QuickSortHelp(xs, q+1, r, wg)
}

func partition(xs []int, p, r int) int {
	e := xs[r]
	i := p - 1
	for j := p; j < r; j++ {
		if xs[j] <= e {
			i++
			swap(&(xs[i]), &(xs[j]))
		}
	}
	swap(&(xs[i+1]), &(xs[r]))
	return i + 1
}

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
