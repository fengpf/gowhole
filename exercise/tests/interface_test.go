package basictests

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
	"unicode"
)

type Stringer interface {
	String() string
}

type Binary uint64

func (i Binary) String() string {
	return strconv.FormatUint(i.Get(), 2)
}

func (i Binary) Get() uint64 {
	return uint64(i)
}

func Test_stringer(t *testing.T) {
	var b Binary = 32
	//首次遇见s := Stringer(b)这样的语句时，golang会生成Stringer接口对应于Binary类型的虚表，并将其缓存
	s := Stringer(b)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", b.Get())
	emptyInterface()

	c := 100
	d := &c
	*d++
	fmt.Printf("c pointer(%p), value(%v) \n\r", &c, c)
	fmt.Printf("d pointer(%p), value(%v) pointerValue(%d)\n\r", &d, d, *d)
}

func emptyInterface() {
	//接口类型的一个极端重要的例子是空接口：interface{},它表示空的方法集合，
	//由于任何值都有零个或者多个方法，所以任何值都可以满足它。 注意，[]T不能直接赋值给[]interface{}
	t := []int{1, 2, 3, 4}
	// var s []interface{} = t  //wrong
	s := make([]interface{}, len(t))
	for i, v := range t {
		s[i] = v
	}
	var value interface{}
	fmt.Printf("value is: %+v\n", value)
	switch str := value.(type) {
	case string:
		fmt.Printf("string value is: %q\n", str)
	case Stringer:
		fmt.Printf("value is not a string\n")
	}
}

type MyCounter int

func (m *MyCounter) Write(b []byte) (int, error) {
	s := string(b)
	var i, j, n int
	fmt.Println(s)
	for i < len(s) {
		for i < len(s) && unicode.IsSpace(rune(s[i])) {
			i++
		}
		j = i
		for i < len(s) && !unicode.IsSpace(rune(s[i])) {
			i++
		}
		if i != j {
			n = n + 1
			*m += MyCounter(n)
			j = i
		}
	}
	return n, nil
}

func Test_myCounter(t *testing.T) {
	var myc MyCounter
	a := " dds dsds "
	myc.Write([]byte(a))
	fmt.Println(myc)
	var any interface{}
	any = true
	fmt.Println(any)
	any = 12.34
	fmt.Println(any)
	any = "dsds"
	fmt.Println(any)
	any = map[string]int{"one": 1}
	fmt.Println(any)
	any = new(bytes.Buffer)
}

func Test_implementInterface(t *testing.T) {
	var w io.Writer
	w = os.Stdout
	w.Write([]byte("Stdout\n"))
	w = new(bytes.Buffer)
	w.Write([]byte("Buffer\n"))
}

func Test_scanWords(t *testing.T) {
	// An artificial input source.
	const input = "Now is the winter of our discontent,\nMade glorious summer by this sun of York.\n"
	scanner := bufio.NewScanner(strings.NewReader(input))
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	// Count the words.
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Printf("%d\n", count)
}

func Test_flag(t *testing.T) {
	var period = flag.Duration("period", 1*time.Second, "sleep period")
	flag.Parse()
	fmt.Printf("Sleeping for %v...\n", *period)
	time.Sleep(*period)

	var temp = CelsiusFlag("temp", 20.0, "the temperature")
	fmt.Printf("temperature is %v...%b\n", *temp, 10)
	fmt.Println()
}

var p = fmt.Println

func Test_flag2(t *testing.T) {
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
