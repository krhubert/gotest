package buildtest

import (
	"go/build"
	"testing"
)

// Import is shorthand for Default.Import.
func Import(t testing.TB, path, srcDir string, mode build.ImportMode) *build.Package {
	t.Helper()
	p0, err := build.Import(path, srcDir, mode)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// ImportDir is shorthand for Default.ImportDir.
func ImportDir(t testing.TB, dir string, mode build.ImportMode) *build.Package {
	t.Helper()
	p0, err := build.ImportDir(dir, mode)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// ArchChar returns "?" and an error.
// In earlier versions of Go, the returned string was used to derive
// the compiler and linker tool names, the default object file suffix,
// and the default linker output name. As of Go 1.5, those strings
// no longer vary by architecture; they are compile, link, .o, and a.out, respectively.
func ArchChar(t testing.TB, goarch string) string {
	t.Helper()
	p0, err := build.ArchChar(goarch)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
