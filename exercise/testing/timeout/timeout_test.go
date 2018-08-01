package timeout

import (
	"testing"
	"time"
)

func TestWillTimeout(t *testing.T) {
	time.Sleep(1 * time.Second)
	// pass if timeout > 2s
}
