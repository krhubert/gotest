package asttest

import (
	"go/ast"
	"go/token"
	"io"
	"testing"
)

// Fprint prints the (sub-)tree starting at AST node x to w.
// If fset != nil, position information is interpreted relative
// to that file set. Otherwise positions are printed as integer
// values (file set specific offsets).
//
// A non-nil [FieldFilter] f may be provided to control the output:
// struct fields for which f(fieldname, fieldvalue) is true are
// printed; all others are filtered from the output. Unexported
// struct fields are never printed.
func Fprint(t testing.TB, w io.Writer, fset *token.FileSet, x any, f ast.FieldFilter) {
	t.Helper()
	err := ast.Fprint(w, fset, x, f)
	if err != nil {
		t.Fatal(err)
	}
	return
}

// Print prints x to standard output, skipping nil fields.
// Print(fset, x) is the same as Fprint(os.Stdout, fset, x, NotNilFilter).
func Print(t testing.TB, fset *token.FileSet, x any) {
	t.Helper()
	err := ast.Print(fset, x)
	if err != nil {
		t.Fatal(err)
	}
	return
}
