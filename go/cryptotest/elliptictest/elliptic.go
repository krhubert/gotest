package elliptictest

import (
	"crypto/elliptic"
	"io"
	"math/big"
	"testing"
)

// GenerateKey returns a public/private key pair. The private key is
// generated using the given reader, which must return random data.
//
// Deprecated: for ECDH, use the GenerateKey methods of the [crypto/ecdh] package;
// for ECDSA, use the GenerateKey function of the crypto/ecdsa package.
func GenerateKey(t testing.TB, curve elliptic.Curve, rand io.Reader) (priv []byte, x, y *big.Int) {
	t.Helper()
	p0, p1, p2, err := elliptic.GenerateKey(curve, rand)
	if err != nil {
		t.Fatal(err)
	}
	return p0, p1, p2
}
