package hextest

import (
	"encoding/hex"
	"testing"
)

// Decode decodes src into [DecodedLen](len(src)) bytes,
// returning the actual number of bytes written to dst.
//
// Decode expects that src contains only hexadecimal
// characters and that src has even length.
// If the input is malformed, Decode returns the number
// of bytes decoded before the error.
func Decode(t testing.TB, dst, src []byte) int {
	t.Helper()
	p0, err := hex.Decode(dst, src)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// AppendDecode appends the hexadecimally decoded src to dst
// and returns the extended buffer.
// If the input is malformed, it returns the partially decoded src and an error.
func AppendDecode(t testing.TB, dst, src []byte) []byte {
	t.Helper()
	p0, err := hex.AppendDecode(dst, src)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// DecodeString returns the bytes represented by the hexadecimal string s.
//
// DecodeString expects that src contains only hexadecimal
// characters and that src has even length.
// If the input is malformed, DecodeString returns
// the bytes decoded before the error.
func DecodeString(t testing.TB, s string) []byte {
	t.Helper()
	p0, err := hex.DecodeString(s)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
