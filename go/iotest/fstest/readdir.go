package fstest

import (
	"io/fs"
	"testing"
)

// ReadDir reads the named directory
// and returns a list of directory entries sorted by filename.
//
// If fs implements [ReadDirFS], ReadDir calls fs.ReadDir.
// Otherwise ReadDir calls fs.Open and uses ReadDir and Close
// on the returned file.
func ReadDir(t testing.TB, fsys fs.FS, name string) []fs.DirEntry {
	t.Helper()
	p0, err := fs.ReadDir(fsys, name)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
