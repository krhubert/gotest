package mimetest

import (
	"mime"
	"testing"
)

// ParseMediaType parses a media type value and any optional
// parameters, per RFC 1521.  Media types are the values in
// Content-Type and Content-Disposition headers (RFC 2183).
// On success, ParseMediaType returns the media type converted
// to lowercase and trimmed of white space and a non-nil map.
// If there is an error parsing the optional parameter,
// the media type will be returned along with the error
// [ErrInvalidMediaParameter].
// The returned map, params, maps from the lowercase
// attribute to the attribute value with its case preserved.
func ParseMediaType(t testing.TB, v string) (mediatype string, params map[string]string) {
	t.Helper()
	p0, p1, err := mime.ParseMediaType(v)
	if err != nil {
		t.Fatal(err)
	}
	return p0, p1
}
