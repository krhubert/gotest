package nettest

import (
	"net"
	"testing"
)

// ResolveTCPAddr returns an address of TCP end point.
//
// The network must be a TCP network name.
//
// If the host in the address parameter is not a literal IP address or
// the port is not a literal port number, ResolveTCPAddr resolves the
// address to an address of TCP end point.
// Otherwise, it parses the address as a pair of literal IP address
// and port number.
// The address parameter can use a host name, but this is not
// recommended, because it will return at most one of the host name's
// IP addresses.
//
// See func [Dial] for a description of the network and address
// parameters.
func ResolveTCPAddr(t testing.TB, network, address string) *net.TCPAddr {
	t.Helper()
	p0, err := net.ResolveTCPAddr(network, address)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// DialTCP acts like [Dial] for TCP networks.
//
// The network must be a TCP network name; see func Dial for details.
//
// If laddr is nil, a local address is automatically chosen.
// If the IP field of raddr is nil or an unspecified IP address, the
// local system is assumed.
func DialTCP(t testing.TB, network string, laddr, raddr *net.TCPAddr) *net.TCPConn {
	t.Helper()
	p0, err := net.DialTCP(network, laddr, raddr)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// ListenTCP acts like [Listen] for TCP networks.
//
// The network must be a TCP network name; see func Dial for details.
//
// If the IP field of laddr is nil or an unspecified IP address,
// ListenTCP listens on all available unicast and anycast IP addresses
// of the local system.
// If the Port field of laddr is 0, a port number is automatically
// chosen.
func ListenTCP(t testing.TB, network string, laddr *net.TCPAddr) *net.TCPListener {
	t.Helper()
	p0, err := net.ListenTCP(network, laddr)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
