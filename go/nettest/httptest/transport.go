package httptest

import (
	"net/http"
	"net/url"
	"testing"
)

// ProxyFromEnvironment returns the URL of the proxy to use for a
// given request, as indicated by the environment variables
// HTTP_PROXY, HTTPS_PROXY and NO_PROXY (or the lowercase versions
// thereof). Requests use the proxy from the environment variable
// matching their scheme, unless excluded by NO_PROXY.
//
// The environment values may be either a complete URL or a
// "host[:port]", in which case the "http" scheme is assumed.
// An error is returned if the value is a different form.
//
// A nil URL and nil error are returned if no proxy is defined in the
// environment, or a proxy should not be used for the given request,
// as defined by NO_PROXY.
//
// As a special case, if req.URL.Host is "localhost" (with or without
// a port number), then a nil URL and nil error will be returned.
func ProxyFromEnvironment(t testing.TB, req *http.Request) *url.URL {
	t.Helper()
	p0, err := http.ProxyFromEnvironment(req)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
