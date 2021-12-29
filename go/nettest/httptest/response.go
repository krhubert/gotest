package httptest

import (
	"bufio"
	"net/http"
	"testing"
)

// ReadResponse reads and returns an HTTP response from r.
// The req parameter optionally specifies the [Request] that corresponds
// to this [Response]. If nil, a GET request is assumed.
// Clients must call resp.Body.Close when finished reading resp.Body.
// After that call, clients can inspect resp.Trailer to find key/value
// pairs included in the response trailer.
func ReadResponse(t testing.TB, r *bufio.Reader, req *http.Request) *http.Response {
	t.Helper()
	p0, err := http.ReadResponse(r, req)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
