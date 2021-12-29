package zlibtest

import (
	"compress/zlib"
	"io"
	"testing"
)

// NewReader creates a new ReadCloser.
// Reads from the returned ReadCloser read and decompress data from r.
// If r does not implement [io.ByteReader], the decompressor may read more
// data than necessary from r.
// It is the caller's responsibility to call Close on the ReadCloser when done.
//
// The [io.ReadCloser] returned by NewReader also implements [Resetter].
func NewReader(t testing.TB, r io.Reader) io.ReadCloser {
	t.Helper()
	p0, err := zlib.NewReader(r)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// NewReaderDict is like [NewReader] but uses a preset dictionary.
// NewReaderDict ignores the dictionary if the compressed data does not refer to it.
// If the compressed data refers to a different dictionary, NewReaderDict returns [ErrDictionary].
//
// The ReadCloser returned by NewReaderDict also implements [Resetter].
func NewReaderDict(t testing.TB, r io.Reader, dict []byte) io.ReadCloser {
	t.Helper()
	p0, err := zlib.NewReaderDict(r, dict)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
