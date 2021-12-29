package x509test

import (
	"crypto"
	"crypto/x509"
	"crypto/x509/pkix"
	"io"
	"testing"
)

// ParsePKIXPublicKey parses a public key in PKIX, ASN.1 DER form. The encoded
// public key is a SubjectPublicKeyInfo structure (see RFC 5280, Section 4.1).
//
// It returns a *[rsa.PublicKey], *[dsa.PublicKey], *[ecdsa.PublicKey],
// [ed25519.PublicKey] (not a pointer), or *[ecdh.PublicKey] (for X25519).
// More types might be supported in the future.
//
// This kind of key is commonly encoded in PEM blocks of type "PUBLIC KEY".
func ParsePKIXPublicKey(t testing.TB, derBytes []byte) (pub any) {
	t.Helper()
	p0, err := x509.ParsePKIXPublicKey(derBytes)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// MarshalPKIXPublicKey converts a public key to PKIX, ASN.1 DER form.
// The encoded public key is a SubjectPublicKeyInfo structure
// (see RFC 5280, Section 4.1).
//
// The following key types are currently supported: *[rsa.PublicKey],
// *[ecdsa.PublicKey], [ed25519.PublicKey] (not a pointer), and *[ecdh.PublicKey].
// Unsupported key types result in an error.
//
// This kind of key is commonly encoded in PEM blocks of type "PUBLIC KEY".
func MarshalPKIXPublicKey(t testing.TB, pub any) []byte {
	t.Helper()
	p0, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// CreateCertificate creates a new X.509 v3 certificate based on a template.
// The following members of template are currently used:
//
//   - AuthorityKeyId
//   - BasicConstraintsValid
//   - CRLDistributionPoints
//   - DNSNames
//   - EmailAddresses
//   - ExcludedDNSDomains
//   - ExcludedEmailAddresses
//   - ExcludedIPRanges
//   - ExcludedURIDomains
//   - ExtKeyUsage
//   - ExtraExtensions
//   - IPAddresses
//   - IsCA
//   - IssuingCertificateURL
//   - KeyUsage
//   - MaxPathLen
//   - MaxPathLenZero
//   - NotAfter
//   - NotBefore
//   - OCSPServer
//   - PermittedDNSDomains
//   - PermittedDNSDomainsCritical
//   - PermittedEmailAddresses
//   - PermittedIPRanges
//   - PermittedURIDomains
//   - PolicyIdentifiers (see note below)
//   - Policies (see note below)
//   - SerialNumber
//   - SignatureAlgorithm
//   - Subject
//   - SubjectKeyId
//   - URIs
//   - UnknownExtKeyUsage
//
// The certificate is signed by parent. If parent is equal to template then the
// certificate is self-signed. The parameter pub is the public key of the
// certificate to be generated and priv is the private key of the signer.
//
// The returned slice is the certificate in DER encoding.
//
// The currently supported key types are *rsa.PublicKey, *ecdsa.PublicKey and
// ed25519.PublicKey. pub must be a supported key type, and priv must be a
// crypto.Signer with a supported public key.
//
// The AuthorityKeyId will be taken from the SubjectKeyId of parent, if any,
// unless the resulting certificate is self-signed. Otherwise the value from
// template will be used.
//
// If SubjectKeyId from template is empty and the template is a CA, SubjectKeyId
// will be generated from the hash of the public key.
//
// The PolicyIdentifier and Policies fields are both used to marshal certificate
// policy OIDs. By default, only the PolicyIdentifier is marshaled, but if the
// GODEBUG setting "x509usepolicies" has the value "1", the Policies field will
// be marshaled instead of the PolicyIdentifier field. The Policies field can
// be used to marshal policy OIDs which have components that are larger than 31
// bits.
func CreateCertificate(t testing.TB, rand io.Reader, template, parent *x509.Certificate, pub, priv any) []byte {
	t.Helper()
	p0, err := x509.CreateCertificate(rand, template, parent, pub, priv)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// ParseCRL parses a CRL from the given bytes. It's often the case that PEM
// encoded CRLs will appear where they should be DER encoded, so this function
// will transparently handle PEM encoding as long as there isn't any leading
// garbage.
//
// Deprecated: Use [ParseRevocationList] instead.
func ParseCRL(t testing.TB, crlBytes []byte) *pkix.CertificateList {
	t.Helper()
	p0, err := x509.ParseCRL(crlBytes)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// ParseDERCRL parses a DER encoded CRL from the given bytes.
//
// Deprecated: Use [ParseRevocationList] instead.
func ParseDERCRL(t testing.TB, derBytes []byte) *pkix.CertificateList {
	t.Helper()
	p0, err := x509.ParseDERCRL(derBytes)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// CreateCertificateRequest creates a new certificate request based on a
// template. The following members of template are used:
//
//   - SignatureAlgorithm
//   - Subject
//   - DNSNames
//   - EmailAddresses
//   - IPAddresses
//   - URIs
//   - ExtraExtensions
//   - Attributes (deprecated)
//
// priv is the private key to sign the CSR with, and the corresponding public
// key will be included in the CSR. It must implement crypto.Signer and its
// Public() method must return a *rsa.PublicKey or a *ecdsa.PublicKey or a
// ed25519.PublicKey. (A *rsa.PrivateKey, *ecdsa.PrivateKey or
// ed25519.PrivateKey satisfies this.)
//
// The returned slice is the certificate request in DER encoding.
func CreateCertificateRequest(t testing.TB, rand io.Reader, template *x509.CertificateRequest, priv any) (csr []byte) {
	t.Helper()
	p0, err := x509.CreateCertificateRequest(rand, template, priv)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// ParseCertificateRequest parses a single certificate request from the
// given ASN.1 DER data.
func ParseCertificateRequest(t testing.TB, asn1Data []byte) *x509.CertificateRequest {
	t.Helper()
	p0, err := x509.ParseCertificateRequest(asn1Data)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// CreateRevocationList creates a new X.509 v2 [Certificate] Revocation List,
// according to RFC 5280, based on template.
//
// The CRL is signed by priv which should be the private key associated with
// the public key in the issuer certificate.
//
// The issuer may not be nil, and the crlSign bit must be set in [KeyUsage] in
// order to use it as a CRL issuer.
//
// The issuer distinguished name CRL field and authority key identifier
// extension are populated using the issuer certificate. issuer must have
// SubjectKeyId set.
func CreateRevocationList(t testing.TB, rand io.Reader, template *x509.RevocationList, issuer *x509.Certificate, priv crypto.Signer) []byte {
	t.Helper()
	p0, err := x509.CreateRevocationList(rand, template, issuer, priv)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
