package x509test

import (
	"crypto/x509"
	"testing"
)

// SystemCertPool returns a copy of the system cert pool.
//
// On Unix systems other than macOS the environment variables SSL_CERT_FILE and
// SSL_CERT_DIR can be used to override the system default locations for the SSL
// certificate file and SSL certificate files directory, respectively. The
// latter can be a colon-separated list.
//
// Any mutations to the returned pool are not written to disk and do not affect
// any other pool returned by SystemCertPool.
//
// New changes in the system cert pool might not be reflected in subsequent calls.
func SystemCertPool(t testing.TB) *x509.CertPool {
	t.Helper()
	p0, err := x509.SystemCertPool()
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
