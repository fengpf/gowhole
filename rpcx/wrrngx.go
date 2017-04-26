package rpcx

import (
	"fmt"
)

func getNextServerIndex(s []*Server, size int) int {
	var index, sum int
	index = -1
	for i := 0; i < size; i++ {
		s[i].CurWeight += s[i].Weight
		sum += s[i].Weight
		if index == -1 || s[index].CurWeight < s[i].CurWeight {
			index = i
		}
	}
	s[index].CurWeight -= sum
	return index
}

func wrrNgx(s []*Server, weights []int, size int) {
	var (
		index int
	)
	index = -1
	sum := getsum(weights, size)
	for i := 0; i < sum; i++ {
		index = getNextServerIndex(s, size)
		fmt.Printf("%s(%d)\n", s[index].Name, s[index].Weight)
	}
	println("----------")
}

// PrintWrrNgx test ngx wrr.
func PrintWrrNgx() {
	weights := []int{4, 2, 1, 3}
	names := []string{"a", "b", "c", "d"}
	size := len(weights)
	s := initServers(names, weights, size)
	for i := 0; i < size; i++ {
		fmt.Printf("%s(%d)\n", s[i].Name, s[i].Weight)
	}
	println("----------")
	println("nwrr_nginx sequence is")
	wrrNgx(s, weights, size)
}
