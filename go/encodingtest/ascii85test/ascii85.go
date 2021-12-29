package ascii85test

import (
	"encoding/ascii85"
	"testing"
)

// Decode decodes src into dst, returning both the number
// of bytes written to dst and the number consumed from src.
// If src contains invalid ascii85 data, Decode will return the
// number of bytes successfully written and a [CorruptInputError].
// Decode ignores space and control characters in src.
// Often, ascii85-encoded data is wrapped in <~ and ~> symbols.
// Decode expects these to have been stripped by the caller.
//
// If flush is true, Decode assumes that src represents the
// end of the input stream and processes it completely rather
// than wait for the completion of another 32-bit block.
//
// [NewDecoder] wraps an [io.Reader] interface around Decode.
func Decode(t testing.TB, dst, src []byte, flush bool) (ndst, nsrc int) {
	t.Helper()
	p0, p1, err := ascii85.Decode(dst, src, flush)
	if err != nil {
		t.Fatal(err)
	}
	return p0, p1
}
