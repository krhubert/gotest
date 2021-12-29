package randtest

import (
	"crypto/rand"
	"io"
	"math/big"
	"testing"
)

// Prime returns a number of the given bit length that is prime with high probability.
// Prime will return error for any error returned by [rand.Read] or if bits < 2.
func Prime(t testing.TB, r io.Reader, bits int) *big.Int {
	t.Helper()
	p0, err := rand.Prime(r, bits)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// Int returns a uniform random value in [0, max). It panics if max <= 0.
func Int(t testing.TB, r io.Reader, max *big.Int) (n *big.Int) {
	t.Helper()
	p0, err := rand.Int(r, max)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
