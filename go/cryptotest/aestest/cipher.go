package aestest

import (
	"crypto/aes"
	"crypto/cipher"
	"testing"
)

// NewCipher creates and returns a new [cipher.Block].
// The key argument should be the AES key,
// either 16, 24, or 32 bytes to select
// AES-128, AES-192, or AES-256.
func NewCipher(t testing.TB, key []byte) cipher.Block {
	t.Helper()
	p0, err := aes.NewCipher(key)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
