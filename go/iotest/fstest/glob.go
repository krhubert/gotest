package fstest

import (
	"io/fs"
	"testing"
)

// Glob returns the names of all files matching pattern or nil
// if there is no matching file. The syntax of patterns is the same
// as in [path.Match]. The pattern may describe hierarchical names such as
// usr/*/bin/ed.
//
// Glob ignores file system errors such as I/O errors reading directories.
// The only possible returned error is [path.ErrBadPattern], reporting that
// the pattern is malformed.
//
// If fs implements [GlobFS], Glob calls fs.Glob.
// Otherwise, Glob uses [ReadDir] to traverse the directory tree
// and look for matches for the pattern.
func Glob(t testing.TB, fsys fs.FS, pattern string) (matches []string) {
	t.Helper()
	p0, err := fs.Glob(fsys, pattern)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
