package basictests

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var (
	qps   = time.Microsecond * 5
	total int64
)

func Test_serial(t *testing.T) {
	now := time.Now()
	for i := 0; i < 20000; i++ {
		time.Sleep(qps)
	}
	fmt.Println("gap:", time.Since(now))
}

func Test_serial_avg(t *testing.T) {
	for i := 0; i < 20000; i++ {
		now := time.Now()
		time.Sleep(qps)
		total += int64(time.Since(now)) / int64(time.Microsecond)
	}
	fmt.Println("avg:", total/20000)
}

func Test_goroutine(t *testing.T) {
	var wg sync.WaitGroup
	now := time.Now()
	for i := 0; i < 20000; i++ {
		time.Sleep(qps)
		wg.Add(1)
		go func() {
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("gap:", time.Since(now))
}

func Test_wgGroup(t *testing.T) {
	aids := []int{1, 2, 3, 4, 5}
	b := make([]int, 0, 5)
	var (
		wg sync.WaitGroup
		l  sync.Mutex
	)
	for _, id := range aids {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			l.Lock()
			b = append(b, id)
			l.Unlock()
		}(id)
	}
	wg.Wait()
	fmt.Println(b)
}
