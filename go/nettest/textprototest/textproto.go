package textprototest

import (
	"net/textproto"
	"testing"
)

// Dial connects to the given address on the given network using [net.Dial]
// and then returns a new [Conn] for the connection.
func Dial(t testing.TB, network, addr string) *textproto.Conn {
	t.Helper()
	p0, err := textproto.Dial(network, addr)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
