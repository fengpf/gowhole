package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	// parallel starts parallel image processing based on the current GOMAXPROCS value.
	// If GOMAXPROCS = 1 it uses no parallelization.
	// If GOMAXPROCS > 1 it spawns N=GOMAXPROCS workers in separate goroutines.

	width := 100
	parallel(width, func(partStart, partEnd int) {
		for x := partStart; x < partEnd; x++ {
			fmt.Println(x)
		}
	})
}

func parallel(dataSize int, fn func(partStart, partEnd int)) {
	numGoroutines := 1
	partSize := dataSize
	numProcs := runtime.GOMAXPROCS(0)
	fmt.Printf("numOfProcs(%d)\n", numProcs)
	if numProcs > 1 {
		numGoroutines = numProcs
		partSize = dataSize / (numGoroutines * 10)
		if partSize < 1 {
			partSize = 1
		}
	}
	if numGoroutines == 1 {
		fn(0, dataSize)
	} else {
		fmt.Printf("numOfGoroutines(%d)\n", numGoroutines)
		var wg sync.WaitGroup
		wg.Add(numGoroutines)
		idx := uint64(0)
		for p := 0; p < numGoroutines; p++ {
			go func() {
				defer wg.Done()
				for {
					partStart := int(atomic.AddUint64(&idx, uint64(partSize))) - partSize
					if partStart >= dataSize {
						break
					}
					partEnd := partStart + partSize
					if partEnd > dataSize {
						partEnd = dataSize
					}
					fn(partStart, partEnd)
				}
			}()
		}
		wg.Wait()
	}
}
