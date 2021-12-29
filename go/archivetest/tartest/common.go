package tartest

import (
	"archive/tar"
	"io/fs"
	"testing"
)

// FileInfoHeader creates a partially-populated [Header] from fi.
// If fi describes a symlink, FileInfoHeader records link as the link target.
// If fi describes a directory, a slash is appended to the name.
//
// Since fs.FileInfo's Name method only returns the base name of
// the file it describes, it may be necessary to modify Header.Name
// to provide the full path name of the file.
//
// If fi implements [FileInfoNames]
// Header.Gname and Header.Uname
// are provided by the methods of the interface.
func FileInfoHeader(t testing.TB, fi fs.FileInfo, link string) *tar.Header {
	t.Helper()
	p0, err := tar.FileInfoHeader(fi, link)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
