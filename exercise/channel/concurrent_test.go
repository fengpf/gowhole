package channel

import (
	"fmt"
	"sort"
	"sync"
	"testing"
)

func Test_slice(t *testing.T) {
	a := make([]int, 0, 10)
	aCh := make(chan int, 10)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			aCh <- i + 1
		}(i)
	}
	wg.Wait()
	close(aCh)
	// fmt.Println(<-aCh)
	for c := range aCh {
		a = append(a, c)
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})
	fmt.Println(a, <-aCh, <-aCh)
}
