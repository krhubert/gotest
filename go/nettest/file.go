package nettest

import (
	"net"
	"os"
	"testing"
)

// FileConn returns a copy of the network connection corresponding to
// the open file f.
// It is the caller's responsibility to close f when finished.
// Closing c does not affect f, and closing f does not affect c.
func FileConn(t testing.TB, f *os.File) (c net.Conn) {
	t.Helper()
	p0, err := net.FileConn(f)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// FileListener returns a copy of the network listener corresponding
// to the open file f.
// It is the caller's responsibility to close ln when finished.
// Closing ln does not affect f, and closing f does not affect ln.
func FileListener(t testing.TB, f *os.File) (ln net.Listener) {
	t.Helper()
	p0, err := net.FileListener(f)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// FilePacketConn returns a copy of the packet network connection
// corresponding to the open file f.
// It is the caller's responsibility to close f when finished.
// Closing c does not affect f, and closing f does not affect c.
func FilePacketConn(t testing.TB, f *os.File) (c net.PacketConn) {
	t.Helper()
	p0, err := net.FilePacketConn(f)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
