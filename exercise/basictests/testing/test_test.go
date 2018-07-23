package basictests

import (
	"runtime"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

// callerName gives the function name (qualified with a package path)
// for the caller after skip frames (where 0 means the current function).
func callerName(skip int) string {
	// Make room for the skip PC.
	var pc [2]uintptr
	n := runtime.Callers(skip+2, pc[:]) // skip + runtime.Callers + callerName
	if n == 0 {
		panic("testing: zero callers found")
	}
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	spew.Dump(frame)
	println(frame.Function)
	return frame.Function
}

// 今天有人问了个很有意思的问题。TestA 内调用 t.Parallel 会让测试变成并发执行。
// 那么问题来了，Parallel 执行的时候，TestA 已经在执行了，又怎么变成并发的呢？
func Test_caller(t *testing.T) {
	t.Parallel() //会发送信号给他的父级test 等待它
	println("A")
	callerName(0)
	// dir, _ := os.Getwd()
	// fmt.Println(dir)
}

// go test -v ./...
// go test -v ./... -run TestTLog/test_4
func TestTLog(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		value int
	}{
		{name: "test 1", value: 1},
		{name: "test 2", value: 2},
		{name: "test 3", value: 3},
		{name: "test 4", value: 4},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			// Here you test tc.value against a test function.
			// Let's use t.Log as our test function :-)
			t.Parallel() //When calling t.Parallel() the test sends a signal to its parent test to stop waiting for it, and then the loops continues.
			t.Log(tc.value)
		})
	}
}
