package petest

import (
	"debug/pe"
	"io"
	"testing"
)

// Open opens the named file using [os.Open] and prepares it for use as a PE binary.
func Open(t testing.TB, name string) *pe.File {
	t.Helper()
	p0, err := pe.Open(name)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// NewFile creates a new [File] for accessing a PE binary in an underlying reader.
func NewFile(t testing.TB, r io.ReaderAt) *pe.File {
	t.Helper()
	p0, err := pe.NewFile(r)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
