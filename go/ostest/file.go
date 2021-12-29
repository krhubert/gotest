package ostest

import (
	"os"
	"testing"
)

// Mkdir creates a new directory with the specified name and permission
// bits (before umask).
// If there is an error, it will be of type *PathError.
func Mkdir(t testing.TB, name string, perm os.FileMode) {
	t.Helper()
	err := os.Mkdir(name, perm)
	if err != nil {
		t.Fatal(err)
	}
	return
}

// Chdir changes the current working directory to the named directory.
// If there is an error, it will be of type *PathError.
func Chdir(t testing.TB, dir string) {
	t.Helper()
	err := os.Chdir(dir)
	if err != nil {
		t.Fatal(err)
	}
	return
}

// Open opens the named file for reading. If successful, methods on
// the returned file can be used for reading; the associated file
// descriptor has mode O_RDONLY.
// If there is an error, it will be of type *PathError.
func Open(t testing.TB, name string) *os.File {
	t.Helper()
	p0, err := os.Open(name)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// Create creates or truncates the named file. If the file already exists,
// it is truncated. If the file does not exist, it is created with mode 0o666
// (before umask). If successful, methods on the returned File can
// be used for I/O; the associated file descriptor has mode O_RDWR.
// If there is an error, it will be of type *PathError.
func Create(t testing.TB, name string) *os.File {
	t.Helper()
	p0, err := os.Create(name)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// OpenFile is the generalized open call; most users will use Open
// or Create instead. It opens the named file with specified flag
// (O_RDONLY etc.). If the file does not exist, and the O_CREATE flag
// is passed, it is created with mode perm (before umask). If successful,
// methods on the returned File can be used for I/O.
// If there is an error, it will be of type *PathError.
func OpenFile(t testing.TB, name string, flag int, perm os.FileMode) *os.File {
	t.Helper()
	p0, err := os.OpenFile(name, flag, perm)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// Rename renames (moves) oldpath to newpath.
// If newpath already exists and is not a directory, Rename replaces it.
// OS-specific restrictions may apply when oldpath and newpath are in different directories.
// Even within the same directory, on non-Unix platforms Rename is not an atomic operation.
// If there is an error, it will be of type *LinkError.
func Rename(t testing.TB, oldpath, newpath string) {
	t.Helper()
	err := os.Rename(oldpath, newpath)
	if err != nil {
		t.Fatal(err)
	}
	return
}

// Readlink returns the destination of the named symbolic link.
// If there is an error, it will be of type *PathError.
//
// If the link destination is relative, Readlink returns the relative path
// without resolving it to an absolute one.
func Readlink(t testing.TB, name string) string {
	t.Helper()
	p0, err := os.Readlink(name)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// UserCacheDir returns the default root directory to use for user-specific
// cached data. Users should create their own application-specific subdirectory
// within this one and use that.
//
// On Unix systems, it returns $XDG_CACHE_HOME as specified by
// https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html if
// non-empty, else $HOME/.cache.
// On Darwin, it returns $HOME/Library/Caches.
// On Windows, it returns %LocalAppData%.
// On Plan 9, it returns $home/lib/cache.
//
// If the location cannot be determined (for example, $HOME is not defined),
// then it will return an error.
func UserCacheDir(t testing.TB) string {
	t.Helper()
	p0, err := os.UserCacheDir()
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// UserConfigDir returns the default root directory to use for user-specific
// configuration data. Users should create their own application-specific
// subdirectory within this one and use that.
//
// On Unix systems, it returns $XDG_CONFIG_HOME as specified by
// https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html if
// non-empty, else $HOME/.config.
// On Darwin, it returns $HOME/Library/Application Support.
// On Windows, it returns %AppData%.
// On Plan 9, it returns $home/lib.
//
// If the location cannot be determined (for example, $HOME is not defined),
// then it will return an error.
func UserConfigDir(t testing.TB) string {
	t.Helper()
	p0, err := os.UserConfigDir()
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// UserHomeDir returns the current user's home directory.
//
// On Unix, including macOS, it returns the $HOME environment variable.
// On Windows, it returns %USERPROFILE%.
// On Plan 9, it returns the $home environment variable.
//
// If the expected variable is not set in the environment, UserHomeDir
// returns either a platform-specific default value or a non-nil error.
func UserHomeDir(t testing.TB) string {
	t.Helper()
	p0, err := os.UserHomeDir()
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// Chmod changes the mode of the named file to mode.
// If the file is a symbolic link, it changes the mode of the link's target.
// If there is an error, it will be of type *PathError.
//
// A different subset of the mode bits are used, depending on the
// operating system.
//
// On Unix, the mode's permission bits, ModeSetuid, ModeSetgid, and
// ModeSticky are used.
//
// On Windows, only the 0o200 bit (owner writable) of mode is used; it
// controls whether the file's read-only attribute is set or cleared.
// The other bits are currently unused. For compatibility with Go 1.12
// and earlier, use a non-zero mode. Use mode 0o400 for a read-only
// file and 0o600 for a readable+writable file.
//
// On Plan 9, the mode's permission bits, ModeAppend, ModeExclusive,
// and ModeTemporary are used.
func Chmod(t testing.TB, name string, mode os.FileMode) {
	t.Helper()
	err := os.Chmod(name, mode)
	if err != nil {
		t.Fatal(err)
	}
	return
}

// ReadFile reads the named file and returns the contents.
// A successful call returns err == nil, not err == EOF.
// Because ReadFile reads the whole file, it does not treat an EOF from Read
// as an error to be reported.
func ReadFile(t testing.TB, name string) []byte {
	t.Helper()
	p0, err := os.ReadFile(name)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// WriteFile writes data to the named file, creating it if necessary.
// If the file does not exist, WriteFile creates it with permissions perm (before umask);
// otherwise WriteFile truncates it before writing, without changing permissions.
// Since WriteFile requires multiple system calls to complete, a failure mid-operation
// can leave the file in a partially written state.
func WriteFile(t testing.TB, name string, data []byte, perm os.FileMode) {
	t.Helper()
	err := os.WriteFile(name, data, perm)
	if err != nil {
		t.Fatal(err)
	}
	return
}
