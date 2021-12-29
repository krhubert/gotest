package x509test

import (
	"crypto/x509"
	"testing"
)

// ParseOID parses a Object Identifier string, represented by ASCII numbers separated by dots.
func ParseOID(t testing.TB, oid string) x509.OID {
	t.Helper()
	p0, err := x509.ParseOID(oid)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// OIDFromInts creates a new OID using ints, each integer is a separate component.
func OIDFromInts(t testing.TB, oid []uint64) x509.OID {
	t.Helper()
	p0, err := x509.OIDFromInts(oid)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
