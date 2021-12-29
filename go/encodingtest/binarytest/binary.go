package binarytest

import (
	"encoding/binary"
	"io"
	"testing"
)

// Read reads structured binary data from r into data.
// Data must be a pointer to a fixed-size value or a slice
// of fixed-size values.
// Bytes read from r are decoded using the specified byte order
// and written to successive fields of the data.
// When decoding boolean values, a zero byte is decoded as false, and
// any other non-zero byte is decoded as true.
// When reading into structs, the field data for fields with
// blank (_) field names is skipped; i.e., blank field names
// may be used for padding.
// When reading into a struct, all non-blank fields must be exported
// or Read may panic.
//
// The error is [io.EOF] only if no bytes were read.
// If an [io.EOF] happens after reading some but not all the bytes,
// Read returns [io.ErrUnexpectedEOF].
func Read(t testing.TB, r io.Reader, order binary.ByteOrder, data any) {
	t.Helper()
	err := binary.Read(r, order, data)
	if err != nil {
		t.Fatal(err)
	}
	return
}

// Decode decodes binary data from buf into data according to
// the given byte order.
// It returns an error if buf is too small, otherwise the number of
// bytes consumed from buf.
func Decode(t testing.TB, buf []byte, order binary.ByteOrder, data any) int {
	t.Helper()
	p0, err := binary.Decode(buf, order, data)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// Write writes the binary representation of data into w.
// Data must be a fixed-size value or a slice of fixed-size
// values, or a pointer to such data.
// Boolean values encode as one byte: 1 for true, and 0 for false.
// Bytes written to w are encoded using the specified byte order
// and read from successive fields of the data.
// When writing structs, zero values are written for fields
// with blank (_) field names.
func Write(t testing.TB, w io.Writer, order binary.ByteOrder, data any) {
	t.Helper()
	err := binary.Write(w, order, data)
	if err != nil {
		t.Fatal(err)
	}
	return
}

// Encode encodes the binary representation of data into buf according to
// the given byte order.
// It returns an error if buf is too small, otherwise the number of
// bytes written into buf.
func Encode(t testing.TB, buf []byte, order binary.ByteOrder, data any) int {
	t.Helper()
	p0, err := binary.Encode(buf, order, data)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// Append appends the binary representation of data to buf.
// buf may be nil, in which case a new buffer will be allocated.
// See [Write] on which data are acceptable.
// It returns the (possibily extended) buffer containing data or an error.
func Append(t testing.TB, buf []byte, order binary.ByteOrder, data any) []byte {
	t.Helper()
	p0, err := binary.Append(buf, order, data)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
