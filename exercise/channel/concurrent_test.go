package channel

import (
	"fmt"
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
	fmt.Println(<-aCh)
	close(aCh)
	for c := range aCh {
		a = append(a, c)
	}
	fmt.Println(a, <-aCh, <-aCh)
}
