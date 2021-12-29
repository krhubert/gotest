package typestest

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

// Eval returns the type and, if constant, the value for the
// expression expr, evaluated at position pos of package pkg,
// which must have been derived from type-checking an AST with
// complete position information relative to the provided file
// set.
//
// The meaning of the parameters fset, pkg, and pos is the
// same as in [CheckExpr]. An error is returned if expr cannot
// be parsed successfully, or the resulting expr AST cannot be
// type-checked.
func Eval(t testing.TB, fset *token.FileSet, pkg *types.Package, pos token.Pos, expr string) types.TypeAndValue {
	t.Helper()
	p0, err := types.Eval(fset, pkg, pos, expr)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// CheckExpr type checks the expression expr as if it had appeared at position
// pos of package pkg. [Type] information about the expression is recorded in
// info. The expression may be an identifier denoting an uninstantiated generic
// function or type.
//
// If pkg == nil, the [Universe] scope is used and the provided
// position pos is ignored. If pkg != nil, and pos is invalid,
// the package scope is used. Otherwise, pos must belong to the
// package.
//
// An error is returned if pos is not within the package or
// if the node cannot be type-checked.
//
// Note: [Eval] and CheckExpr should not be used instead of running Check
// to compute types and values, but in addition to Check, as these
// functions ignore the context in which an expression is used (e.g., an
// assignment). Thus, top-level untyped constants will return an
// untyped type rather than the respective context-specific type.
func CheckExpr(t testing.TB, fset *token.FileSet, pkg *types.Package, pos token.Pos, expr ast.Expr, info *types.Info) {
	t.Helper()
	err := types.CheckExpr(fset, pkg, pos, expr, info)
	if err != nil {
		t.Fatal(err)
	}
	return
}
