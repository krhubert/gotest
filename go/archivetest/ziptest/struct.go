package ziptest

import (
	"archive/zip"
	"io/fs"
	"testing"
)

// FileInfoHeader creates a partially-populated [FileHeader] from an
// fs.FileInfo.
// Because fs.FileInfo's Name method returns only the base name of
// the file it describes, it may be necessary to modify the Name field
// of the returned header to provide the full path name of the file.
// If compression is desired, callers should set the FileHeader.Method
// field; it is unset by default.
func FileInfoHeader(t testing.TB, fi fs.FileInfo) *zip.FileHeader {
	t.Helper()
	p0, err := zip.FileInfoHeader(fi)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
