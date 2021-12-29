package rpctest

import (
	"net/rpc"
	"testing"
)

// Register publishes the receiver's methods in the [DefaultServer].
func Register(t testing.TB, rcvr any) {
	t.Helper()
	err := rpc.Register(rcvr)
	if err != nil {
		t.Fatal(err)
	}
	return
}

// RegisterName is like [Register] but uses the provided name for the type
// instead of the receiver's concrete type.
func RegisterName(t testing.TB, name string, rcvr any) {
	t.Helper()
	err := rpc.RegisterName(name, rcvr)
	if err != nil {
		t.Fatal(err)
	}
	return
}

// ServeRequest is like [ServeCodec] but synchronously serves a single request.
// It does not close the codec upon completion.
func ServeRequest(t testing.TB, codec rpc.ServerCodec) {
	t.Helper()
	err := rpc.ServeRequest(codec)
	if err != nil {
		t.Fatal(err)
	}
	return
}
