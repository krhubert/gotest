package machotest

import (
	"debug/macho"
	"io"
	"testing"
)

// Open opens the named file using [os.Open] and prepares it for use as a Mach-O binary.
func Open(t testing.TB, name string) *macho.File {
	t.Helper()
	p0, err := macho.Open(name)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// NewFile creates a new [File] for accessing a Mach-O binary in an underlying reader.
// The Mach-O binary is expected to start at position 0 in the ReaderAt.
func NewFile(t testing.TB, r io.ReaderAt) *macho.File {
	t.Helper()
	p0, err := macho.NewFile(r)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
