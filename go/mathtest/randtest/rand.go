package randtest

import (
	"crypto/rand"
	"testing"
)

// Read generates len(p) random bytes from the default [Source] and
// writes them into p. It always returns len(p) and a nil error.
// Read, unlike the [Rand.Read] method, is safe for concurrent use.
//
// Deprecated: For almost all use cases, [crypto/rand.Read] is more appropriate.
// If a deterministic source is required, use [math/rand/v2.ChaCha8.Read].
func Read(t testing.TB, p []byte) (n int) {
	t.Helper()
	p0, err := rand.Read(p)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
