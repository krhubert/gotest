package httptest

import (
	"net/http"
	"testing"
)

// ParseCookie parses a Cookie header value and returns all the cookies
// which were set in it. Since the same cookie name can appear multiple times
// the returned Values can contain more than one value for a given key.
func ParseCookie(t testing.TB, line string) []*http.Cookie {
	t.Helper()
	p0, err := http.ParseCookie(line)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// ParseSetCookie parses a Set-Cookie header value and returns a cookie.
// It returns an error on syntax error.
func ParseSetCookie(t testing.TB, line string) *http.Cookie {
	t.Helper()
	p0, err := http.ParseSetCookie(line)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
