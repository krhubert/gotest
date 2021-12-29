package fstest

import (
	"io/fs"
	"testing"
)

// WalkDir walks the file tree rooted at root, calling fn for each file or
// directory in the tree, including root.
//
// All errors that arise visiting files and directories are filtered by fn:
// see the [fs.WalkDirFunc] documentation for details.
//
// The files are walked in lexical order, which makes the output deterministic
// but requires WalkDir to read an entire directory into memory before proceeding
// to walk that directory.
//
// WalkDir does not follow symbolic links found in directories,
// but if root itself is a symbolic link, its target will be walked.
func WalkDir(t testing.TB, fsys fs.FS, root string, fn fs.WalkDirFunc) {
	t.Helper()
	err := fs.WalkDir(fsys, root, fn)
	if err != nil {
		t.Fatal(err)
	}
	return
}
