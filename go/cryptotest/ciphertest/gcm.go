package ciphertest

import (
	"crypto/cipher"
	"testing"
)

// NewGCM returns the given 128-bit, block cipher wrapped in Galois Counter Mode
// with the standard nonce length.
//
// In general, the GHASH operation performed by this implementation of GCM is not constant-time.
// An exception is when the underlying [Block] was created by aes.NewCipher
// on systems with hardware support for AES. See the [crypto/aes] package documentation for details.
func NewGCM(t testing.TB, block cipher.Block) cipher.AEAD {
	t.Helper()
	p0, err := cipher.NewGCM(block)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// NewGCMWithNonceSize returns the given 128-bit, block cipher wrapped in Galois
// Counter Mode, which accepts nonces of the given length. The length must not
// be zero.
//
// Only use this function if you require compatibility with an existing
// cryptosystem that uses non-standard nonce lengths. All other users should use
// [NewGCM], which is faster and more resistant to misuse.
func NewGCMWithNonceSize(t testing.TB, block cipher.Block, size int) cipher.AEAD {
	t.Helper()
	p0, err := cipher.NewGCMWithNonceSize(block, size)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// NewGCMWithTagSize returns the given 128-bit, block cipher wrapped in Galois
// Counter Mode, which generates tags with the given length.
//
// Tag sizes between 12 and 16 bytes are allowed.
//
// Only use this function if you require compatibility with an existing
// cryptosystem that uses non-standard tag lengths. All other users should use
// [NewGCM], which is more resistant to misuse.
func NewGCMWithTagSize(t testing.TB, block cipher.Block, tagSize int) cipher.AEAD {
	t.Helper()
	p0, err := cipher.NewGCMWithTagSize(block, tagSize)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
