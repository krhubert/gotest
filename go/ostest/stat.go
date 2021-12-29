package ostest

import (
	"os"
	"testing"
)

// Stat returns a [FileInfo] describing the named file.
// If there is an error, it will be of type [*PathError].
func Stat(t testing.TB, name string) os.FileInfo {
	t.Helper()
	p0, err := os.Stat(name)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// Lstat returns a [FileInfo] describing the named file.
// If the file is a symbolic link, the returned FileInfo
// describes the symbolic link. Lstat makes no attempt to follow the link.
// If there is an error, it will be of type [*PathError].
//
// On Windows, if the file is a reparse point that is a surrogate for another
// named entity (such as a symbolic link or mounted folder), the returned
// FileInfo describes the reparse point, and makes no attempt to resolve it.
func Lstat(t testing.TB, name string) os.FileInfo {
	t.Helper()
	p0, err := os.Lstat(name)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
