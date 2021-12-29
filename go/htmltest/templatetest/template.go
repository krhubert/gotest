package templatetest

import (
	"html/template"
	"io/fs"
	"testing"
)

// ParseFiles creates a new [Template] and parses the template definitions from
// the named files. The returned template's name will have the (base) name and
// (parsed) contents of the first file. There must be at least one file.
// If an error occurs, parsing stops and the returned [*Template] is nil.
//
// When parsing multiple files with the same name in different directories,
// the last one mentioned will be the one that results.
// For instance, ParseFiles("a/foo", "b/foo") stores "b/foo" as the template
// named "foo", while "a/foo" is unavailable.
func ParseFiles(t testing.TB, filenames ...string) *template.Template {
	t.Helper()
	p0, err := template.ParseFiles(filenames...)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// ParseGlob creates a new [Template] and parses the template definitions from
// the files identified by the pattern. The files are matched according to the
// semantics of filepath.Match, and the pattern must match at least one file.
// The returned template will have the (base) name and (parsed) contents of the
// first file matched by the pattern. ParseGlob is equivalent to calling
// [ParseFiles] with the list of files matched by the pattern.
//
// When parsing multiple files with the same name in different directories,
// the last one mentioned will be the one that results.
func ParseGlob(t testing.TB, pattern string) *template.Template {
	t.Helper()
	p0, err := template.ParseGlob(pattern)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// ParseFS is like [ParseFiles] or [ParseGlob] but reads from the file system fs
// instead of the host operating system's file system.
// It accepts a list of glob patterns.
// (Note that most file names serve as glob patterns matching only themselves.)
func ParseFS(t testing.TB, fs fs.FS, patterns ...string) *template.Template {
	t.Helper()
	p0, err := template.ParseFS(fs, patterns...)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
