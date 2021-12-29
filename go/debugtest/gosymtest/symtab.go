package gosymtest

import (
	"debug/gosym"
	"testing"
)

// NewTable decodes the Go symbol table (the ".gosymtab" section in ELF),
// returning an in-memory representation.
// Starting with Go 1.3, the Go symbol table no longer includes symbol data.
func NewTable(t testing.TB, symtab []byte, pcln *gosym.LineTable) *gosym.Table {
	t.Helper()
	p0, err := gosym.NewTable(symtab, pcln)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
