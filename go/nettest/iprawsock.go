package nettest

import (
	"net"
	"testing"
)

// ResolveIPAddr returns an address of IP end point.
//
// The network must be an IP network name.
//
// If the host in the address parameter is not a literal IP address,
// ResolveIPAddr resolves the address to an address of IP end point.
// Otherwise, it parses the address as a literal IP address.
// The address parameter can use a host name, but this is not
// recommended, because it will return at most one of the host name's
// IP addresses.
//
// See func [Dial] for a description of the network and address
// parameters.
func ResolveIPAddr(t testing.TB, network, address string) *net.IPAddr {
	t.Helper()
	p0, err := net.ResolveIPAddr(network, address)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// DialIP acts like [Dial] for IP networks.
//
// The network must be an IP network name; see func Dial for details.
//
// If laddr is nil, a local address is automatically chosen.
// If the IP field of raddr is nil or an unspecified IP address, the
// local system is assumed.
func DialIP(t testing.TB, network string, laddr, raddr *net.IPAddr) *net.IPConn {
	t.Helper()
	p0, err := net.DialIP(network, laddr, raddr)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// ListenIP acts like [ListenPacket] for IP networks.
//
// The network must be an IP network name; see func Dial for details.
//
// If the IP field of laddr is nil or an unspecified IP address,
// ListenIP listens on all available IP addresses of the local system
// except multicast IP addresses.
func ListenIP(t testing.TB, network string, laddr *net.IPAddr) *net.IPConn {
	t.Helper()
	p0, err := net.ListenIP(network, laddr)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
