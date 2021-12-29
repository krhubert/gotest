package urltest

import (
	"net/url"
	"testing"
)

// QueryUnescape does the inverse transformation of [QueryEscape],
// converting each 3-byte encoded substring of the form "%AB" into the
// hex-decoded byte 0xAB.
// It returns an error if any % is not followed by two hexadecimal
// digits.
func QueryUnescape(t testing.TB, s string) string {
	t.Helper()
	p0, err := url.QueryUnescape(s)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// PathUnescape does the inverse transformation of [PathEscape],
// converting each 3-byte encoded substring of the form "%AB" into the
// hex-decoded byte 0xAB. It returns an error if any % is not followed
// by two hexadecimal digits.
//
// PathUnescape is identical to [QueryUnescape] except that it does not
// unescape '+' to ' ' (space).
func PathUnescape(t testing.TB, s string) string {
	t.Helper()
	p0, err := url.PathUnescape(s)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// Parse parses a raw url into a [URL] structure.
//
// The url may be relative (a path, without a host) or absolute
// (starting with a scheme). Trying to parse a hostname and path
// without a scheme is invalid but may not necessarily return an
// error, due to parsing ambiguities.
func Parse(t testing.TB, rawURL string) *url.URL {
	t.Helper()
	p0, err := url.Parse(rawURL)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// ParseRequestURI parses a raw url into a [URL] structure. It assumes that
// url was received in an HTTP request, so the url is interpreted
// only as an absolute URI or an absolute path.
// The string url is assumed not to have a #fragment suffix.
// (Web browsers strip #fragment before sending the URL to a web server.)
func ParseRequestURI(t testing.TB, rawURL string) *url.URL {
	t.Helper()
	p0, err := url.ParseRequestURI(rawURL)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// ParseQuery parses the URL-encoded query string and returns
// a map listing the values specified for each key.
// ParseQuery always returns a non-nil map containing all the
// valid query parameters found; err describes the first decoding error
// encountered, if any.
//
// Query is expected to be a list of key=value settings separated by ampersands.
// A setting without an equals sign is interpreted as a key set to an empty
// value.
// Settings containing a non-URL-encoded semicolon are considered invalid.
func ParseQuery(t testing.TB, query string) url.Values {
	t.Helper()
	p0, err := url.ParseQuery(query)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// JoinPath returns a [URL] string with the provided path elements joined to
// the existing path of base and the resulting path cleaned of any ./ or ../ elements.
func JoinPath(t testing.TB, base string, elem ...string) (result string) {
	t.Helper()
	p0, err := url.JoinPath(base, elem...)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
