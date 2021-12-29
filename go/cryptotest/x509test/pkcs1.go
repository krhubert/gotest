package x509test

import (
	"crypto/rsa"
	"crypto/x509"
	"testing"
)

// ParsePKCS1PrivateKey parses an [RSA] private key in PKCS #1, ASN.1 DER form.
//
// This kind of key is commonly encoded in PEM blocks of type "RSA PRIVATE KEY".
func ParsePKCS1PrivateKey(t testing.TB, der []byte) *rsa.PrivateKey {
	t.Helper()
	p0, err := x509.ParsePKCS1PrivateKey(der)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// ParsePKCS1PublicKey parses an [RSA] public key in PKCS #1, ASN.1 DER form.
//
// This kind of key is commonly encoded in PEM blocks of type "RSA PUBLIC KEY".
func ParsePKCS1PublicKey(t testing.TB, der []byte) *rsa.PublicKey {
	t.Helper()
	p0, err := x509.ParsePKCS1PublicKey(der)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
