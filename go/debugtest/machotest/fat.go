package machotest

import (
	"debug/macho"
	"io"
	"testing"
)

// NewFatFile creates a new [FatFile] for accessing all the Mach-O images in a
// universal binary. The Mach-O binary is expected to start at position 0 in
// the ReaderAt.
func NewFatFile(t testing.TB, r io.ReaderAt) *macho.FatFile {
	t.Helper()
	p0, err := macho.NewFatFile(r)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// OpenFat opens the named file using [os.Open] and prepares it for use as a Mach-O
// universal binary.
func OpenFat(t testing.TB, name string) *macho.FatFile {
	t.Helper()
	p0, err := macho.OpenFat(name)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
