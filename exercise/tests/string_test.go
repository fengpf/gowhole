package tests

import (
	"fmt"
	"path"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"unicode"
	"unsafe"

	"github.com/qiniu/log"
)

// Substr get part of a string.
func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0
	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length
	if start > end {
		start, end = end, start
	}
	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}

func myTrim(s string) string {
	var (
		b    strings.Builder
		i, j int
	)
	for j < len(s) { //开启整个字符串的遍历
		for j < len(s) && unicode.IsSpace(rune(s[j])) { //空格位置计数
			println("unicode.IsSpace----", j, s[j])
			j++
		}
		i = j //先记录非空格位置
		println(i, j)
		for j < len(s) && !unicode.IsSpace(rune(s[j])) { //非空格位置计数
			j++
			println("!unicode.IsSpace>>>>", j, s[j])
		}
		if i != j { //如果空格标记不等于非空格标记则获取i-j段数据压入字节数组,并且重置当前位置
			println(i, j)
			b.Write([]byte(s[i:j]))
			i = j
			println(b.String())
		}
	}
	return b.String()
}

func Test_trim(t *testing.T) {
	str := "   111 ddssdds dsddsfff dsdsd dsds 222 "
	// fmt.Println(str)
	// str2 := strings.TrimSpace(str)
	// fmt.Println(str2)
	fmt.Println(myTrim(str))
}

type Bytes []byte

func StringBytes(s string) Bytes {
	return *(*Bytes)(unsafe.Pointer(&s))
}

func BytesString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// 获取&s[0]，即存储字符串的字节数组的地址指针，Go里不允许这种操作。
func StringPointer(s string) unsafe.Pointer {
	p := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return unsafe.Pointer(p.Data)
}

// r获取&b[0]，即[]byte底层数组的地址指针，Go里不允许这种操作
func BytesPointer(b []byte) unsafe.Pointer {
	p := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	return unsafe.Pointer(p.Data)
}
func Test_byte2String(t *testing.T) {
	s := "dd"
	fmt.Println(StringBytes(s))
	fmt.Println([]byte(s))
	fmt.Println(BytesString([]byte{255, 255, 255, 248}))
	fmt.Println(StringPointer(s))
	fmt.Println(BytesPointer([]byte{255, 255, 255, 248}))

	const nihongo = "汉语aas爱好"
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}
}

func Test_join(t *testing.T) {
	fmt.Println(joinPaths("/a", "test/"))
}

func lastChar(s string) uint8 {
	if s == "" {
		panic("The length of the string can't be 0")
	}
	return s[len(s)-1]
}

func joinPaths(absolutePath, relativePath string) string {
	if relativePath == "" {
		return absolutePath
	}
	finalPath := path.Join(absolutePath, relativePath)
	appendSlash := lastChar(relativePath) == '/' && lastChar(finalPath) != '/'
	if appendSlash {
		return finalPath + "/"
	}
	return finalPath
}

func Test_split(t *testing.T) {
	var tids []int64
	tidsStr := "1,2,3"
	tidSlice := strings.Split(tidsStr, ",")
	tids = make([]int64, 0, len(tidSlice))
	for _, v := range tidSlice {
		tid, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			log.Error("strconv.ParseInt(%s) error(%v)", v, err)
			return
		}
		tids = append(tids, tid)
	}
	fmt.Println(tids)
}
