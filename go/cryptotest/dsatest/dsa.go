package dsatest

import (
	"crypto/dsa"
	"io"
	"math/big"
	"testing"
)

// GenerateParameters puts a random, valid set of DSA parameters into params.
// This function can take many seconds, even on fast machines.
func GenerateParameters(t testing.TB, params *dsa.Parameters, rand io.Reader, sizes dsa.ParameterSizes) {
	t.Helper()
	err := dsa.GenerateParameters(params, rand, sizes)
	if err != nil {
		t.Fatal(err)
	}
	return
}

// GenerateKey generates a public&private key pair. The Parameters of the
// [PrivateKey] must already be valid (see [GenerateParameters]).
func GenerateKey(t testing.TB, priv *dsa.PrivateKey, rand io.Reader) {
	t.Helper()
	err := dsa.GenerateKey(priv, rand)
	if err != nil {
		t.Fatal(err)
	}
	return
}

// Sign signs an arbitrary length hash (which should be the result of hashing a
// larger message) using the private key, priv. It returns the signature as a
// pair of integers. The security of the private key depends on the entropy of
// rand.
//
// Note that FIPS 186-3 section 4.6 specifies that the hash should be truncated
// to the byte-length of the subgroup. This function does not perform that
// truncation itself.
//
// Be aware that calling Sign with an attacker-controlled [PrivateKey] may
// require an arbitrary amount of CPU.
func Sign(t testing.TB, rand io.Reader, priv *dsa.PrivateKey, hash []byte) (r, s *big.Int) {
	t.Helper()
	p0, p1, err := dsa.Sign(rand, priv, hash)
	if err != nil {
		t.Fatal(err)
	}
	return p0, p1
}
