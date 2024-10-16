package ioutiltest

import (
	"io"
	"io/fs"
	"io/ioutil"
	"testing"
)

// ReadAll reads from r until an error or EOF and returns the data it read.
// A successful call returns err == nil, not err == EOF. Because ReadAll is
// defined to read from src until EOF, it does not treat an EOF from Read
// as an error to be reported.
//
// Deprecated: As of Go 1.16, this function simply calls [io.ReadAll].
func ReadAll(t testing.TB, r io.Reader) []byte {
	t.Helper()
	p0, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// ReadFile reads the file named by filename and returns the contents.
// A successful call returns err == nil, not err == EOF. Because ReadFile
// reads the whole file, it does not treat an EOF from Read as an error
// to be reported.
//
// Deprecated: As of Go 1.16, this function simply calls [os.ReadFile].
func ReadFile(t testing.TB, filename string) []byte {
	t.Helper()
	p0, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// WriteFile writes data to a file named by filename.
// If the file does not exist, WriteFile creates it with permissions perm
// (before umask); otherwise WriteFile truncates it before writing, without changing permissions.
//
// Deprecated: As of Go 1.16, this function simply calls [os.WriteFile].
func WriteFile(t testing.TB, filename string, data []byte, perm fs.FileMode) {
	t.Helper()
	err := ioutil.WriteFile(filename, data, perm)
	if err != nil {
		t.Fatal(err)
	}
	return
}

// ReadDir reads the directory named by dirname and returns
// a list of fs.FileInfo for the directory's contents,
// sorted by filename. If an error occurs reading the directory,
// ReadDir returns no directory entries along with the error.
//
// Deprecated: As of Go 1.16, [os.ReadDir] is a more efficient and correct choice:
// it returns a list of [fs.DirEntry] instead of [fs.FileInfo],
// and it returns partial results in the case of an error
// midway through reading a directory.
//
// If you must continue obtaining a list of [fs.FileInfo], you still can:
//
//	entries, err := os.ReadDir(dirname)
//	if err != nil { ... }
//	infos := make([]fs.FileInfo, 0, len(entries))
//	for _, entry := range entries {
//		info, err := entry.Info()
//		if err != nil { ... }
//		infos = append(infos, info)
//	}
func ReadDir(t testing.TB, dirname string) []fs.FileInfo {
	t.Helper()
	p0, err := ioutil.ReadDir(dirname)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
