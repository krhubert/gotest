package parsetest

import (
	"testing"
	"text/template/parse"
)

// Parse returns a map from template name to [Tree], created by parsing the
// templates described in the argument string. The top-level template will be
// given the specified name. If an error is encountered, parsing stops and an
// empty map is returned with the error.
func Parse(t testing.TB, name, text, leftDelim, rightDelim string, funcs ...map[string]any) map[string]*parse.Tree {
	t.Helper()
	p0, err := parse.Parse(name, text, leftDelim, rightDelim, funcs...)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
