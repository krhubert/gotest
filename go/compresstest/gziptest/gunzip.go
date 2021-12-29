package gziptest

import (
	"compress/gzip"
	"io"
	"testing"
)

// NewReader creates a new [Reader] reading the given reader.
// If r does not also implement [io.ByteReader],
// the decompressor may read more data than necessary from r.
//
// It is the caller's responsibility to call Close on the [Reader] when done.
//
// The [Reader.Header] fields will be valid in the [Reader] returned.
func NewReader(t testing.TB, r io.Reader) *gzip.Reader {
	t.Helper()
	p0, err := gzip.NewReader(r)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
