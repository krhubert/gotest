package nettest

import (
	"net"
	"testing"
)

// ResolveUnixAddr returns an address of Unix domain socket end point.
//
// The network must be a Unix network name.
//
// See func [Dial] for a description of the network and address
// parameters.
func ResolveUnixAddr(t testing.TB, network, address string) *net.UnixAddr {
	t.Helper()
	p0, err := net.ResolveUnixAddr(network, address)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// DialUnix acts like [Dial] for Unix networks.
//
// The network must be a Unix network name; see func Dial for details.
//
// If laddr is non-nil, it is used as the local address for the
// connection.
func DialUnix(t testing.TB, network string, laddr, raddr *net.UnixAddr) *net.UnixConn {
	t.Helper()
	p0, err := net.DialUnix(network, laddr, raddr)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// ListenUnix acts like [Listen] for Unix networks.
//
// The network must be "unix" or "unixpacket".
func ListenUnix(t testing.TB, network string, laddr *net.UnixAddr) *net.UnixListener {
	t.Helper()
	p0, err := net.ListenUnix(network, laddr)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// ListenUnixgram acts like [ListenPacket] for Unix networks.
//
// The network must be "unixgram".
func ListenUnixgram(t testing.TB, network string, laddr *net.UnixAddr) *net.UnixConn {
	t.Helper()
	p0, err := net.ListenUnixgram(network, laddr)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
