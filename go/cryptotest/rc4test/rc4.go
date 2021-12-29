package rc4test

import (
	"crypto/rc4"
	"testing"
)

// NewCipher creates and returns a new [Cipher]. The key argument should be the
// RC4 key, at least 1 byte and at most 256 bytes.
func NewCipher(t testing.TB, key []byte) *rc4.Cipher {
	t.Helper()
	p0, err := rc4.NewCipher(key)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
