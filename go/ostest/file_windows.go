package ostest

import (
	"os"
	"testing"
)

// Truncate changes the size of the named file.
// If the file is a symbolic link, it changes the size of the link's target.
func Truncate(t testing.TB, name string, size int64) {
	t.Helper()
	err := os.Truncate(name, size)
	if err != nil {
		t.Fatal(err)
	}
	return
}

// Remove removes the named file or directory.
// If there is an error, it will be of type *PathError.
func Remove(t testing.TB, name string) {
	t.Helper()
	err := os.Remove(name)
	if err != nil {
		t.Fatal(err)
	}
	return
}

// Pipe returns a connected pair of Files; reads from r return bytes written to w.
// It returns the files and an error, if any. The Windows handles underlying
// the returned files are marked as inheritable by child processes.
func Pipe(t testing.TB) (r *File, w *File) {
	t.Helper()
	p0, p1, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	return p0, p1
}

// Link creates newname as a hard link to the oldname file.
// If there is an error, it will be of type *LinkError.
func Link(t testing.TB, oldname, newname string) {
	t.Helper()
	err := os.Link(oldname, newname)
	if err != nil {
		t.Fatal(err)
	}
	return
}

// Symlink creates newname as a symbolic link to oldname.
// On Windows, a symlink to a non-existent oldname creates a file symlink;
// if oldname is later created as a directory the symlink will not work.
// If there is an error, it will be of type *LinkError.
func Symlink(t testing.TB, oldname, newname string) {
	t.Helper()
	err := os.Symlink(oldname, newname)
	if err != nil {
		t.Fatal(err)
	}
	return
}
