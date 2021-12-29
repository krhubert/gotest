package doctest

import (
	"go/ast"
	"go/doc"
	"go/token"
	"testing"
)

// NewFromFiles computes documentation for a package.
//
// The package is specified by a list of *ast.Files and corresponding
// file set, which must not be nil.
// NewFromFiles uses all provided files when computing documentation,
// so it is the caller's responsibility to provide only the files that
// match the desired build context. "go/build".Context.MatchFile can
// be used for determining whether a file matches a build context with
// the desired GOOS and GOARCH values, and other build constraints.
// The import path of the package is specified by importPath.
//
// Examples found in _test.go files are associated with the corresponding
// type, function, method, or the package, based on their name.
// If the example has a suffix in its name, it is set in the
// [Example.Suffix] field. [Examples] with malformed names are skipped.
//
// Optionally, a single extra argument of type [Mode] can be provided to
// control low-level aspects of the documentation extraction behavior.
//
// NewFromFiles takes ownership of the AST files and may edit them,
// unless the PreserveAST Mode bit is on.
func NewFromFiles(t testing.TB, fset *token.FileSet, files []*ast.File, importPath string, opts ...any) *doc.Package {
	t.Helper()
	p0, err := doc.NewFromFiles(fset, files, importPath, opts)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
