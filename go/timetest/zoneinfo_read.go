package timetest

import (
	"testing"
	"time"
)

// LoadLocationFromTZData returns a Location with the given name
// initialized from the IANA Time Zone database-formatted data.
// The data should be in the format of a standard IANA time zone file
// (for example, the content of /etc/localtime on Unix systems).
func LoadLocationFromTZData(t testing.TB, name string, data []byte) *time.Location {
	t.Helper()
	p0, err := time.LoadLocationFromTZData(name, data)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
