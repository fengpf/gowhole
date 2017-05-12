package rpcx

import (
	"fmt"
)

// 参考- http://blog.csdn.net/gqtcgq/article/details/52076997

// Server weight-round-robin.
type Server struct {
	Name      string
	Weight    int
	CurWeight int
}

func getsum(set []int, size int) (sum int) {
	for i := 0; i < size; i++ {
		sum += set[i]
	}
	return
}

// greatest common divisor
func gcd(a, b int) int {
	var c int
	for b > 0 {
		c = a % b
		a = b
		b = c
	}
	return a
}

func getgcd(set []int, size int) (res int) {
	res = set[0]
	for i := 0; i < size; i++ {
		res = gcd(res, set[i])
	}
	return
}

func getmax(set []int, size int) (max int) {
	max = set[0]
	for i := 0; i < size; i++ {
		if max < set[i] {
			max = set[i]
		}
	}
	return
}

// get weight-round-robin
func getwrr(s []*Server, size, gcd, max int, i, cw *int) int {
	for {
		*i = (*i + 1) % size
		if *i == 0 {
			*cw = *cw - gcd
			if *cw <= 0 {
				*cw = max
				if *cw == 0 {
					return -1
				}
			}
		}
		if s[*i].Weight >= *cw {
			return *i
		}
	}
}

func wrr(s []*Server, weights []int, size int) {
	var (
		index, curweight int
	)
	index = -1
	curweight = 0
	gcd := getgcd(weights, size)
	max := getmax(weights, size)
	sum := getsum(weights, size)
	for i := 0; i < sum; i++ {
		getwrr(s, size, gcd, max, &index, &curweight)
		fmt.Printf("%s(%d)\n", s[index].Name, s[index].Weight)
	}
	println("----------")
}

func initServers(names []string, ws []int, size int) (ss []*Server) {
	ss = make([]*Server, 0, size)
	for i := 0; i < size; i++ {
		s := &Server{}
		s.Name = names[i]
		s.Weight = ws[i]
		ss = append(ss, s)
	}
	return
}

// PrintWrr test wrr.
func PrintWrr() {
	weights := []int{1, 2, 4}
	names := []string{"a", "b", "c"}
	size := len(weights)
	s := initServers(names, weights, size)
	for i := 0; i < size; i++ {
		fmt.Printf("%s(%d)\n", s[i].Name, s[i].Weight)
	}
	println("----------")
	println("wrr sequence is")
	wrr(s, weights, size)
}
