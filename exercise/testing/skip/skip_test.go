package skip

import (
	"os"
	"testing"
)

func Test_testingLimit(t *testing.T) {
	os.Setenv("SOME_ACCESS_TOKEN", "123")
	if os.Getenv("SOME_ACCESS_TOKEN") != "123" {
		t.Skip("skipping test, $SOME_ACCESS_TOKEN not set")
	}
	if testing.Short() {
		t.Skip("skipping malloc count in short mode")
	}
}
