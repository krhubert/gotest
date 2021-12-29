package ostest

import (
	"os"
	"testing"
)

// MkdirAll creates a directory named path,
// along with any necessary parents, and returns nil,
// or else returns an error.
// The permission bits perm (before umask) are used for all
// directories that MkdirAll creates.
// If path is already a directory, MkdirAll does nothing
// and returns nil.
func MkdirAll(t testing.TB, path string, perm os.FileMode) {
	t.Helper()
	err := os.MkdirAll(path, perm)
	if err != nil {
		t.Fatal(err)
	}
	return
}

// RemoveAll removes path and any children it contains.
// It removes everything it can but returns the first error
// it encounters. If the path does not exist, RemoveAll
// returns nil (no error).
// If there is an error, it will be of type [*PathError].
func RemoveAll(t testing.TB, path string) {
	t.Helper()
	err := os.RemoveAll(path)
	if err != nil {
		t.Fatal(err)
	}
	return
}
