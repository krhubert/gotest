package fstest

import (
	"io/fs"
	"testing"
)

// Stat returns a [FileInfo] describing the named file from the file system.
//
// If fs implements [StatFS], Stat calls fs.Stat.
// Otherwise, Stat opens the [File] to stat it.
func Stat(t testing.TB, fsys fs.FS, name string) fs.FileInfo {
	t.Helper()
	p0, err := fs.Stat(fsys, name)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
