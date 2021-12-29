package nettest

import (
	"net"
	"testing"
)

// ParseCIDR parses s as a CIDR notation IP address and prefix length,
// like "192.0.2.0/24" or "2001:db8::/32", as defined in
// RFC 4632 and RFC 4291.
//
// It returns the IP address and the network implied by the IP and
// prefix length.
// For example, ParseCIDR("192.0.2.1/24") returns the IP address
// 192.0.2.1 and the network 192.0.2.0/24.
func ParseCIDR(t testing.TB, s string) (net.IP, *net.IPNet) {
	t.Helper()
	p0, p1, err := net.ParseCIDR(s)
	if err != nil {
		t.Fatal(err)
	}
	return p0, p1
}
