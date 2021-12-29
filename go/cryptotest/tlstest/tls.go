package tlstest

import (
	"crypto/tls"
	"net"
	"testing"
)

// Listen creates a TLS listener accepting connections on the
// given network address using net.Listen.
// The configuration config must be non-nil and must include
// at least one certificate or else set GetCertificate.
func Listen(t testing.TB, network, laddr string, config *tls.Config) net.Listener {
	t.Helper()
	p0, err := tls.Listen(network, laddr, config)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// DialWithDialer connects to the given network address using dialer.Dial and
// then initiates a TLS handshake, returning the resulting TLS connection. Any
// timeout or deadline given in the dialer apply to connection and TLS
// handshake as a whole.
//
// DialWithDialer interprets a nil configuration as equivalent to the zero
// configuration; see the documentation of [Config] for the defaults.
//
// DialWithDialer uses context.Background internally; to specify the context,
// use [Dialer.DialContext] with NetDialer set to the desired dialer.
func DialWithDialer(t testing.TB, dialer *net.Dialer, network, addr string, config *tls.Config) *tls.Conn {
	t.Helper()
	p0, err := tls.DialWithDialer(dialer, network, addr, config)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// Dial connects to the given network address using net.Dial
// and then initiates a TLS handshake, returning the resulting
// TLS connection.
// Dial interprets a nil configuration as equivalent to
// the zero configuration; see the documentation of Config
// for the defaults.
func Dial(t testing.TB, network, addr string, config *tls.Config) *tls.Conn {
	t.Helper()
	p0, err := tls.Dial(network, addr, config)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// LoadX509KeyPair reads and parses a public/private key pair from a pair of
// files. The files must contain PEM encoded data. The certificate file may
// contain intermediate certificates following the leaf certificate to form a
// certificate chain. On successful return, Certificate.Leaf will be populated.
//
// Before Go 1.23 Certificate.Leaf was left nil, and the parsed certificate was
// discarded. This behavior can be re-enabled by setting "x509keypairleaf=0"
// in the GODEBUG environment variable.
func LoadX509KeyPair(t testing.TB, certFile, keyFile string) tls.Certificate {
	t.Helper()
	p0, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// X509KeyPair parses a public/private key pair from a pair of
// PEM encoded data. On successful return, Certificate.Leaf will be populated.
//
// Before Go 1.23 Certificate.Leaf was left nil, and the parsed certificate was
// discarded. This behavior can be re-enabled by setting "x509keypairleaf=0"
// in the GODEBUG environment variable.
func X509KeyPair(t testing.TB, certPEMBlock, keyPEMBlock []byte) tls.Certificate {
	t.Helper()
	p0, err := tls.X509KeyPair(certPEMBlock, keyPEMBlock)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
