package ziptest

import (
	"archive/zip"
	"io"
	"testing"
)

// OpenReader will open the Zip file specified by name and return a ReadCloser.
//
// If any file inside the archive uses a non-local name
// (as defined by [filepath.IsLocal]) or a name containing backslashes
// and the GODEBUG environment variable contains `zipinsecurepath=0`,
// OpenReader returns the reader with an ErrInsecurePath error.
// A future version of Go may introduce this behavior by default.
// Programs that want to accept non-local names can ignore
// the ErrInsecurePath error and use the returned reader.
func OpenReader(t testing.TB, name string) *zip.ReadCloser {
	t.Helper()
	p0, err := zip.OpenReader(name)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// NewReader returns a new [Reader] reading from r, which is assumed to
// have the given size in bytes.
//
// If any file inside the archive uses a non-local name
// (as defined by [filepath.IsLocal]) or a name containing backslashes
// and the GODEBUG environment variable contains `zipinsecurepath=0`,
// NewReader returns the reader with an [ErrInsecurePath] error.
// A future version of Go may introduce this behavior by default.
// Programs that want to accept non-local names can ignore
// the [ErrInsecurePath] error and use the returned reader.
func NewReader(t testing.TB, r io.ReaderAt, size int64) *zip.Reader {
	t.Helper()
	p0, err := zip.NewReader(r, size)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
