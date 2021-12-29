package nettest

import (
	"net"
	"testing"
)

// Interfaces returns a list of the system's network interfaces.
func Interfaces(t testing.TB) []net.Interface {
	t.Helper()
	p0, err := net.Interfaces()
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// InterfaceAddrs returns a list of the system's unicast interface
// addresses.
//
// The returned list does not identify the associated interface; use
// Interfaces and [Interface.Addrs] for more detail.
func InterfaceAddrs(t testing.TB) []net.Addr {
	t.Helper()
	p0, err := net.InterfaceAddrs()
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// InterfaceByIndex returns the interface specified by index.
//
// On Solaris, it returns one of the logical network interfaces
// sharing the logical data link; for more precision use
// [InterfaceByName].
func InterfaceByIndex(t testing.TB, index int) *net.Interface {
	t.Helper()
	p0, err := net.InterfaceByIndex(index)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// InterfaceByName returns the interface specified by name.
func InterfaceByName(t testing.TB, name string) *net.Interface {
	t.Helper()
	p0, err := net.InterfaceByName(name)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
