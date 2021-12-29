package rpctest

import (
	"net/rpc"
	"testing"
)

// DialHTTP connects to an HTTP RPC server at the specified network address
// listening on the default HTTP RPC path.
func DialHTTP(t testing.TB, network, address string) *rpc.Client {
	t.Helper()
	p0, err := rpc.DialHTTP(network, address)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// DialHTTPPath connects to an HTTP RPC server
// at the specified network address and path.
func DialHTTPPath(t testing.TB, network, address, path string) *rpc.Client {
	t.Helper()
	p0, err := rpc.DialHTTPPath(network, address, path)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// Dial connects to an RPC server at the specified network address.
func Dial(t testing.TB, network, address string) *rpc.Client {
	t.Helper()
	p0, err := rpc.Dial(network, address)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
