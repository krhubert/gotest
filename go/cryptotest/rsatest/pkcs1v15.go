package rsatest

import (
	"crypto"
	"crypto/rsa"
	"io"
	"testing"
)

// EncryptPKCS1v15 encrypts the given message with RSA and the padding
// scheme from PKCS #1 v1.5.  The message must be no longer than the
// length of the public modulus minus 11 bytes.
//
// The random parameter is used as a source of entropy to ensure that
// encrypting the same message twice doesn't result in the same
// ciphertext. Most applications should use [crypto/rand.Reader]
// as random. Note that the returned ciphertext does not depend
// deterministically on the bytes read from random, and may change
// between calls and/or between versions.
//
// WARNING: use of this function to encrypt plaintexts other than
// session keys is dangerous. Use RSA OAEP in new protocols.
func EncryptPKCS1v15(t testing.TB, random io.Reader, pub *rsa.PublicKey, msg []byte) []byte {
	t.Helper()
	p0, err := rsa.EncryptPKCS1v15(random, pub, msg)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// DecryptPKCS1v15 decrypts a plaintext using RSA and the padding scheme from PKCS #1 v1.5.
// The random parameter is legacy and ignored, and it can be nil.
//
// Note that whether this function returns an error or not discloses secret
// information. If an attacker can cause this function to run repeatedly and
// learn whether each instance returned an error then they can decrypt and
// forge signatures as if they had the private key. See
// DecryptPKCS1v15SessionKey for a way of solving this problem.
func DecryptPKCS1v15(t testing.TB, random io.Reader, priv *rsa.PrivateKey, ciphertext []byte) []byte {
	t.Helper()
	p0, err := rsa.DecryptPKCS1v15(random, priv, ciphertext)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// DecryptPKCS1v15SessionKey decrypts a session key using RSA and the padding
// scheme from PKCS #1 v1.5. The random parameter is legacy and ignored, and it
// can be nil.
//
// DecryptPKCS1v15SessionKey returns an error if the ciphertext is the wrong
// length or if the ciphertext is greater than the public modulus. Otherwise, no
// error is returned. If the padding is valid, the resulting plaintext message
// is copied into key. Otherwise, key is unchanged. These alternatives occur in
// constant time. It is intended that the user of this function generate a
// random session key beforehand and continue the protocol with the resulting
// value.
//
// Note that if the session key is too small then it may be possible for an
// attacker to brute-force it. If they can do that then they can learn whether a
// random value was used (because it'll be different for the same ciphertext)
// and thus whether the padding was correct. This also defeats the point of this
// function. Using at least a 16-byte key will protect against this attack.
//
// This method implements protections against Bleichenbacher chosen ciphertext
// attacks [0] described in RFC 3218 Section 2.3.2 [1]. While these protections
// make a Bleichenbacher attack significantly more difficult, the protections
// are only effective if the rest of the protocol which uses
// DecryptPKCS1v15SessionKey is designed with these considerations in mind. In
// particular, if any subsequent operations which use the decrypted session key
// leak any information about the key (e.g. whether it is a static or random
// key) then the mitigations are defeated. This method must be used extremely
// carefully, and typically should only be used when absolutely necessary for
// compatibility with an existing protocol (such as TLS) that is designed with
// these properties in mind.
//
//   - [0] “Chosen Ciphertext Attacks Against Protocols Based on the RSA Encryption
//     Standard PKCS #1”, Daniel Bleichenbacher, Advances in Cryptology (Crypto '98)
//   - [1] RFC 3218, Preventing the Million Message Attack on CMS,
//     https://www.rfc-editor.org/rfc/rfc3218.html
func DecryptPKCS1v15SessionKey(t testing.TB, random io.Reader, priv *rsa.PrivateKey, ciphertext []byte, key []byte) {
	t.Helper()
	err := rsa.DecryptPKCS1v15SessionKey(random, priv, ciphertext, key)
	if err != nil {
		t.Fatal(err)
	}
	return
}

// SignPKCS1v15 calculates the signature of hashed using
// RSASSA-PKCS1-V1_5-SIGN from RSA PKCS #1 v1.5.  Note that hashed must
// be the result of hashing the input message using the given hash
// function. If hash is zero, hashed is signed directly. This isn't
// advisable except for interoperability.
//
// The random parameter is legacy and ignored, and it can be nil.
//
// This function is deterministic. Thus, if the set of possible
// messages is small, an attacker may be able to build a map from
// messages to signatures and identify the signed messages. As ever,
// signatures provide authenticity, not confidentiality.
func SignPKCS1v15(t testing.TB, random io.Reader, priv *rsa.PrivateKey, hash crypto.Hash, hashed []byte) []byte {
	t.Helper()
	p0, err := rsa.SignPKCS1v15(random, priv, hash, hashed)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// VerifyPKCS1v15 verifies an RSA PKCS #1 v1.5 signature.
// hashed is the result of hashing the input message using the given hash
// function and sig is the signature. A valid signature is indicated by
// returning a nil error. If hash is zero then hashed is used directly. This
// isn't advisable except for interoperability.
//
// The inputs are not considered confidential, and may leak through timing side
// channels, or if an attacker has control of part of the inputs.
func VerifyPKCS1v15(t testing.TB, pub *rsa.PublicKey, hash crypto.Hash, hashed []byte, sig []byte) {
	t.Helper()
	err := rsa.VerifyPKCS1v15(pub, hash, hashed, sig)
	if err != nil {
		t.Fatal(err)
	}
	return
}
