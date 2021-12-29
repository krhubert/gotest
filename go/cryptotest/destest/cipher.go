package destest

import (
	"crypto/cipher"
	"crypto/des"
	"testing"
)

// NewCipher creates and returns a new [cipher.Block].
func NewCipher(t testing.TB, key []byte) cipher.Block {
	t.Helper()
	p0, err := des.NewCipher(key)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// NewTripleDESCipher creates and returns a new [cipher.Block].
func NewTripleDESCipher(t testing.TB, key []byte) cipher.Block {
	t.Helper()
	p0, err := des.NewTripleDESCipher(key)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
