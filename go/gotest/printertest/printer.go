package printertest

import (
	"go/printer"
	"go/token"
	"io"
	"testing"
)

// Fprint "pretty-prints" an AST node to output.
// It calls [Config.Fprint] with default settings.
// Note that gofmt uses tabs for indentation but spaces for alignment;
// use format.Node (package go/format) for output that matches gofmt.
func Fprint(t testing.TB, output io.Writer, fset *token.FileSet, node any) {
	t.Helper()
	err := printer.Fprint(output, fset, node)
	if err != nil {
		t.Fatal(err)
	}
	return
}
