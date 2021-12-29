package fmttest

import (
	"fmt"
	"io"
	"testing"
)

// Fprintf formats according to a format specifier and writes to w.
// It returns the number of bytes written and any write error encountered.
func Fprintf(t testing.TB, w io.Writer, format string, a ...any) (n int) {
	t.Helper()
	p0, err := fmt.Fprintf(w, format, a)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func Printf(t testing.TB, format string, a ...any) (n int) {
	t.Helper()
	p0, err := fmt.Printf(format, a)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// Fprint formats using the default formats for its operands and writes to w.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func Fprint(t testing.TB, w io.Writer, a ...any) (n int) {
	t.Helper()
	p0, err := fmt.Fprint(w, a)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// Print formats using the default formats for its operands and writes to standard output.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func Print(t testing.TB, a ...any) (n int) {
	t.Helper()
	p0, err := fmt.Print(a)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// Fprintln formats using the default formats for its operands and writes to w.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func Fprintln(t testing.TB, w io.Writer, a ...any) (n int) {
	t.Helper()
	p0, err := fmt.Fprintln(w, a)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func Println(t testing.TB, a ...any) (n int) {
	t.Helper()
	p0, err := fmt.Println(a)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
