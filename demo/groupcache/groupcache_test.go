package consistenthash

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/golang/groupcache/consistenthash"
	"github.com/golang/groupcache/lru"
	"github.com/golang/groupcache/singleflight"
)

func Test_consistenthash(t *testing.T) {
	c := consistenthash.New(5, nil)
	c.Add("A", "B", "C", "D", "E")
	for _, k := range []string{"what", "nice", "what", "nice", "good", "yes!"} {
		fmt.Printf("%s -> %s\n", k, c.Get(k))
	}
}

func NewDelayReturn(dur time.Duration, n int) func() (interface{}, error) {
	return func() (interface{}, error) {
		time.Sleep(dur)
		return n, nil
	}
}
func Test_singleflight(t *testing.T) {
	g := singleflight.Group{}
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		ret, err := g.Do("key", NewDelayReturn(time.Second*1, 1))
		if err != nil {
			panic(err)
		}
		fmt.Printf("key-1 get %v\n", ret)
		wg.Done()
	}()
	go func() {
		time.Sleep(100 * time.Millisecond) // make sure this is call is later
		ret, err := g.Do("key", NewDelayReturn(time.Second*2, 2))
		if err != nil {
			panic(err)
		}
		fmt.Printf("key-2 get %v\n", ret)
		wg.Done()
	}()
	wg.Wait()
}

func Test_lru(t *testing.T) {
	cache := lru.New(2)
	cache.Add("x", "x0")
	cache.Add("y", "y0")
	yval, ok := cache.Get("y")
	if ok {
		fmt.Printf("y is %v\n", yval)
	}
	cache.Add("z", "z0")

	fmt.Printf("cache length is %d\n", cache.Len())
	_, ok = cache.Get("x")
	if !ok {
		fmt.Printf("x key was weeded out\n")
	}
}
