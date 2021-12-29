package nettest

import (
	"net"
	"testing"
)

// SplitHostPort splits a network address of the form "host:port",
// "host%zone:port", "[host]:port" or "[host%zone]:port" into host or
// host%zone and port.
//
// A literal IPv6 address in hostport must be enclosed in square
// brackets, as in "[::1]:80", "[::1%lo0]:80".
//
// See func Dial for a description of the hostport parameter, and host
// and port results.
func SplitHostPort(t testing.TB, hostport string) (host, port string) {
	t.Helper()
	p0, p1, err := net.SplitHostPort(hostport)
	if err != nil {
		t.Fatal(err)
	}
	return p0, p1
}
