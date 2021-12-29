package elftest

import (
	"debug/elf"
	"io"
	"testing"
)

// Open opens the named file using [os.Open] and prepares it for use as an ELF binary.
func Open(t testing.TB, name string) *elf.File {
	t.Helper()
	p0, err := elf.Open(name)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// NewFile creates a new [File] for accessing an ELF binary in an underlying reader.
// The ELF binary is expected to start at position 0 in the ReaderAt.
func NewFile(t testing.TB, r io.ReaderAt) *elf.File {
	t.Helper()
	p0, err := elf.NewFile(r)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
