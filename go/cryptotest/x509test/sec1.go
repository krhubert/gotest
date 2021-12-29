package x509test

import (
	"crypto/ecdsa"
	"crypto/x509"
	"testing"
)

// ParseECPrivateKey parses an EC private key in SEC 1, ASN.1 DER form.
//
// This kind of key is commonly encoded in PEM blocks of type "EC PRIVATE KEY".
func ParseECPrivateKey(t testing.TB, der []byte) *ecdsa.PrivateKey {
	t.Helper()
	p0, err := x509.ParseECPrivateKey(der)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// MarshalECPrivateKey converts an EC private key to SEC 1, ASN.1 DER form.
//
// This kind of key is commonly encoded in PEM blocks of type "EC PRIVATE KEY".
// For a more flexible key format which is not EC specific, use
// [MarshalPKCS8PrivateKey].
func MarshalECPrivateKey(t testing.TB, key *ecdsa.PrivateKey) []byte {
	t.Helper()
	p0, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
