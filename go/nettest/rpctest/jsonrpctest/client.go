package jsonrpctest

import (
	"net/rpc"
	"net/rpc/jsonrpc"
	"testing"
)

// Dial connects to a JSON-RPC server at the specified network address.
func Dial(t testing.TB, network, address string) *rpc.Client {
	t.Helper()
	p0, err := jsonrpc.Dial(network, address)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
