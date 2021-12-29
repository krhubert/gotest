package ostest

import (
	"os"
	"testing"
)

// Getwd returns a rooted path name corresponding to the
// current directory. If the current directory can be
// reached via multiple paths (due to symbolic links),
// Getwd may return any one of them.
func Getwd(t testing.TB) (dir string) {
	t.Helper()
	p0, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
