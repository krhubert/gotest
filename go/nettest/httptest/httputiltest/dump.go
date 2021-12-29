package httputiltest

import (
	"net/http"
	"net/http/httputil"
	"testing"
)

// DumpRequestOut is like [DumpRequest] but for outgoing client requests. It
// includes any headers that the standard [http.Transport] adds, such as
// User-Agent.
func DumpRequestOut(t testing.TB, req *http.Request, body bool) []byte {
	t.Helper()
	p0, err := httputil.DumpRequestOut(req, body)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// DumpRequest returns the given request in its HTTP/1.x wire
// representation. It should only be used by servers to debug client
// requests. The returned representation is an approximation only;
// some details of the initial request are lost while parsing it into
// an [http.Request]. In particular, the order and case of header field
// names are lost. The order of values in multi-valued headers is kept
// intact. HTTP/2 requests are dumped in HTTP/1.x form, not in their
// original binary representations.
//
// If body is true, DumpRequest also returns the body. To do so, it
// consumes req.Body and then replaces it with a new [io.ReadCloser]
// that yields the same bytes. If DumpRequest returns an error,
// the state of req is undefined.
//
// The documentation for [http.Request.Write] details which fields
// of req are included in the dump.
func DumpRequest(t testing.TB, req *http.Request, body bool) []byte {
	t.Helper()
	p0, err := httputil.DumpRequest(req, body)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// DumpResponse is like DumpRequest but dumps a response.
func DumpResponse(t testing.TB, resp *http.Response, body bool) []byte {
	t.Helper()
	p0, err := httputil.DumpResponse(resp, body)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
