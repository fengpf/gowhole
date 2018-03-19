package main

import (
	"strconv"
	"sync/atomic"
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

func main() {
	s := new(severity)
	s.set(warningLog)
	println(s.get())
	println(s.string())
}
