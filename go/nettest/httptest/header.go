package httptest

import (
	"net/http"
	"testing"
	"time"
)

// ParseTime parses a time header (such as the Date: header),
// trying each of the three formats allowed by HTTP/1.1:
// [TimeFormat], [time.RFC850], and [time.ANSIC].
func ParseTime(t testing.TB, text string) time.Time {
	t.Helper()
	p0, err := http.ParseTime(text)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
