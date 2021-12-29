package ostest

import (
	"os"
	"testing"
)

// Hostname returns the host name reported by the kernel.
func Hostname(t testing.TB) (name string) {
	t.Helper()
	p0, err := os.Hostname()
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
