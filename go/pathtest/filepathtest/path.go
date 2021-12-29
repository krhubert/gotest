package filepathtest

import (
	"io/fs"
	"path/filepath"
	"testing"
)

// Localize converts a slash-separated path into an operating system path.
// The input path must be a valid path as reported by [io/fs.ValidPath].
//
// Localize returns an error if the path cannot be represented by the operating system.
// For example, the path a\b is rejected on Windows, on which \ is a separator
// character and cannot be part of a filename.
//
// The path returned by Localize will always be local, as reported by IsLocal.
func Localize(t testing.TB, path string) string {
	t.Helper()
	p0, err := filepath.Localize(path)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// EvalSymlinks returns the path name after the evaluation of any symbolic
// links.
// If path is relative the result will be relative to the current directory,
// unless one of the components is an absolute symbolic link.
// EvalSymlinks calls [Clean] on the result.
func EvalSymlinks(t testing.TB, path string) string {
	t.Helper()
	p0, err := filepath.EvalSymlinks(path)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// Abs returns an absolute representation of path.
// If the path is not absolute it will be joined with the current
// working directory to turn it into an absolute path. The absolute
// path name for a given file is not guaranteed to be unique.
// Abs calls [Clean] on the result.
func Abs(t testing.TB, path string) string {
	t.Helper()
	p0, err := filepath.Abs(path)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// Rel returns a relative path that is lexically equivalent to targpath when
// joined to basepath with an intervening separator. That is,
// [Join](basepath, Rel(basepath, targpath)) is equivalent to targpath itself.
// On success, the returned path will always be relative to basepath,
// even if basepath and targpath share no elements.
// An error is returned if targpath can't be made relative to basepath or if
// knowing the current working directory would be necessary to compute it.
// Rel calls [Clean] on the result.
func Rel(t testing.TB, basepath, targpath string) string {
	t.Helper()
	p0, err := filepath.Rel(basepath, targpath)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

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
// WalkDir does not follow symbolic links.
//
// WalkDir calls fn with paths that use the separator character appropriate
// for the operating system. This is unlike [io/fs.WalkDir], which always
// uses slash separated paths.
func WalkDir(t testing.TB, root string, fn fs.WalkDirFunc) {
	t.Helper()
	err := filepath.WalkDir(root, fn)
	if err != nil {
		t.Fatal(err)
	}
	return
}

// Walk walks the file tree rooted at root, calling fn for each file or
// directory in the tree, including root.
//
// All errors that arise visiting files and directories are filtered by fn:
// see the [WalkFunc] documentation for details.
//
// The files are walked in lexical order, which makes the output deterministic
// but requires Walk to read an entire directory into memory before proceeding
// to walk that directory.
//
// Walk does not follow symbolic links.
//
// Walk is less efficient than [WalkDir], introduced in Go 1.16,
// which avoids calling os.Lstat on every visited file or directory.
func Walk(t testing.TB, root string, fn filepath.WalkFunc) {
	t.Helper()
	err := filepath.Walk(root, fn)
	if err != nil {
		t.Fatal(err)
	}
	return
}
