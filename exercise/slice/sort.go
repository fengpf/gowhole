package slice

import (
	"fmt"
	"sort"
	"testing"
)

type total []int64

func (t total) Len() int           { return len(t) }
func (t total) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t total) Less(i, j int) bool { return t[i] > t[j] }

func Test_sort(t *testing.T) {
	m := []int64{789, 123, 456}
	// sort.Sort(total(t))
	sort.Slice(m, func(i, j int) bool {
		return m[i] < m[j]
	})
	fmt.Printf("%v\n", m)
}
