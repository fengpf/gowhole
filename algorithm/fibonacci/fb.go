package main

import (
	"fmt"
	"time"
)

const N = 20

var fbs [N]uint64

func main() {

	start := time.Now()

	// fmt.Printf("fb(%d)\n", fb(N)) //elapsed(83.703µs)

	for i := 0; i < N; i++ {
		// fmt.Printf("fb(%d)\n", fb(i)) //elapsed(171.264µs)
		fmt.Printf("fb2(%d)\n", fb2(i)) //elapsed(57.846µs)
	}

	elapsed := time.Since(start)

	fmt.Printf("elapsed(%v)\n", elapsed)
}

func fb(n int) (res int) {
	if n <= 1 {
		res = 1
		return
	}

	res = fb(n-1) + fb(n-2)
	return
}

func fb2(n int) (res uint64) {
	if fbs[n] != 0 {
		res = fbs[n]
		return
	}

	if n <= 1 {
		res = 1
	} else {
		res = fb2(n-1) + fb2(n-2)
	}

	fbs[n] = res
	return
}
