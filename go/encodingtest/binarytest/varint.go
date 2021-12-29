package binarytest

import (
	"encoding/binary"
	"io"
	"testing"
)

// ReadUvarint reads an encoded unsigned integer from r and returns it as a uint64.
// The error is [io.EOF] only if no bytes were read.
// If an [io.EOF] happens after reading some but not all the bytes,
// ReadUvarint returns [io.ErrUnexpectedEOF].
func ReadUvarint(t testing.TB, r io.ByteReader) uint64 {
	t.Helper()
	p0, err := binary.ReadUvarint(r)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// ReadVarint reads an encoded signed integer from r and returns it as an int64.
// The error is [io.EOF] only if no bytes were read.
// If an [io.EOF] happens after reading some but not all the bytes,
// ReadVarint returns [io.ErrUnexpectedEOF].
func ReadVarint(t testing.TB, r io.ByteReader) int64 {
	t.Helper()
	p0, err := binary.ReadVarint(r)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
