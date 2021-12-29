package asn1test

import (
	"encoding/asn1"
	"testing"
)

// Marshal returns the ASN.1 encoding of val.
//
// In addition to the struct tags recognized by Unmarshal, the following can be
// used:
//
//	ia5:         causes strings to be marshaled as ASN.1, IA5String values
//	omitempty:   causes empty slices to be skipped
//	printable:   causes strings to be marshaled as ASN.1, PrintableString values
//	utf8:        causes strings to be marshaled as ASN.1, UTF8String values
//	utc:         causes time.Time to be marshaled as ASN.1, UTCTime values
//	generalized: causes time.Time to be marshaled as ASN.1, GeneralizedTime values
func Marshal(t testing.TB, val any) []byte {
	t.Helper()
	p0, err := asn1.Marshal(val)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// MarshalWithParams allows field parameters to be specified for the
// top-level element. The form of the params is the same as the field tags.
func MarshalWithParams(t testing.TB, val any, params string) []byte {
	t.Helper()
	p0, err := asn1.MarshalWithParams(val, params)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
