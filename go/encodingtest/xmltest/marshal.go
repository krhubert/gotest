package xmltest

import (
	"encoding/xml"
	"testing"
)

// Marshal returns the XML encoding of v.
//
// Marshal handles an array or slice by marshaling each of the elements.
// Marshal handles a pointer by marshaling the value it points at or, if the
// pointer is nil, by writing nothing. Marshal handles an interface value by
// marshaling the value it contains or, if the interface value is nil, by
// writing nothing. Marshal handles all other data by writing one or more XML
// elements containing the data.
//
// The name for the XML elements is taken from, in order of preference:
//   - the tag on the XMLName field, if the data is a struct
//   - the value of the XMLName field of type [Name]
//   - the tag of the struct field used to obtain the data
//   - the name of the struct field used to obtain the data
//   - the name of the marshaled type
//
// The XML element for a struct contains marshaled elements for each of the
// exported fields of the struct, with these exceptions:
//   - the XMLName field, described above, is omitted.
//   - a field with tag "-" is omitted.
//   - a field with tag "name,attr" becomes an attribute with
//     the given name in the XML element.
//   - a field with tag ",attr" becomes an attribute with the
//     field name in the XML element.
//   - a field with tag ",chardata" is written as character data,
//     not as an XML element.
//   - a field with tag ",cdata" is written as character data
//     wrapped in one or more <![CDATA[ ... ]]> tags, not as an XML element.
//   - a field with tag ",innerxml" is written verbatim, not subject
//     to the usual marshaling procedure.
//   - a field with tag ",comment" is written as an XML comment, not
//     subject to the usual marshaling procedure. It must not contain
//     the "--" string within it.
//   - a field with a tag including the "omitempty" option is omitted
//     if the field value is empty. The empty values are false, 0, any
//     nil pointer or interface value, and any array, slice, map, or
//     string of length zero.
//   - an anonymous struct field is handled as if the fields of its
//     value were part of the outer struct.
//   - a field implementing [Marshaler] is written by calling its MarshalXML
//     method.
//   - a field implementing [encoding.TextMarshaler] is written by encoding the
//     result of its MarshalText method as text.
//
// If a field uses a tag "a>b>c", then the element c will be nested inside
// parent elements a and b. Fields that appear next to each other that name
// the same parent will be enclosed in one XML element.
//
// If the XML name for a struct field is defined by both the field tag and the
// struct's XMLName field, the names must match.
//
// See [MarshalIndent] for an example.
//
// Marshal will return an error if asked to marshal a channel, function, or map.
func Marshal(t testing.TB, v any) []byte {
	t.Helper()
	p0, err := xml.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// MarshalIndent works like [Marshal], but each XML element begins on a new
// indented line that starts with prefix and is followed by one or more
// copies of indent according to the nesting depth.
func MarshalIndent(t testing.TB, v any, prefix, indent string) []byte {
	t.Helper()
	p0, err := xml.MarshalIndent(v, prefix, indent)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
