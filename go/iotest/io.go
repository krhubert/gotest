package iotest

import (
	"io"
	"testing"
)

// WriteString writes the contents of the string s to w, which accepts a slice of bytes.
// If w implements [StringWriter], [StringWriter.WriteString] is invoked directly.
// Otherwise, [Writer.Write] is called exactly once.
func WriteString(t testing.TB, w io.Writer, s string) (n int) {
	t.Helper()
	p0, err := io.WriteString(w, s)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// ReadAtLeast reads from r into buf until it has read at least min bytes.
// It returns the number of bytes copied and an error if fewer bytes were read.
// The error is EOF only if no bytes were read.
// If an EOF happens after reading fewer than min bytes,
// ReadAtLeast returns [ErrUnexpectedEOF].
// If min is greater than the length of buf, ReadAtLeast returns [ErrShortBuffer].
// On return, n >= min if and only if err == nil.
// If r returns an error having read at least min bytes, the error is dropped.
func ReadAtLeast(t testing.TB, r io.Reader, buf []byte, min int) (n int) {
	t.Helper()
	p0, err := io.ReadAtLeast(r, buf, min)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// ReadFull reads exactly len(buf) bytes from r into buf.
// It returns the number of bytes copied and an error if fewer bytes were read.
// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// ReadFull returns [ErrUnexpectedEOF].
// On return, n == len(buf) if and only if err == nil.
// If r returns an error having read at least len(buf) bytes, the error is dropped.
func ReadFull(t testing.TB, r io.Reader, buf []byte) (n int) {
	t.Helper()
	p0, err := io.ReadFull(r, buf)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// CopyN copies n bytes (or until an error) from src to dst.
// It returns the number of bytes copied and the earliest
// error encountered while copying.
// On return, written == n if and only if err == nil.
//
// If dst implements [ReaderFrom], the copy is implemented using it.
func CopyN(t testing.TB, dst io.Writer, src io.Reader, n int64) (written int64) {
	t.Helper()
	p0, err := io.CopyN(dst, src, n)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// Copy copies from src to dst until either EOF is reached
// on src or an error occurs. It returns the number of bytes
// copied and the first error encountered while copying, if any.
//
// A successful Copy returns err == nil, not err == EOF.
// Because Copy is defined to read from src until EOF, it does
// not treat an EOF from Read as an error to be reported.
//
// If src implements [WriterTo],
// the copy is implemented by calling src.WriteTo(dst).
// Otherwise, if dst implements [ReaderFrom],
// the copy is implemented by calling dst.ReadFrom(src).
func Copy(t testing.TB, dst io.Writer, src io.Reader) (written int64) {
	t.Helper()
	p0, err := io.Copy(dst, src)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// CopyBuffer is identical to Copy except that it stages through the
// provided buffer (if one is required) rather than allocating a
// temporary one. If buf is nil, one is allocated; otherwise if it has
// zero length, CopyBuffer panics.
//
// If either src implements [WriterTo] or dst implements [ReaderFrom],
// buf will not be used to perform the copy.
func CopyBuffer(t testing.TB, dst io.Writer, src io.Reader, buf []byte) (written int64) {
	t.Helper()
	p0, err := io.CopyBuffer(dst, src, buf)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// ReadAll reads from r until an error or EOF and returns the data it read.
// A successful call returns err == nil, not err == EOF. Because ReadAll is
// defined to read from src until EOF, it does not treat an EOF from Read
// as an error to be reported.
func ReadAll(t testing.TB, r io.Reader) []byte {
	t.Helper()
	p0, err := io.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
