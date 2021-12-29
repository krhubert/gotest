package fstest

import (
	"io/fs"
	"testing"
)

// ReadFile reads the named file from the file system fs and returns its contents.
// A successful call returns a nil error, not [io.EOF].
// (Because ReadFile reads the whole file, the expected EOF
// from the final Read is not treated as an error to be reported.)
//
// If fs implements [ReadFileFS], ReadFile calls fs.ReadFile.
// Otherwise ReadFile calls fs.Open and uses Read and Close
// on the returned [File].
func ReadFile(t testing.TB, fsys fs.FS, name string) []byte {
	t.Helper()
	p0, err := fs.ReadFile(fsys, name)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
