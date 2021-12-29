package syslogtest

import (
	"log"
	"log/syslog"
	"testing"
)

// New establishes a new connection to the system log daemon. Each
// write to the returned writer sends a log message with the given
// priority (a combination of the syslog facility and severity) and
// prefix tag. If tag is empty, the [os.Args][0] is used.
func New(t testing.TB, priority syslog.Priority, tag string) *syslog.Writer {
	t.Helper()
	p0, err := syslog.New(priority, tag)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// Dial establishes a connection to a log daemon by connecting to
// address raddr on the specified network. Each write to the returned
// writer sends a log message with the facility and severity
// (from priority) and tag. If tag is empty, the [os.Args][0] is used.
// If network is empty, Dial will connect to the local syslog server.
// Otherwise, see the documentation for net.Dial for valid values
// of network and raddr.
func Dial(t testing.TB, network, raddr string, priority syslog.Priority, tag string) *syslog.Writer {
	t.Helper()
	p0, err := syslog.Dial(network, raddr, priority, tag)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// NewLogger creates a [log.Logger] whose output is written to the
// system log service with the specified priority, a combination of
// the syslog facility and severity. The logFlag argument is the flag
// set passed through to [log.New] to create the Logger.
func NewLogger(t testing.TB, p syslog.Priority, logFlag int) *log.Logger {
	t.Helper()
	p0, err := syslog.NewLogger(p, logFlag)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
