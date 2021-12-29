package xmltest

import (
	"encoding/xml"
	"testing"
)

// Unmarshal parses the XML-encoded data and stores the result in
// the value pointed to by v, which must be an arbitrary struct,
// slice, or string. Well-formed data that does not fit into v is
// discarded.
//
// Because Unmarshal uses the reflect package, it can only assign
// to exported (upper case) fields. Unmarshal uses a case-sensitive
// comparison to match XML element names to tag values and struct
// field names.
//
// Unmarshal maps an XML element to a struct using the following rules.
// In the rules, the tag of a field refers to the value associated with the
// key 'xml' in the struct field's tag (see the example above).
//
//   - If the struct has a field of type []byte or string with tag
//     ",innerxml", Unmarshal accumulates the raw XML nested inside the
//     element in that field. The rest of the rules still apply.
//
//   - If the struct has a field named XMLName of type Name,
//     Unmarshal records the element name in that field.
//
//   - If the XMLName field has an associated tag of the form
//     "name" or "namespace-URL name", the XML element must have
//     the given name (and, optionally, name space) or else Unmarshal
//     returns an error.
//
//   - If the XML element has an attribute whose name matches a
//     struct field name with an associated tag containing ",attr" or
//     the explicit name in a struct field tag of the form "name,attr",
//     Unmarshal records the attribute value in that field.
//
//   - If the XML element has an attribute not handled by the previous
//     rule and the struct has a field with an associated tag containing
//     ",any,attr", Unmarshal records the attribute value in the first
//     such field.
//
//   - If the XML element contains character data, that data is
//     accumulated in the first struct field that has tag ",chardata".
//     The struct field may have type []byte or string.
//     If there is no such field, the character data is discarded.
//
//   - If the XML element contains comments, they are accumulated in
//     the first struct field that has tag ",comment".  The struct
//     field may have type []byte or string. If there is no such
//     field, the comments are discarded.
//
//   - If the XML element contains a sub-element whose name matches
//     the prefix of a tag formatted as "a" or "a>b>c", unmarshal
//     will descend into the XML structure looking for elements with the
//     given names, and will map the innermost elements to that struct
//     field. A tag starting with ">" is equivalent to one starting
//     with the field name followed by ">".
//
//   - If the XML element contains a sub-element whose name matches
//     a struct field's XMLName tag and the struct field has no
//     explicit name tag as per the previous rule, unmarshal maps
//     the sub-element to that struct field.
//
//   - If the XML element contains a sub-element whose name matches a
//     field without any mode flags (",attr", ",chardata", etc), Unmarshal
//     maps the sub-element to that struct field.
//
//   - If the XML element contains a sub-element that hasn't matched any
//     of the above rules and the struct has a field with tag ",any",
//     unmarshal maps the sub-element to that struct field.
//
//   - An anonymous struct field is handled as if the fields of its
//     value were part of the outer struct.
//
//   - A struct field with tag "-" is never unmarshaled into.
//
// If Unmarshal encounters a field type that implements the Unmarshaler
// interface, Unmarshal calls its UnmarshalXML method to produce the value from
// the XML element.  Otherwise, if the value implements
// [encoding.TextUnmarshaler], Unmarshal calls that value's UnmarshalText method.
//
// Unmarshal maps an XML element to a string or []byte by saving the
// concatenation of that element's character data in the string or
// []byte. The saved []byte is never nil.
//
// Unmarshal maps an attribute value to a string or []byte by saving
// the value in the string or slice.
//
// Unmarshal maps an attribute value to an [Attr] by saving the attribute,
// including its name, in the Attr.
//
// Unmarshal maps an XML element or attribute value to a slice by
// extending the length of the slice and mapping the element or attribute
// to the newly created value.
//
// Unmarshal maps an XML element or attribute value to a bool by
// setting it to the boolean value represented by the string. Whitespace
// is trimmed and ignored.
//
// Unmarshal maps an XML element or attribute value to an integer or
// floating-point field by setting the field to the result of
// interpreting the string value in decimal. There is no check for
// overflow. Whitespace is trimmed and ignored.
//
// Unmarshal maps an XML element to a Name by recording the element
// name.
//
// Unmarshal maps an XML element to a pointer by setting the pointer
// to a freshly allocated value and then mapping the element to that value.
//
// A missing element or empty attribute value will be unmarshaled as a zero value.
// If the field is a slice, a zero value will be appended to the field. Otherwise, the
// field will be set to its zero value.
func Unmarshal(t testing.TB, data []byte, v any) {
	t.Helper()
	err := xml.Unmarshal(data, v)
	if err != nil {
		t.Fatal(err)
	}
	return
}
