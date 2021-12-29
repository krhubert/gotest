package httptest

import (
	"net"
	"net/http"
	"testing"
)

// Serve accepts incoming HTTP connections on the listener l,
// creating a new service goroutine for each. The service goroutines
// read requests and then call handler to reply to them.
//
// The handler is typically nil, in which case [DefaultServeMux] is used.
//
// HTTP/2 support is only enabled if the Listener returns [*tls.Conn]
// connections and they were configured with "h2" in the TLS
// Config.NextProtos.
//
// Serve always returns a non-nil error.
func Serve(t testing.TB, l net.Listener, handler http.Handler) {
	t.Helper()
	err := http.Serve(l, handler)
	if err != nil {
		t.Fatal(err)
	}
	return
}

// ServeTLS accepts incoming HTTPS connections on the listener l,
// creating a new service goroutine for each. The service goroutines
// read requests and then call handler to reply to them.
//
// The handler is typically nil, in which case [DefaultServeMux] is used.
//
// Additionally, files containing a certificate and matching private key
// for the server must be provided. If the certificate is signed by a
// certificate authority, the certFile should be the concatenation
// of the server's certificate, any intermediates, and the CA's certificate.
//
// ServeTLS always returns a non-nil error.
func ServeTLS(t testing.TB, l net.Listener, handler http.Handler, certFile, keyFile string) {
	t.Helper()
	err := http.ServeTLS(l, handler, certFile, keyFile)
	if err != nil {
		t.Fatal(err)
	}
	return
}

// ListenAndServe listens on the TCP network address addr and then calls
// [Serve] with handler to handle requests on incoming connections.
// Accepted connections are configured to enable TCP keep-alives.
//
// The handler is typically nil, in which case [DefaultServeMux] is used.
//
// ListenAndServe always returns a non-nil error.
func ListenAndServe(t testing.TB, addr string, handler http.Handler) {
	t.Helper()
	err := http.ListenAndServe(addr, handler)
	if err != nil {
		t.Fatal(err)
	}
	return
}

// ListenAndServeTLS acts identically to [ListenAndServe], except that it
// expects HTTPS connections. Additionally, files containing a certificate and
// matching private key for the server must be provided. If the certificate
// is signed by a certificate authority, the certFile should be the concatenation
// of the server's certificate, any intermediates, and the CA's certificate.
func ListenAndServeTLS(t testing.TB, addr, certFile, keyFile string, handler http.Handler) {
	t.Helper()
	err := http.ListenAndServeTLS(addr, certFile, keyFile, handler)
	if err != nil {
		t.Fatal(err)
	}
	return
}
