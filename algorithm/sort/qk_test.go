package sort

import (
	"math/rand"
	"sort"
	"sync"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
)

var (
	nums = GenerateRandomNumber(0, 50000, 20000)
)

func Benchmark_quick(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.Slice(nums, func(i, j int) bool {
			return nums[i] > nums[j]
		})
		spew.Dump(nums)
	}
}

func QuickSort(src []int, wg *sync.WaitGroup) {
	defer wg.Done()
	go sort.Slice(src, func(i, j int) bool {
		return src[i] > src[j]
	})
}

func Test_quick(t *testing.T) {
	nums := []int{2, 1, 3, 4}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go QuickSort3(nums, 0, len(nums)-1, wg)
	// QuickSort(nums, wg)
	wg.Wait()
	spew.Dump(nums)
}

func GenerateRandomNumber(start int, end int, count int) []int {
	if end < start || (end-start) < count {
		return nil
	}
	nums := make([]int, 0)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		num := r.Intn((end - start)) + start
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}
		if !exist {
			nums = append(nums, num)
		}
	}
	return nums
}

func TestRandomNum(t *testing.T) {
	s := []int{11, 12, 13, 14, 15}
	keys := GenerateRandomNumber(0, len(s), len(s))
	spew.Dump(keys)
	spew.Dump(s)

	ns := make([]int, 0, len(s))
	for _, k := range keys {
		ns = append(ns, s[k])
	}
	spew.Dump(ns)
}
