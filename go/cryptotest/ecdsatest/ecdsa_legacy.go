package ecdsatest

import (
	"crypto/ecdsa"
	"io"
	"math/big"
	"testing"
)

// Sign signs a hash (which should be the result of hashing a larger message)
// using the private key, priv. If the hash is longer than the bit-length of the
// private key's curve order, the hash will be truncated to that length. It
// returns the signature as a pair of integers. Most applications should use
// [SignASN1] instead of dealing directly with r, s.
func Sign(t testing.TB, rand io.Reader, priv *ecdsa.PrivateKey, hash []byte) (r, s *big.Int) {
	t.Helper()
	p0, p1, err := ecdsa.Sign(rand, priv, hash)
	if err != nil {
		t.Fatal(err)
	}
	return p0, p1
}
