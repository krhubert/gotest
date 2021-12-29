package xmltest

import (
	"encoding/xml"
	"io"
	"testing"
)

// EscapeText writes to w the properly escaped XML equivalent
// of the plain text data s.
func EscapeText(t testing.TB, w io.Writer, s []byte) {
	t.Helper()
	err := xml.EscapeText(w, s)
	if err != nil {
		t.Fatal(err)
	}
	return
}
