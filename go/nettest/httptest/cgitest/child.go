package cgitest

import (
	"net/http"
	"net/http/cgi"
	"testing"
)

// Request returns the HTTP request as represented in the current
// environment. This assumes the current program is being run
// by a web server in a CGI environment.
// The returned Request's Body is populated, if applicable.
func Request(t testing.TB) *http.Request {
	t.Helper()
	p0, err := cgi.Request()
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// RequestFromMap creates an [http.Request] from CGI variables.
// The returned Request's Body field is not populated.
func RequestFromMap(t testing.TB, params map[string]string) *http.Request {
	t.Helper()
	p0, err := cgi.RequestFromMap(params)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// Serve executes the provided [Handler] on the currently active CGI
// request, if any. If there's no current CGI environment
// an error is returned. The provided handler may be nil to use
// [http.DefaultServeMux].
func Serve(t testing.TB, handler http.Handler) {
	t.Helper()
	err := cgi.Serve(handler)
	if err != nil {
		t.Fatal(err)
	}
	return
}
