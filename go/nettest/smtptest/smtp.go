package smtptest

import (
	"net"
	"net/smtp"
	"testing"
)

// Dial returns a new [Client] connected to an SMTP server at addr.
// The addr must include a port, as in "mail.example.com:smtp".
func Dial(t testing.TB, addr string) *smtp.Client {
	t.Helper()
	p0, err := smtp.Dial(addr)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// NewClient returns a new [Client] using an existing connection and host as a
// server name to be used when authenticating.
func NewClient(t testing.TB, conn net.Conn, host string) *smtp.Client {
	t.Helper()
	p0, err := smtp.NewClient(conn, host)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// SendMail connects to the server at addr, switches to TLS if
// possible, authenticates with the optional mechanism a if possible,
// and then sends an email from address from, to addresses to, with
// message msg.
// The addr must include a port, as in "mail.example.com:smtp".
//
// The addresses in the to parameter are the SMTP RCPT addresses.
//
// The msg parameter should be an RFC 822-style email with headers
// first, a blank line, and then the message body. The lines of msg
// should be CRLF terminated. The msg headers should usually include
// fields such as "From", "To", "Subject", and "Cc".  Sending "Bcc"
// messages is accomplished by including an email address in the to
// parameter but not including it in the msg headers.
//
// The SendMail function and the net/smtp package are low-level
// mechanisms and provide no support for DKIM signing, MIME
// attachments (see the mime/multipart package), or other mail
// functionality. Higher-level packages exist outside of the standard
// library.
func SendMail(t testing.TB, addr string, a smtp.Auth, from string, to []string, msg []byte) {
	t.Helper()
	err := smtp.SendMail(addr, a, from, to, msg)
	if err != nil {
		t.Fatal(err)
	}
	return
}
