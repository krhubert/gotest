package ecdsatest

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"io"
	"testing"
)

// GenerateKey generates a new ECDSA private key for the specified curve.
//
// Most applications should use [crypto/rand.Reader] as rand. Note that the
// returned key does not depend deterministically on the bytes read from rand,
// and may change between calls and/or between versions.
func GenerateKey(t testing.TB, c elliptic.Curve, rand io.Reader) *ecdsa.PrivateKey {
	t.Helper()
	p0, err := ecdsa.GenerateKey(c, rand)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// SignASN1 signs a hash (which should be the result of hashing a larger message)
// using the private key, priv. If the hash is longer than the bit-length of the
// private key's curve order, the hash will be truncated to that length. It
// returns the ASN.1 encoded signature.
//
// The signature is randomized. Most applications should use [crypto/rand.Reader]
// as rand. Note that the returned signature does not depend deterministically on
// the bytes read from rand, and may change between calls and/or between versions.
func SignASN1(t testing.TB, rand io.Reader, priv *ecdsa.PrivateKey, hash []byte) []byte {
	t.Helper()
	p0, err := ecdsa.SignASN1(rand, priv, hash)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
