package main

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// mid倒置补0 （补满共10位）+0 +yyyyMMdd          – 累计数据
// mid倒置补0 （补满共10位）+1 +yyyyMMdd          – 7日数据
// mid倒置补0 （补满共10位）+2 +yyyyMMdd        – 30日数据
// mid倒置补0 （补满共10位）+3 +yyyyMMdd        – 90日数据
func fansRowKey(mid int64, ty int) string {
	midStr := strconv.FormatInt(mid, 10)
	s := reverseString(midStr)
	l := len(s)
	if l < 10 {
		n := 10 - l
		for i := 0; i < n; i++ {
			s = s + "0"
		}
	}
	s = s + strconv.Itoa(ty) + time.Now().AddDate(0, 0, -1).Add(-12*time.Hour).Format("20060102")
	return s
}

// reverse for string.
func reverseString(s string) string {
	rs := []rune(s)
	l := len(rs) - 1
	for f, t := 0, l; f < t; f, t = f+1, t-1 {
		rs[f], rs[t] = rs[t], rs[f]
	}
	return string(rs)
}

func Test_reverse(t *testing.T) {
	id, ty := int64(45631212789), 0
	start := time.Now()
	println(fansRowKey(id, ty))
	elapsed := time.Since(start)
	fmt.Println("App elapsed: ", elapsed)
}

func main() {

}
