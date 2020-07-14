package raft

import "fmt"

// Debugging
const Debug = 1

func DPrintf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}
