package buildinfotest

import (
	"debug/buildinfo"
	"io"
	"testing"
)

// ReadFile returns build information embedded in a Go binary
// file at the given path. Most information is only available for binaries built
// with module support.
func ReadFile(t testing.TB, name string) (info *buildinfo.BuildInfo) {
	t.Helper()
	p0, err := buildinfo.ReadFile(name)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// Read returns build information embedded in a Go binary file
// accessed through the given ReaderAt. Most information is only available for
// binaries built with module support.
func Read(t testing.TB, r io.ReaderAt) *buildinfo.BuildInfo {
	t.Helper()
	p0, err := buildinfo.Read(r)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
