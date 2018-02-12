package main

import (
	"fmt"
	"sort"
)

type total []int64

func (t total) Len() int           { return len(t) }
func (t total) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t total) Less(i, j int) bool { return t[i] > t[j] }

func main() {
	t := []int64{789, 123, 456}
	// sort.Sort(total(t))
	sort.Slice(t, func(i, j int) bool {
		return t[i] < t[j]
	})
	fmt.Printf("%v\n", t)
}
