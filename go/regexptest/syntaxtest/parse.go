package syntaxtest

import (
	"regexp/syntax"
	"testing"
)

// Parse parses a regular expression string s, controlled by the specified
// Flags, and returns a regular expression parse tree. The syntax is
// described in the top-level comment.
func Parse(t testing.TB, s string, flags syntax.Flags) *syntax.Regexp {
	t.Helper()
	p0, err := syntax.Parse(s, flags)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
