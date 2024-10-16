package asn1test

import (
	"encoding/asn1"
	"testing"
)

// Unmarshal parses the DER-encoded ASN.1 data structure b
// and uses the reflect package to fill in an arbitrary value pointed at by val.
// Because Unmarshal uses the reflect package, the structs
// being written to must use upper case field names. If val
// is nil or not a pointer, Unmarshal returns an error.
//
// After parsing b, any bytes that were leftover and not used to fill
// val will be returned in rest. When parsing a SEQUENCE into a struct,
// any trailing elements of the SEQUENCE that do not have matching
// fields in val will not be included in rest, as these are considered
// valid elements of the SEQUENCE and not trailing data.
//
//   - An ASN.1 INTEGER can be written to an int, int32, int64,
//     or *[big.Int].
//     If the encoded value does not fit in the Go type,
//     Unmarshal returns a parse error.
//
//   - An ASN.1 BIT STRING can be written to a [BitString].
//
//   - An ASN.1 OCTET STRING can be written to a []byte.
//
//   - An ASN.1 OBJECT IDENTIFIER can be written to an [ObjectIdentifier].
//
//   - An ASN.1 ENUMERATED can be written to an [Enumerated].
//
//   - An ASN.1 UTCTIME or GENERALIZEDTIME can be written to a [time.Time].
//
//   - An ASN.1 PrintableString, IA5String, or NumericString can be written to a string.
//
//   - Any of the above ASN.1 values can be written to an interface{}.
//     The value stored in the interface has the corresponding Go type.
//     For integers, that type is int64.
//
//   - An ASN.1 SEQUENCE OF x or SET OF x can be written
//     to a slice if an x can be written to the slice's element type.
//
//   - An ASN.1 SEQUENCE or SET can be written to a struct
//     if each of the elements in the sequence can be
//     written to the corresponding element in the struct.
//
// The following tags on struct fields have special meaning to Unmarshal:
//
//	application specifies that an APPLICATION tag is used
//	private     specifies that a PRIVATE tag is used
//	default:x   sets the default value for optional integer fields (only used if optional is also present)
//	explicit    specifies that an additional, explicit tag wraps the implicit one
//	optional    marks the field as ASN.1 OPTIONAL
//	set         causes a SET, rather than a SEQUENCE type to be expected
//	tag:x       specifies the ASN.1 tag number; implies ASN.1 CONTEXT SPECIFIC
//
// When decoding an ASN.1 value with an IMPLICIT tag into a string field,
// Unmarshal will default to a PrintableString, which doesn't support
// characters such as '@' and '&'. To force other encodings, use the following
// tags:
//
//	ia5     causes strings to be unmarshaled as ASN.1 IA5String values
//	numeric causes strings to be unmarshaled as ASN.1 NumericString values
//	utf8    causes strings to be unmarshaled as ASN.1 UTF8String values
//
// If the type of the first field of a structure is RawContent then the raw
// ASN1 contents of the struct will be stored in it.
//
// If the name of a slice type ends with "SET" then it's treated as if
// the "set" tag was set on it. This results in interpreting the type as a
// SET OF x rather than a SEQUENCE OF x. This can be used with nested slices
// where a struct tag cannot be given.
//
// Other ASN.1 types are not supported; if it encounters them,
// Unmarshal returns a parse error.
func Unmarshal(t testing.TB, b []byte, val any) (rest []byte) {
	t.Helper()
	p0, err := asn1.Unmarshal(b, val)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// UnmarshalWithParams allows field parameters to be specified for the
// top-level element. The form of the params is the same as the field tags.
func UnmarshalWithParams(t testing.TB, b []byte, val any, params string) (rest []byte) {
	t.Helper()
	p0, err := asn1.UnmarshalWithParams(b, val, params)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
