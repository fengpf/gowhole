package test

import (
	"testing"
)


func Test_RPCAll(t *testing.T) {
	once.Do(StartServer)
	CallRPC(serverAddr)
}
