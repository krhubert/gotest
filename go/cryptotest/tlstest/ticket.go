package tlstest

import (
	"crypto/tls"
	"testing"
)

// ParseSessionState parses a [SessionState] encoded by [SessionState.Bytes].
func ParseSessionState(t testing.TB, data []byte) *tls.SessionState {
	t.Helper()
	p0, err := tls.ParseSessionState(data)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// NewResumptionState returns a state value that can be returned by
// [ClientSessionCache.Get] to resume a previous session.
//
// state needs to be returned by [ParseSessionState], and the ticket and session
// state must have been returned by [ClientSessionState.ResumptionState].
func NewResumptionState(t testing.TB, ticket []byte, state *tls.SessionState) *tls.ClientSessionState {
	t.Helper()
	p0, err := tls.NewResumptionState(ticket, state)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
