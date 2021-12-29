package strconvtest

import (
	"strconv"
	"testing"
)

// UnquoteChar decodes the first character or byte in the escaped string
// or character literal represented by the string s.
// It returns four values:
//
//  1. value, the decoded Unicode code point or byte value;
//  2. multibyte, a boolean indicating whether the decoded character requires a multibyte UTF-8 representation;
//  3. tail, the remainder of the string after the character; and
//  4. an error that will be nil if the character is syntactically valid.
//
// The second argument, quote, specifies the type of literal being parsed
// and therefore which escaped quote character is permitted.
// If set to a single quote, it permits the sequence \' and disallows unescaped '.
// If set to a double quote, it permits \" and disallows unescaped ".
// If set to zero, it does not permit either escape and allows both quote characters to appear unescaped.
func UnquoteChar(t testing.TB, s string, quote byte) (value rune, multibyte bool, tail string) {
	t.Helper()
	p0, p1, p2, err := strconv.UnquoteChar(s, quote)
	if err != nil {
		t.Fatal(err)
	}
	return p0, p1, p2
}

// QuotedPrefix returns the quoted string (as understood by [Unquote]) at the prefix of s.
// If s does not start with a valid quoted string, QuotedPrefix returns an error.
func QuotedPrefix(t testing.TB, s string) string {
	t.Helper()
	p0, err := strconv.QuotedPrefix(s)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// Unquote interprets s as a single-quoted, double-quoted,
// or backquoted Go string literal, returning the string value
// that s quotes.  (If s is single-quoted, it would be a Go
// character literal; Unquote returns the corresponding
// one-character string.)
func Unquote(t testing.TB, s string) string {
	t.Helper()
	p0, err := strconv.Unquote(s)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
