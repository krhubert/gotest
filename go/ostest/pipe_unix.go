package ostest

import (
	"os"
	"testing"
)

// Pipe returns a connected pair of Files; reads from r return bytes written to w.
// It returns the files and an error, if any.
func Pipe(t testing.TB) (r *os.File, w *os.File) {
	t.Helper()
	p0, p1, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	return p0, p1
}
