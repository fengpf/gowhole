package test

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

// KE 代表键-元素对。
type KE struct {
	key     string
	element int
}

// BenchmarkMapPut 用于对字典的添加和修改操作进行测试。
func BenchmarkMapPut(b *testing.B) {
	max := 5
	var kes []KE
	for i := 0; i <= max; i++ {
		kes = append(kes, KE{strconv.Itoa(i), rand.Intn(1000000)})
	}
	m := make(map[string]int)
	b.ResetTimer()
	for _, ke := range kes {
		k, e := ke.key, ke.element
		b.Run(fmt.Sprintf("Key: %s, Element: %#v", k, e), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				m[k] = e + i
			}
		})
	}
}
