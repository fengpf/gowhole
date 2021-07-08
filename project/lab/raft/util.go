package raft

import (
	"gowhole/middleware/log"
)

// Debugging
const Debug = 1

func DPrintf(format string, a ...interface{}) {
	log.Info(format, a...)
}
