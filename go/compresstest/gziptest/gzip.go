package gziptest

import (
	"compress/gzip"
	"io"
	"testing"
)

// NewWriterLevel is like [NewWriter] but specifies the compression level instead
// of assuming [DefaultCompression].
//
// The compression level can be [DefaultCompression], [NoCompression], [HuffmanOnly]
// or any integer value between [BestSpeed] and [BestCompression] inclusive.
// The error returned will be nil if the level is valid.
func NewWriterLevel(t testing.TB, w io.Writer, level int) *gzip.Writer {
	t.Helper()
	p0, err := gzip.NewWriterLevel(w, level)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
