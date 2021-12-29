package rsatest

import (
	"crypto/rsa"
	"hash"
	"io"
	"testing"
)

// GenerateKey generates a random RSA private key of the given bit size.
//
// Most applications should use [crypto/rand.Reader] as rand. Note that the
// returned key does not depend deterministically on the bytes read from rand,
// and may change between calls and/or between versions.
func GenerateKey(t testing.TB, random io.Reader, bits int) *rsa.PrivateKey {
	t.Helper()
	p0, err := rsa.GenerateKey(random, bits)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// GenerateMultiPrimeKey generates a multi-prime RSA keypair of the given bit
// size and the given random source.
//
// Table 1 in "[On the Security of Multi-prime RSA]" suggests maximum numbers of
// primes for a given bit size.
//
// Although the public keys are compatible (actually, indistinguishable) from
// the 2-prime case, the private keys are not. Thus it may not be possible to
// export multi-prime private keys in certain formats or to subsequently import
// them into other code.
//
// This package does not implement CRT optimizations for multi-prime RSA, so the
// keys with more than two primes will have worse performance.
//
// Deprecated: The use of this function with a number of primes different from
// two is not recommended for the above security, compatibility, and performance
// reasons. Use [GenerateKey] instead.
//
// [On the Security of Multi-prime RSA]: http://www.cacr.math.uwaterloo.ca/techreports/2006/cacr2006-16.pdf
func GenerateMultiPrimeKey(t testing.TB, random io.Reader, nprimes int, bits int) *rsa.PrivateKey {
	t.Helper()
	p0, err := rsa.GenerateMultiPrimeKey(random, nprimes, bits)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// EncryptOAEP encrypts the given message with RSA-OAEP.
//
// OAEP is parameterised by a hash function that is used as a random oracle.
// Encryption and decryption of a given message must use the same hash function
// and sha256.New() is a reasonable choice.
//
// The random parameter is used as a source of entropy to ensure that
// encrypting the same message twice doesn't result in the same ciphertext.
// Most applications should use [crypto/rand.Reader] as random.
//
// The label parameter may contain arbitrary data that will not be encrypted,
// but which gives important context to the message. For example, if a given
// public key is used to encrypt two types of messages then distinct label
// values could be used to ensure that a ciphertext for one purpose cannot be
// used for another by an attacker. If not required it can be empty.
//
// The message must be no longer than the length of the public modulus minus
// twice the hash length, minus a further 2.
func EncryptOAEP(t testing.TB, hash hash.Hash, random io.Reader, pub *rsa.PublicKey, msg []byte, label []byte) []byte {
	t.Helper()
	p0, err := rsa.EncryptOAEP(hash, random, pub, msg, label)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// DecryptOAEP decrypts ciphertext using RSA-OAEP.
//
// OAEP is parameterised by a hash function that is used as a random oracle.
// Encryption and decryption of a given message must use the same hash function
// and sha256.New() is a reasonable choice.
//
// The random parameter is legacy and ignored, and it can be nil.
//
// The label parameter must match the value given when encrypting. See
// [EncryptOAEP] for details.
func DecryptOAEP(t testing.TB, hash hash.Hash, random io.Reader, priv *rsa.PrivateKey, ciphertext []byte, label []byte) []byte {
	t.Helper()
	p0, err := rsa.DecryptOAEP(hash, random, priv, ciphertext, label)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
