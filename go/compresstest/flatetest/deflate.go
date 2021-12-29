package flatetest

import (
	"compress/flate"
	"io"
	"testing"
)

// NewWriter returns a new [Writer] compressing data at the given level.
// Following zlib, levels range from 1 ([BestSpeed]) to 9 ([BestCompression]);
// higher levels typically run slower but compress more. Level 0
// ([NoCompression]) does not attempt any compression; it only adds the
// necessary DEFLATE framing.
// Level -1 ([DefaultCompression]) uses the default compression level.
// Level -2 ([HuffmanOnly]) will use Huffman compression only, giving
// a very fast compression for all types of input, but sacrificing considerable
// compression efficiency.
//
// If level is in the range [-2, 9] then the error returned will be nil.
// Otherwise the error returned will be non-nil.
func NewWriter(t testing.TB, w io.Writer, level int) *flate.Writer {
	t.Helper()
	p0, err := flate.NewWriter(w, level)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// NewWriterDict is like [NewWriter] but initializes the new
// [Writer] with a preset dictionary. The returned [Writer] behaves
// as if the dictionary had been written to it without producing
// any compressed output. The compressed data written to w
// can only be decompressed by a [Reader] initialized with the
// same dictionary.
func NewWriterDict(t testing.TB, w io.Writer, level int, dict []byte) *flate.Writer {
	t.Helper()
	p0, err := flate.NewWriterDict(w, level, dict)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
