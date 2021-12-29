package x509test

import (
	"crypto/x509"
	"testing"
)

// ParseCertificate parses a single certificate from the given ASN.1 DER data.
//
// Before Go 1.23, ParseCertificate accepted certificates with negative serial
// numbers. This behavior can be restored by including "x509negativeserial=1" in
// the GODEBUG environment variable.
func ParseCertificate(t testing.TB, der []byte) *x509.Certificate {
	t.Helper()
	p0, err := x509.ParseCertificate(der)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// ParseCertificates parses one or more certificates from the given ASN.1 DER
// data. The certificates must be concatenated with no intermediate padding.
func ParseCertificates(t testing.TB, der []byte) []*x509.Certificate {
	t.Helper()
	p0, err := x509.ParseCertificates(der)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// ParseRevocationList parses a X509 v2 [Certificate] Revocation List from the given
// ASN.1 DER data.
func ParseRevocationList(t testing.TB, der []byte) *x509.RevocationList {
	t.Helper()
	p0, err := x509.ParseRevocationList(der)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
