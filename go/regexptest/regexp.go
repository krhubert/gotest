package regexptest

import (
	"io"
	"regexp"
	"testing"
)

// Compile parses a regular expression and returns, if successful,
// a [Regexp] object that can be used to match against text.
//
// When matching against text, the regexp returns a match that
// begins as early as possible in the input (leftmost), and among those
// it chooses the one that a backtracking search would have found first.
// This so-called leftmost-first matching is the same semantics
// that Perl, Python, and other implementations use, although this
// package implements it without the expense of backtracking.
// For POSIX leftmost-longest matching, see [CompilePOSIX].
func Compile(t testing.TB, expr string) *regexp.Regexp {
	t.Helper()
	p0, err := regexp.Compile(expr)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// CompilePOSIX is like [Compile] but restricts the regular expression
// to POSIX ERE (egrep) syntax and changes the match semantics to
// leftmost-longest.
//
// That is, when matching against text, the regexp returns a match that
// begins as early as possible in the input (leftmost), and among those
// it chooses a match that is as long as possible.
// This so-called leftmost-longest matching is the same semantics
// that early regular expression implementations used and that POSIX
// specifies.
//
// However, there can be multiple leftmost-longest matches, with different
// submatch choices, and here this package diverges from POSIX.
// Among the possible leftmost-longest matches, this package chooses
// the one that a backtracking search would have found first, while POSIX
// specifies that the match be chosen to maximize the length of the first
// subexpression, then the second, and so on from left to right.
// The POSIX rule is computationally prohibitive and not even well-defined.
// See https://swtch.com/~rsc/regexp/regexp2.html#posix for details.
func CompilePOSIX(t testing.TB, expr string) *regexp.Regexp {
	t.Helper()
	p0, err := regexp.CompilePOSIX(expr)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// MatchReader reports whether the text returned by the [io.RuneReader]
// contains any match of the regular expression pattern.
// More complicated queries need to use [Compile] and the full [Regexp] interface.
func MatchReader(t testing.TB, pattern string, r io.RuneReader) (matched bool) {
	t.Helper()
	p0, err := regexp.MatchReader(pattern, r)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// MatchString reports whether the string s
// contains any match of the regular expression pattern.
// More complicated queries need to use [Compile] and the full [Regexp] interface.
func MatchString(t testing.TB, pattern string, s string) (matched bool) {
	t.Helper()
	p0, err := regexp.MatchString(pattern, s)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// Match reports whether the byte slice b
// contains any match of the regular expression pattern.
// More complicated queries need to use [Compile] and the full [Regexp] interface.
func Match(t testing.TB, pattern string, b []byte) (matched bool) {
	t.Helper()
	p0, err := regexp.Match(pattern, b)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
