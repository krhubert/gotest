package nettest

import (
	"net"
	"testing"
)

// LookupHost looks up the given host using the local resolver.
// It returns a slice of that host's addresses.
//
// LookupHost uses [context.Background] internally; to specify the context, use
// [Resolver.LookupHost].
func LookupHost(t testing.TB, host string) (addrs []string) {
	t.Helper()
	p0, err := net.LookupHost(host)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// LookupIP looks up host using the local resolver.
// It returns a slice of that host's IPv4 and IPv6 addresses.
func LookupIP(t testing.TB, host string) []net.IP {
	t.Helper()
	p0, err := net.LookupIP(host)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// LookupPort looks up the port for the given network and service.
//
// LookupPort uses [context.Background] internally; to specify the context, use
// [Resolver.LookupPort].
func LookupPort(t testing.TB, network, service string) (port int) {
	t.Helper()
	p0, err := net.LookupPort(network, service)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// LookupCNAME returns the canonical name for the given host.
// Callers that do not care about the canonical name can call
// [LookupHost] or [LookupIP] directly; both take care of resolving
// the canonical name as part of the lookup.
//
// A canonical name is the final name after following zero
// or more CNAME records.
// LookupCNAME does not return an error if host does not
// contain DNS "CNAME" records, as long as host resolves to
// address records.
//
// The returned canonical name is validated to be a properly
// formatted presentation-format domain name.
//
// LookupCNAME uses [context.Background] internally; to specify the context, use
// [Resolver.LookupCNAME].
func LookupCNAME(t testing.TB, host string) (cname string) {
	t.Helper()
	p0, err := net.LookupCNAME(host)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// LookupSRV tries to resolve an [SRV] query of the given service,
// protocol, and domain name. The proto is "tcp" or "udp".
// The returned records are sorted by priority and randomized
// by weight within a priority.
//
// LookupSRV constructs the DNS name to look up following RFC 2782.
// That is, it looks up _service._proto.name. To accommodate services
// publishing SRV records under non-standard names, if both service
// and proto are empty strings, LookupSRV looks up name directly.
//
// The returned service names are validated to be properly
// formatted presentation-format domain names. If the response contains
// invalid names, those records are filtered out and an error
// will be returned alongside the remaining results, if any.
func LookupSRV(t testing.TB, service, proto, name string) (cname string, addrs []*net.SRV) {
	t.Helper()
	p0, p1, err := net.LookupSRV(service, proto, name)
	if err != nil {
		t.Fatal(err)
	}
	return p0, p1
}

// LookupMX returns the DNS MX records for the given domain name sorted by preference.
//
// The returned mail server names are validated to be properly
// formatted presentation-format domain names. If the response contains
// invalid names, those records are filtered out and an error
// will be returned alongside the remaining results, if any.
//
// LookupMX uses [context.Background] internally; to specify the context, use
// [Resolver.LookupMX].
func LookupMX(t testing.TB, name string) []*net.MX {
	t.Helper()
	p0, err := net.LookupMX(name)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// LookupNS returns the DNS NS records for the given domain name.
//
// The returned name server names are validated to be properly
// formatted presentation-format domain names. If the response contains
// invalid names, those records are filtered out and an error
// will be returned alongside the remaining results, if any.
//
// LookupNS uses [context.Background] internally; to specify the context, use
// [Resolver.LookupNS].
func LookupNS(t testing.TB, name string) []*net.NS {
	t.Helper()
	p0, err := net.LookupNS(name)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// LookupTXT returns the DNS TXT records for the given domain name.
//
// LookupTXT uses [context.Background] internally; to specify the context, use
// [Resolver.LookupTXT].
func LookupTXT(t testing.TB, name string) []string {
	t.Helper()
	p0, err := net.LookupTXT(name)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// LookupAddr performs a reverse lookup for the given address, returning a list
// of names mapping to that address.
//
// The returned names are validated to be properly formatted presentation-format
// domain names. If the response contains invalid names, those records are filtered
// out and an error will be returned alongside the remaining results, if any.
//
// When using the host C library resolver, at most one result will be
// returned. To bypass the host resolver, use a custom [Resolver].
//
// LookupAddr uses [context.Background] internally; to specify the context, use
// [Resolver.LookupAddr].
func LookupAddr(t testing.TB, addr string) (names []string) {
	t.Helper()
	p0, err := net.LookupAddr(addr)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
