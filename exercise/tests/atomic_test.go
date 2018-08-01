package basictests

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type severity int32

const (
	infoLog severity = iota
	warningLog
	errorLog
	fatalLog
	numSeverity = 4
)

const severityChar = "IWEF"

var severityName = []string{
	infoLog:    "INFO",
	warningLog: "WARNING",
	errorLog:   "ERROR",
	fatalLog:   "FATAL",
}

func (s *severity) get() severity {
	return severity(atomic.LoadInt32((*int32)(s)))
}

func (s *severity) set(val severity) {
	atomic.StoreInt32((*int32)(s), int32(val))
}

func (s *severity) string() string {
	return strconv.FormatInt(int64(*s), 10)
}

func Test_set_get(t *testing.T) {
	s := new(severity)
	s.set(warningLog)
	println(s.get())
	println(s.string())
}

type atomicInt struct {
	value int
	lock  sync.Mutex
}

func (a *atomicInt) increment() {
	fmt.Println("safe increment")
	func() {
		a.lock.Lock()
		defer a.lock.Unlock()
		a.value++
	}()
}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}

func Test_incr(t *testing.T) {
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}

func Test_com(t *testing.T) {
	addVal(int32(1))
}

func addVal(delta int32) {
	var val int32
	for {
		old := val
		new := val + delta
		fmt.Println(val, old, new)
		if atomic.CompareAndSwapInt32(&val, old, new) {
			fmt.Println(val, old, new)
			break
		}
	}
}
