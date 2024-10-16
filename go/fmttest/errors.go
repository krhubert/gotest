package fmttest

import (
	"fmt"
	"testing"
)

// Errorf formats according to a format specifier and returns the string as a
// value that satisfies error.
//
// If the format specifier includes a %w verb with an error operand,
// the returned error will implement an Unwrap method returning the operand.
// If there is more than one %w verb, the returned error will implement an
// Unwrap method returning a []error containing all the %w operands in the
// order they appear in the arguments.
// It is invalid to supply the %w verb with an operand that does not implement
// the error interface. The %w verb is otherwise a synonym for %v.
func Errorf(t testing.TB, format string, a ...any) {
	t.Helper()
	err := fmt.Errorf(format, a)
	if err != nil {
		t.Fatal(err)
	}
	return
}
