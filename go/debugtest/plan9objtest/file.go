package plan9objtest

import (
	"debug/plan9obj"
	"io"
	"testing"
)

// Open opens the named file using [os.Open] and prepares it for use as a Plan 9 a.out binary.
func Open(t testing.TB, name string) *plan9obj.File {
	t.Helper()
	p0, err := plan9obj.Open(name)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// NewFile creates a new [File] for accessing a Plan 9 binary in an underlying reader.
// The Plan 9 binary is expected to start at position 0 in the ReaderAt.
func NewFile(t testing.TB, r io.ReaderAt) *plan9obj.File {
	t.Helper()
	p0, err := plan9obj.NewFile(r)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
