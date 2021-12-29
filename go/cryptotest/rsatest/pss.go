package rsatest

import (
	"crypto"
	"crypto/rsa"
	"io"
	"testing"
)

// SignPSS calculates the signature of digest using PSS.
//
// digest must be the result of hashing the input message using the given hash
// function. The opts argument may be nil, in which case sensible defaults are
// used. If opts.Hash is set, it overrides hash.
//
// The signature is randomized depending on the message, key, and salt size,
// using bytes from rand. Most applications should use [crypto/rand.Reader] as
// rand.
func SignPSS(t testing.TB, rand io.Reader, priv *rsa.PrivateKey, hash crypto.Hash, digest []byte, opts *rsa.PSSOptions) []byte {
	t.Helper()
	p0, err := rsa.SignPSS(rand, priv, hash, digest, opts)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// VerifyPSS verifies a PSS signature.
//
// A valid signature is indicated by returning a nil error. digest must be the
// result of hashing the input message using the given hash function. The opts
// argument may be nil, in which case sensible defaults are used. opts.Hash is
// ignored.
//
// The inputs are not considered confidential, and may leak through timing side
// channels, or if an attacker has control of part of the inputs.
func VerifyPSS(t testing.TB, pub *rsa.PublicKey, hash crypto.Hash, digest []byte, sig []byte, opts *rsa.PSSOptions) {
	t.Helper()
	err := rsa.VerifyPSS(pub, hash, digest, sig, opts)
	if err != nil {
		t.Fatal(err)
	}
	return
}
