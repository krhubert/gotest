package asttest

import (
	"go/ast"
	"go/token"
	"testing"
)

// NewPackage creates a new [Package] node from a set of [File] nodes. It resolves
// unresolved identifiers across files and updates each file's Unresolved list
// accordingly. If a non-nil importer and universe scope are provided, they are
// used to resolve identifiers not declared in any of the package files. Any
// remaining unresolved identifiers are reported as undeclared. If the files
// belong to different packages, one package name is selected and files with
// different package names are reported and then ignored.
// The result is a package node and a [scanner.ErrorList] if there were errors.
//
// Deprecated: use the type checker [go/types] instead; see [Object].
func NewPackage(t testing.TB, fset *token.FileSet, files map[string]*ast.File, importer ast.Importer, universe *ast.Scope) *ast.Package {
	t.Helper()
	p0, err := ast.NewPackage(fset, files, importer, universe)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
