package httptest

import (
	"bufio"
	"context"
	"io"
	"net/http"
	"testing"
)

// NewRequest wraps [NewRequestWithContext] using [context.Background].
func NewRequest(t testing.TB, method, url string, body io.Reader) *http.Request {
	t.Helper()
	p0, err := http.NewRequest(method, url, body)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// NewRequestWithContext returns a new [Request] given a method, URL, and
// optional body.
//
// If the provided body is also an [io.Closer], the returned
// [Request.Body] is set to body and will be closed (possibly
// asynchronously) by the Client methods Do, Post, and PostForm,
// and [Transport.RoundTrip].
//
// NewRequestWithContext returns a Request suitable for use with
// [Client.Do] or [Transport.RoundTrip]. To create a request for use with
// testing a Server Handler, either use the [NewRequest] function in the
// net/http/httptest package, use [ReadRequest], or manually update the
// Request fields. For an outgoing client request, the context
// controls the entire lifetime of a request and its response:
// obtaining a connection, sending the request, and reading the
// response headers and body. See the Request type's documentation for
// the difference between inbound and outbound request fields.
//
// If body is of type [*bytes.Buffer], [*bytes.Reader], or
// [*strings.Reader], the returned request's ContentLength is set to its
// exact value (instead of -1), GetBody is populated (so 307 and 308
// redirects can replay the body), and Body is set to [NoBody] if the
// ContentLength is 0.
func NewRequestWithContext(t testing.TB, ctx context.Context, method, url string, body io.Reader) *http.Request {
	t.Helper()
	p0, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// ReadRequest reads and parses an incoming request from b.
//
// ReadRequest is a low-level function and should only be used for
// specialized applications; most code should use the [Server] to read
// requests and handle them via the [Handler] interface. ReadRequest
// only supports HTTP/1.x requests. For HTTP/2, use golang.org/x/net/http2.
func ReadRequest(t testing.TB, b *bufio.Reader) *http.Request {
	t.Helper()
	p0, err := http.ReadRequest(b)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
