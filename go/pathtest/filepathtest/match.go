package filepathtest

import (
	"path/filepath"
	"testing"
)

// Match reports whether name matches the shell file name pattern.
// The pattern syntax is:
//
//	pattern:
//		{ term }
//	term:
//		'*'         matches any sequence of non-Separator characters
//		'?'         matches any single non-Separator character
//		'[' [ '^' ] { character-range } ']'
//		            character class (must be non-empty)
//		c           matches character c (c != '*', '?', '\\', '[')
//		'\\' c      matches character c
//
//	character-range:
//		c           matches character c (c != '\\', '-', ']')
//		'\\' c      matches character c
//		lo '-' hi   matches character c for lo <= c <= hi
//
// Match requires pattern to match all of name, not just a substring.
// The only possible returned error is [ErrBadPattern], when pattern
// is malformed.
//
// On Windows, escaping is disabled. Instead, '\\' is treated as
// path separator.
func Match(t testing.TB, pattern, name string) (matched bool) {
	t.Helper()
	p0, err := filepath.Match(pattern, name)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// Glob returns the names of all files matching pattern or nil
// if there is no matching file. The syntax of patterns is the same
// as in [Match]. The pattern may describe hierarchical names such as
// /usr/*/bin/ed (assuming the [Separator] is '/').
//
// Glob ignores file system errors such as I/O errors reading directories.
// The only possible returned error is [ErrBadPattern], when pattern
// is malformed.
func Glob(t testing.TB, pattern string) (matches []string) {
	t.Helper()
	p0, err := filepath.Glob(pattern)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
