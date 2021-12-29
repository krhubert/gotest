package parsertest

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"testing"
)

// ParseFile parses the source code of a single Go source file and returns
// the corresponding [ast.File] node. The source code may be provided via
// the filename of the source file, or via the src parameter.
//
// If src != nil, ParseFile parses the source from src and the filename is
// only used when recording position information. The type of the argument
// for the src parameter must be string, []byte, or [io.Reader].
// If src == nil, ParseFile parses the file specified by filename.
//
// The mode parameter controls the amount of source text parsed and
// other optional parser functionality. If the [SkipObjectResolution]
// mode bit is set (recommended), the object resolution phase of
// parsing will be skipped, causing File.Scope, File.Unresolved, and
// all Ident.Obj fields to be nil. Those fields are deprecated; see
// [ast.Object] for details.
//
// Position information is recorded in the file set fset, which must not be
// nil.
//
// If the source couldn't be read, the returned AST is nil and the error
// indicates the specific failure. If the source was read but syntax
// errors were found, the result is a partial AST (with [ast.Bad]* nodes
// representing the fragments of erroneous source code). Multiple errors
// are returned via a scanner.ErrorList which is sorted by source position.
func ParseFile(t testing.TB, fset *token.FileSet, filename string, src any, mode parser.Mode) (f *ast.File) {
	t.Helper()
	p0, err := parser.ParseFile(fset, filename, src, mode)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// ParseDir calls [ParseFile] for all files with names ending in ".go" in the
// directory specified by path and returns a map of package name -> package
// AST with all the packages found.
//
// If filter != nil, only the files with [fs.FileInfo] entries passing through
// the filter (and ending in ".go") are considered. The mode bits are passed
// to [ParseFile] unchanged. Position information is recorded in fset, which
// must not be nil.
//
// If the directory couldn't be read, a nil map and the respective error are
// returned. If a parse error occurred, a non-nil but incomplete map and the
// first error encountered are returned.
func ParseDir(t testing.TB, fset *token.FileSet, path string, filter func(fs.FileInfo) bool, mode parser.Mode) (pkgs map[string]*ast.Package) {
	t.Helper()
	p0, err := parser.ParseDir(fset, path, filter, mode)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// ParseExprFrom is a convenience function for parsing an expression.
// The arguments have the same meaning as for [ParseFile], but the source must
// be a valid Go (type or value) expression. Specifically, fset must not
// be nil.
//
// If the source couldn't be read, the returned AST is nil and the error
// indicates the specific failure. If the source was read but syntax
// errors were found, the result is a partial AST (with [ast.Bad]* nodes
// representing the fragments of erroneous source code). Multiple errors
// are returned via a scanner.ErrorList which is sorted by source position.
func ParseExprFrom(t testing.TB, fset *token.FileSet, filename string, src any, mode parser.Mode) (expr ast.Expr) {
	t.Helper()
	p0, err := parser.ParseExprFrom(fset, filename, src, mode)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// ParseExpr is a convenience function for obtaining the AST of an expression x.
// The position information recorded in the AST is undefined. The filename used
// in error messages is the empty string.
//
// If syntax errors were found, the result is a partial AST (with [ast.Bad]* nodes
// representing the fragments of erroneous source code). Multiple errors are
// returned via a scanner.ErrorList which is sorted by source position.
func ParseExpr(t testing.TB, x string) ast.Expr {
	t.Helper()
	p0, err := parser.ParseExpr(x)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
