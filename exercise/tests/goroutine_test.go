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

func Test_rely(t *testing.T) {
	a := func() []int {
		b := []int{1, 2, 3}
		return b
	}()

	go func() {
		for _, v := range a {
			i := v + 1
			fmt.Println(i)
		}
	}()
}

func Test_loop(t *testing.T) {
	values := []int{1, 2, 3}
	for _, val := range values {
		go func() {
			fmt.Println(val)
		}()
	}
}

func Test_loop1(t *testing.T) {
	values := []int{1, 2, 3}
	for _, val := range values {
		go func(val interface{}) {
			fmt.Println(val)
		}(val)
	}
}

func Test_loop2(t *testing.T) {
	values := []int{1, 2, 3}
	for i := range values {
		val := values[i]
		go func() {
			fmt.Println(val)
		}()
	}
}

func Test_loop3(t *testing.T) {
	values := []int{1, 2, 3}
	for i := range values {
		val := values[i]
		go func() {
			fmt.Println(val)
		}()
	}
}

func Test_loop4(t *testing.T) {
	for i := 1; i <= 10; i++ {
		func() {
			fmt.Println(i)
		}()
	}
}

// type val interface{}

// func (v *val) MyMethod() {
// 	fmt.Println(v)
// }

// func Test_loop5(t *testing.T) {
// 	values := []int{1, 2, 3}
// 	for _, val := range values {
// 		go &val.MyMethod()
// 	}
// }

// type newVal interface{}

// func (v *newVal) MyMethod() {
// 	fmt.Println(v)
// }
// func Test_loop6(t *testing.T) {
// 	values := []int{1, 2, 3}
// 	for _, val := range values {
// 		newVal := val
// 		go newVal.MyMethod()
// 	}
// }
