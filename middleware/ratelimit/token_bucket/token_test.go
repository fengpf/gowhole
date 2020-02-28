package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestSimultaneousRequests(t *testing.T) {
	const (
		rate   = 1 //1s 1个
		bucket = 5 //桶容量

		numReq = 15 //请求数量
	)

	lim := NewLimiter(rate, bucket)

	var (
		wg    sync.WaitGroup
		numOK = uint32(0)
	)

	wg.Add(numReq)
	for i := 0; i < numReq; i++ {

		j := i
		go func() {
			if ok := lim.Take(time.Now(), 1, 0); ok {
				fmt.Println(j)
				atomic.AddUint32(&numOK, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	if numOK != bucket {
		t.Errorf("numOK = %d, want %d", numOK, bucket)
	}
}

func TestLongRunningQPS(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping in short mode")
	}
	if runtime.GOOS == "openbsd" {
		t.Skip("low resolution time.Sleep invalidates test (golang.org/issue/14183)")
		return
	}

	// The test runs for a few seconds executing many requests and then checks
	// that overall number of requests is reasonable.
	const (
		limit  = 100
		bucket = 100
	)
	var numOK = int32(0)

	lim := NewLimiter(limit, bucket)

	var wg sync.WaitGroup
	f := func() {
		if ok := lim.Take(time.Now(), 1, 0); ok {
			atomic.AddInt32(&numOK, 1)
		}
		wg.Done()
	}

	start := time.Now()
	end := start.Add(5 * time.Second)
	for time.Now().Before(end) {
		wg.Add(1)
		go f()

		// This will still offer ~500 requests per second, but won't consume
		// outrageous amount of CPU.
		time.Sleep(2 * time.Millisecond)
	}
	wg.Wait()

	elapsed := time.Since(start)
	ideal := bucket + (limit * float64(elapsed) / float64(time.Second))

	want1, want2 := int32(ideal+1), int32(0.999*ideal)

	fmt.Println(numOK, ideal, want1, want2)

	// We should never get more requests than allowed.
	if numOK > want1 {
		t.Errorf("numOK = %d, want %d (ideal %f)", numOK, want1, ideal)
	}
	// We should get very close to the number of requests allowed.
	if numOK < want2 {
		t.Errorf("numOK = %d, want %d (ideal %f)", numOK, want2, ideal)
	}
}
