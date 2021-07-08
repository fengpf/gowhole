package log_zap

import "testing"

func Test_log(t *testing.T) {
	Info("test id:%d", 1)
}