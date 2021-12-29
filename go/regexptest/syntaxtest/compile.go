package syntaxtest

import (
	"regexp/syntax"
	"testing"
)

// Compile compiles the regexp into a program to be executed.
// The regexp should have been simplified already (returned from re.Simplify).
func Compile(t testing.TB, re *syntax.Regexp) *syntax.Prog {
	t.Helper()
	p0, err := syntax.Compile(re)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
