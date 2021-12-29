package constrainttest

import (
	"go/build/constraint"
	"testing"
)

// Parse parses a single build constraint line of the form “//go:build ...” or “// +build ...”
// and returns the corresponding boolean expression.
func Parse(t testing.TB, line string) constraint.Expr {
	t.Helper()
	p0, err := constraint.Parse(line)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// PlusBuildLines returns a sequence of “// +build” lines that evaluate to the build expression x.
// If the expression is too complex to convert directly to “// +build” lines, PlusBuildLines returns an error.
func PlusBuildLines(t testing.TB, x constraint.Expr) []string {
	t.Helper()
	p0, err := constraint.PlusBuildLines(x)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
