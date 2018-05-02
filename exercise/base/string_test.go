package base

import (
	"fmt"
	"strings"
	"testing"
	"unicode"
)

var p = fmt.Println

func Test_flag(t *testing.T) {
	p("Contains:  ", strings.Contains("test", "es"))
	p("Count:     ", strings.Count("test", "t"))
	p("HasPrefix: ", strings.HasPrefix("test", "te"))
	p("HasSuffix: ", strings.HasSuffix("test", "st"))
	p("Index:     ", strings.Index("test", "e"))
	p("Join:      ", strings.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", strings.Repeat("a", 5))
	p("Replace:   ", strings.Replace("foo", "o", "0", -1))
	p("Replace:   ", strings.Replace("foo", "o", "0", 1))
	p("Split:     ", strings.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", strings.ToLower("TEST"))
	p("ToUpper:   ", strings.ToUpper("test"))
	p()
	p("Len: ", len("hello"))
	p("Char:", "hello"[1])
}

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
	for j < len(s) {
		for j < len(s) && !unicode.IsSpace(rune(s[j])) {
			j++
			println("!unicode.IsSpace>>>>", j, s[j])
		}
		if i != j {
			b.Write([]byte(s[i:j]))
			i = j
			println(b.String())
		}
		for j < len(s) && unicode.IsSpace(rune(s[j])) {
			println("unicode.IsSpace----", j, s[j])
			j++
		}
		println(i, j)
		i = j
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
