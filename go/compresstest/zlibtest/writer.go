package zlibtest

import (
	"compress/zlib"
	"io"
	"testing"
)

// NewWriterLevel is like NewWriter but specifies the compression level instead
// of assuming DefaultCompression.
//
// The compression level can be DefaultCompression, NoCompression, HuffmanOnly
// or any integer value between BestSpeed and BestCompression inclusive.
// The error returned will be nil if the level is valid.
func NewWriterLevel(t testing.TB, w io.Writer, level int) *zlib.Writer {
	t.Helper()
	p0, err := zlib.NewWriterLevel(w, level)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// NewWriterLevelDict is like NewWriterLevel but specifies a dictionary to
// compress with.
//
// The dictionary may be nil. If not, its contents should not be modified until
// the Writer is closed.
func NewWriterLevelDict(t testing.TB, w io.Writer, level int, dict []byte) *zlib.Writer {
	t.Helper()
	p0, err := zlib.NewWriterLevelDict(w, level, dict)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
