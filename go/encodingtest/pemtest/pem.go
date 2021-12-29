package pemtest

import (
	"encoding/pem"
	"io"
	"testing"
)

// Encode writes the PEM encoding of b to out.
func Encode(t testing.TB, out io.Writer, b *pem.Block) {
	t.Helper()
	err := pem.Encode(out, b)
	if err != nil {
		t.Fatal(err)
	}
	return
}
