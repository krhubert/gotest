package ostest

import (
	"os"
	"testing"
	"time"
)

// Truncate changes the size of the named file.
// If the file is a symbolic link, it changes the size of the link's target.
// If there is an error, it will be of type *PathError.
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

// Chtimes changes the access and modification times of the named
// file, similar to the Unix utime() or utimes() functions.
// A zero time.Time value will leave the corresponding file time unchanged.
//
// The underlying filesystem may truncate or round the values to a
// less precise time unit.
// If there is an error, it will be of type *PathError.
func Chtimes(t testing.TB, name string, atime time.Time, mtime time.Time) {
	t.Helper()
	err := os.Chtimes(name, atime, mtime)
	if err != nil {
		t.Fatal(err)
	}
	return
}

// Pipe returns a connected pair of Files; reads from r return bytes
// written to w. It returns the files and an error, if any.
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

// Chown changes the numeric uid and gid of the named file.
// If the file is a symbolic link, it changes the uid and gid of the link's target.
// A uid or gid of -1 means to not change that value.
// If there is an error, it will be of type *PathError.
//
// On Windows or Plan 9, Chown always returns the syscall.EWINDOWS or
// EPLAN9 error, wrapped in *PathError.
func Chown(t testing.TB, name string, uid, gid int) {
	t.Helper()
	err := os.Chown(name, uid, gid)
	if err != nil {
		t.Fatal(err)
	}
	return
}

// Lchown changes the numeric uid and gid of the named file.
// If the file is a symbolic link, it changes the uid and gid of the link itself.
// If there is an error, it will be of type *PathError.
func Lchown(t testing.TB, name string, uid, gid int) {
	t.Helper()
	err := os.Lchown(name, uid, gid)
	if err != nil {
		t.Fatal(err)
	}
	return
}
