package fmttest

import (
	"fmt"
	"io"
	"testing"
)

// Scan scans text read from standard input, storing successive
// space-separated values into successive arguments. Newlines count
// as space. It returns the number of items successfully scanned.
// If that is less than the number of arguments, err will report why.
func Scan(t testing.TB, a ...any) (n int) {
	t.Helper()
	p0, err := fmt.Scan(a)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// Scanln is similar to [Scan], but stops scanning at a newline and
// after the final item there must be a newline or EOF.
func Scanln(t testing.TB, a ...any) (n int) {
	t.Helper()
	p0, err := fmt.Scanln(a)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// Scanf scans text read from standard input, storing successive
// space-separated values into successive arguments as determined by
// the format. It returns the number of items successfully scanned.
// If that is less than the number of arguments, err will report why.
// Newlines in the input must match newlines in the format.
// The one exception: the verb %c always scans the next rune in the
// input, even if it is a space (or tab etc.) or newline.
func Scanf(t testing.TB, format string, a ...any) (n int) {
	t.Helper()
	p0, err := fmt.Scanf(format, a)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// Sscan scans the argument string, storing successive space-separated
// values into successive arguments. Newlines count as space. It
// returns the number of items successfully scanned. If that is less
// than the number of arguments, err will report why.
func Sscan(t testing.TB, str string, a ...any) (n int) {
	t.Helper()
	p0, err := fmt.Sscan(str, a)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// Sscanln is similar to [Sscan], but stops scanning at a newline and
// after the final item there must be a newline or EOF.
func Sscanln(t testing.TB, str string, a ...any) (n int) {
	t.Helper()
	p0, err := fmt.Sscanln(str, a)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// Sscanf scans the argument string, storing successive space-separated
// values into successive arguments as determined by the format. It
// returns the number of items successfully parsed.
// Newlines in the input must match newlines in the format.
func Sscanf(t testing.TB, str string, format string, a ...any) (n int) {
	t.Helper()
	p0, err := fmt.Sscanf(str, format, a)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// Fscan scans text read from r, storing successive space-separated
// values into successive arguments. Newlines count as space. It
// returns the number of items successfully scanned. If that is less
// than the number of arguments, err will report why.
func Fscan(t testing.TB, r io.Reader, a ...any) (n int) {
	t.Helper()
	p0, err := fmt.Fscan(r, a)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// Fscanln is similar to [Fscan], but stops scanning at a newline and
// after the final item there must be a newline or EOF.
func Fscanln(t testing.TB, r io.Reader, a ...any) (n int) {
	t.Helper()
	p0, err := fmt.Fscanln(r, a)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// Fscanf scans text read from r, storing successive space-separated
// values into successive arguments as determined by the format. It
// returns the number of items successfully parsed.
// Newlines in the input must match newlines in the format.
func Fscanf(t testing.TB, r io.Reader, format string, a ...any) (n int) {
	t.Helper()
	p0, err := fmt.Fscanf(r, format, a)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
