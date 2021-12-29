package ostest

import (
	"io/fs"
	"os"
	"testing"
)

// ReadDir reads the named directory,
// returning all its directory entries sorted by filename.
// If an error occurs reading the directory,
// ReadDir returns the entries it was able to read before the error,
// along with the error.
func ReadDir(t testing.TB, name string) []os.DirEntry {
	t.Helper()
	p0, err := os.ReadDir(name)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// CopyFS copies the file system fsys into the directory dir,
// creating dir if necessary.
//
// Files are created with mode 0o666 plus any execute permissions
// from the source, and directories are created with mode 0o777
// (before umask).
//
// CopyFS will not overwrite existing files, and returns an error
// if a file name in fsys already exists in the destination.
//
// Symbolic links in fsys are not supported. A *PathError with Err set
// to ErrInvalid is returned when copying from a symbolic link.
//
// Symbolic links in dir are followed.
//
// Copying stops at and returns the first error encountered.
func CopyFS(t testing.TB, dir string, fsys fs.FS) {
	t.Helper()
	err := os.CopyFS(dir, fsys)
	if err != nil {
		t.Fatal(err)
	}
	return
}
