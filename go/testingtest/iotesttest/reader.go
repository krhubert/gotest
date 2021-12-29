package iotesttest

import (
	"io"
	"testing"
	"testing/iotest"
)

// TestReader tests that reading from r returns the expected file content.
// It does reads of different sizes, until EOF.
// If r implements [io.ReaderAt] or [io.Seeker], TestReader also checks
// that those operations behave as they should.
//
// If TestReader finds any misbehaviors, it returns an error reporting them.
// The error text may span multiple lines.
func TestReader(t testing.TB, r io.Reader, content []byte) {
	t.Helper()
	err := iotest.TestReader(r, content)
	if err != nil {
		t.Fatal(err)
	}
	return
}
