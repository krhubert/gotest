package bigtest

import (
	"math/big"
	"testing"
)

// ParseFloat is like f.Parse(s, base) with f set to the given precision
// and rounding mode.
func ParseFloat(t testing.TB, s string, base int, prec uint, mode big.RoundingMode) (f *big.Float, b int) {
	t.Helper()
	p0, p1, err := big.ParseFloat(s, base, prec, mode)
	if err != nil {
		t.Fatal(err)
	}
	return p0, p1
}
