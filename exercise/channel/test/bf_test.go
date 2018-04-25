package main

import (
	"fmt"
	"gowhole/exercise/channel/buffer"
	"sync"
	"testing"
	"time"
)

func main() {

}

func Test_set_get(t *testing.T) {
	for i := 0; i < 5; i++ {
		buffer.Set(i, i, 6000000000)
	}
	var wg sync.WaitGroup
	for j := 0; j < 5; j++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			println(i)
			fmt.Println(buffer.Get(i))
		}(j)
		time.Sleep(time.Duration(time.Second))
	}
	wg.Wait()
}
