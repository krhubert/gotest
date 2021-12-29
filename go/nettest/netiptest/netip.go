package netiptest

import (
	"net/netip"
	"testing"
)

// ParseAddr parses s as an IP address, returning the result. The string
// s can be in dotted decimal ("192.0.2.1"), IPv6 ("2001:db8::68"),
// or IPv6 with a scoped addressing zone ("fe80::1cc0:3e8c:119f:c2e1%ens18").
func ParseAddr(t testing.TB, s string) netip.Addr {
	t.Helper()
	p0, err := netip.ParseAddr(s)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// ParseAddrPort parses s as an [AddrPort].
//
// It doesn't do any name resolution: both the address and the port
// must be numeric.
func ParseAddrPort(t testing.TB, s string) netip.AddrPort {
	t.Helper()
	p0, err := netip.ParseAddrPort(s)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// ParsePrefix parses s as an IP address prefix.
// The string can be in the form "192.168.1.0/24" or "2001:db8::/32",
// the CIDR notation defined in RFC 4632 and RFC 4291.
// IPv6 zones are not permitted in prefixes, and an error will be returned if a
// zone is present.
//
// Note that masked address bits are not zeroed. Use Masked for that.
func ParsePrefix(t testing.TB, s string) netip.Prefix {
	t.Helper()
	p0, err := netip.ParsePrefix(s)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
