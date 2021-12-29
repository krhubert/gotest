package cookiejartest

import (
	"net/http/cookiejar"
	"testing"
)

// New returns a new cookie jar. A nil [*Options] is equivalent to a zero
// Options.
func New(t testing.TB, o *cookiejar.Options) *cookiejar.Jar {
	t.Helper()
	p0, err := cookiejar.New(o)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
