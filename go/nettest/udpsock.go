package nettest

import (
	"net"
	"testing"
)

// ResolveUDPAddr returns an address of UDP end point.
//
// The network must be a UDP network name.
//
// If the host in the address parameter is not a literal IP address or
// the port is not a literal port number, ResolveUDPAddr resolves the
// address to an address of UDP end point.
// Otherwise, it parses the address as a pair of literal IP address
// and port number.
// The address parameter can use a host name, but this is not
// recommended, because it will return at most one of the host name's
// IP addresses.
//
// See func Dial for a description of the network and address
// parameters.
func ResolveUDPAddr(t testing.TB, network, address string) *net.UDPAddr {
	t.Helper()
	p0, err := net.ResolveUDPAddr(network, address)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// DialUDP acts like Dial for UDP networks.
//
// The network must be a UDP network name; see func Dial for details.
//
// If laddr is nil, a local address is automatically chosen.
// If the IP field of raddr is nil or an unspecified IP address, the
// local system is assumed.
func DialUDP(t testing.TB, network string, laddr, raddr *net.UDPAddr) *net.UDPConn {
	t.Helper()
	p0, err := net.DialUDP(network, laddr, raddr)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// ListenUDP acts like ListenPacket for UDP networks.
//
// The network must be a UDP network name; see func Dial for details.
//
// If the IP field of laddr is nil or an unspecified IP address,
// ListenUDP listens on all available IP addresses of the local system
// except multicast IP addresses.
// If the Port field of laddr is 0, a port number is automatically
// chosen.
func ListenUDP(t testing.TB, network string, laddr *net.UDPAddr) *net.UDPConn {
	t.Helper()
	p0, err := net.ListenUDP(network, laddr)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// ListenMulticastUDP acts like ListenPacket for UDP networks but
// takes a group address on a specific network interface.
//
// The network must be a UDP network name; see func Dial for details.
//
// ListenMulticastUDP listens on all available IP addresses of the
// local system including the group, multicast IP address.
// If ifi is nil, ListenMulticastUDP uses the system-assigned
// multicast interface, although this is not recommended because the
// assignment depends on platforms and sometimes it might require
// routing configuration.
// If the Port field of gaddr is 0, a port number is automatically
// chosen.
//
// ListenMulticastUDP is just for convenience of simple, small
// applications. There are golang.org/x/net/ipv4 and
// golang.org/x/net/ipv6 packages for general purpose uses.
//
// Note that ListenMulticastUDP will set the IP_MULTICAST_LOOP socket option
// to 0 under IPPROTO_IP, to disable loopback of multicast packets.
func ListenMulticastUDP(t testing.TB, network string, ifi *net.Interface, gaddr *net.UDPAddr) *net.UDPConn {
	t.Helper()
	p0, err := net.ListenMulticastUDP(network, ifi, gaddr)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
