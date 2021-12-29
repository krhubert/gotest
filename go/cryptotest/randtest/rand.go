package randtest

import (
	"crypto/rand"
	"testing"
)

// Read is a helper function that calls Reader.Read using io.ReadFull.
// On return, n == len(b) if and only if err == nil.
func Read(t testing.TB, b []byte) (n int) {
	t.Helper()
	p0, err := rand.Read(b)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
