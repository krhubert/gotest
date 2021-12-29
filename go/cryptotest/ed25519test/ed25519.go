package ed25519test

import (
	"crypto/ed25519"
	"io"
	"testing"
)

// GenerateKey generates a public/private key pair using entropy from rand.
// If rand is nil, [crypto/rand.Reader] will be used.
//
// The output of this function is deterministic, and equivalent to reading
// [SeedSize] bytes from rand, and passing them to [NewKeyFromSeed].
func GenerateKey(t testing.TB, rand io.Reader) (ed25519.PublicKey, ed25519.PrivateKey) {
	t.Helper()
	p0, p1, err := ed25519.GenerateKey(rand)
	if err != nil {
		t.Fatal(err)
	}
	return p0, p1
}

// VerifyWithOptions reports whether sig is a valid signature of message by
// publicKey. A valid signature is indicated by returning a nil error. It will
// panic if len(publicKey) is not [PublicKeySize].
//
// If opts.Hash is [crypto.SHA512], the pre-hashed variant Ed25519ph is used and
// message is expected to be a SHA-512 hash, otherwise opts.Hash must be
// [crypto.Hash](0) and the message must not be hashed, as Ed25519 performs two
// passes over messages to be signed.
//
// The inputs are not considered confidential, and may leak through timing side
// channels, or if an attacker has control of part of the inputs.
func VerifyWithOptions(t testing.TB, publicKey ed25519.PublicKey, message, sig []byte, opts *ed25519.Options) {
	t.Helper()
	err := ed25519.VerifyWithOptions(publicKey, message, sig, opts)
	if err != nil {
		t.Fatal(err)
	}
	return
}
