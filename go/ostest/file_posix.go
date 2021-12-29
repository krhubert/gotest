package ostest

import (
	"os"
	"testing"
	"time"
)

// Chown changes the numeric uid and gid of the named file.
// If the file is a symbolic link, it changes the uid and gid of the link's target.
// A uid or gid of -1 means to not change that value.
// If there is an error, it will be of type [*PathError].
//
// On Windows or Plan 9, Chown always returns the [syscall.EWINDOWS] or
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
// If there is an error, it will be of type [*PathError].
//
// On Windows, it always returns the [syscall.EWINDOWS] error, wrapped
// in *PathError.
func Lchown(t testing.TB, name string, uid, gid int) {
	t.Helper()
	err := os.Lchown(name, uid, gid)
	if err != nil {
		t.Fatal(err)
	}
	return
}

// Chtimes changes the access and modification times of the named
// file, similar to the Unix utime() or utimes() functions.
// A zero [time.Time] value will leave the corresponding file time unchanged.
//
// The underlying filesystem may truncate or round the values to a
// less precise time unit.
// If there is an error, it will be of type [*PathError].
func Chtimes(t testing.TB, name string, atime time.Time, mtime time.Time) {
	t.Helper()
	err := os.Chtimes(name, atime, mtime)
	if err != nil {
		t.Fatal(err)
	}
	return
}
