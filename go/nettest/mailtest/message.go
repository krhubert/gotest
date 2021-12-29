package mailtest

import (
	"io"
	"net/mail"
	"testing"
	"time"
)

// ReadMessage reads a message from r.
// The headers are parsed, and the body of the message will be available
// for reading from msg.Body.
func ReadMessage(t testing.TB, r io.Reader) (msg *mail.Message) {
	t.Helper()
	p0, err := mail.ReadMessage(r)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// ParseDate parses an RFC 5322 date string.
func ParseDate(t testing.TB, date string) time.Time {
	t.Helper()
	p0, err := mail.ParseDate(date)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// ParseAddress parses a single RFC 5322 address, e.g. "Barry Gibbs <bg@example.com>"
func ParseAddress(t testing.TB, address string) *mail.Address {
	t.Helper()
	p0, err := mail.ParseAddress(address)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// ParseAddressList parses the given string as a list of addresses.
func ParseAddressList(t testing.TB, list string) []*mail.Address {
	t.Helper()
	p0, err := mail.ParseAddressList(list)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
