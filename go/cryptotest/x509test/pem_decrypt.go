package x509test

import (
	"crypto/x509"
	"encoding/pem"
	"io"
	"testing"
)

// DecryptPEMBlock takes a PEM block encrypted according to RFC 1423 and the
// password used to encrypt it and returns a slice of decrypted DER encoded
// bytes. It inspects the DEK-Info header to determine the algorithm used for
// decryption. If no DEK-Info header is present, an error is returned. If an
// incorrect password is detected an [IncorrectPasswordError] is returned. Because
// of deficiencies in the format, it's not always possible to detect an
// incorrect password. In these cases no error will be returned but the
// decrypted DER bytes will be random noise.
//
// Deprecated: Legacy PEM encryption as specified in RFC 1423 is insecure by
// design. Since it does not authenticate the ciphertext, it is vulnerable to
// padding oracle attacks that can let an attacker recover the plaintext.
func DecryptPEMBlock(t testing.TB, b *pem.Block, password []byte) []byte {
	t.Helper()
	p0, err := x509.DecryptPEMBlock(b, password)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// EncryptPEMBlock returns a PEM block of the specified type holding the
// given DER encoded data encrypted with the specified algorithm and
// password according to RFC 1423.
//
// Deprecated: Legacy PEM encryption as specified in RFC 1423 is insecure by
// design. Since it does not authenticate the ciphertext, it is vulnerable to
// padding oracle attacks that can let an attacker recover the plaintext.
func EncryptPEMBlock(t testing.TB, rand io.Reader, blockType string, data, password []byte, alg x509.PEMCipher) *pem.Block {
	t.Helper()
	p0, err := x509.EncryptPEMBlock(rand, blockType, data, password, alg)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
