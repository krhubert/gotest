package fcgitest

import (
	"net"
	"net/http"
	"net/http/fcgi"
	"testing"
)

// Serve accepts incoming FastCGI connections on the listener l, creating a new
// goroutine for each. The goroutine reads requests and then calls handler
// to reply to them.
// If l is nil, Serve accepts connections from os.Stdin.
// If handler is nil, [http.DefaultServeMux] is used.
func Serve(t testing.TB, l net.Listener, handler http.Handler) {
	t.Helper()
	err := fcgi.Serve(l, handler)
	if err != nil {
		t.Fatal(err)
	}
	return
}
