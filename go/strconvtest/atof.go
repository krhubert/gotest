package strconvtest

import (
	"strconv"
	"testing"
)

// ParseFloat converts the string s to a floating-point number
// with the precision specified by bitSize: 32 for float32, or 64 for float64.
// When bitSize=32, the result still has type float64, but it will be
// convertible to float32 without changing its value.
//
// ParseFloat accepts decimal and hexadecimal floating-point numbers
// as defined by the Go syntax for [floating-point literals].
// If s is well-formed and near a valid floating-point number,
// ParseFloat returns the nearest floating-point number rounded
// using IEEE754 unbiased rounding.
// (Parsing a hexadecimal floating-point value only rounds when
// there are more bits in the hexadecimal representation than
// will fit in the mantissa.)
//
// The errors that ParseFloat returns have concrete type *NumError
// and include err.Num = s.
//
// If s is not syntactically well-formed, ParseFloat returns err.Err = ErrSyntax.
//
// If s is syntactically well-formed but is more than 1/2 ULP
// away from the largest floating point number of the given size,
// ParseFloat returns f = ±Inf, err.Err = ErrRange.
//
// ParseFloat recognizes the string "NaN", and the (possibly signed) strings "Inf" and "Infinity"
// as their respective special floating point values. It ignores case when matching.
//
// [floating-point literals]: https://go.dev/ref/spec#Floating-point_literals
func ParseFloat(t testing.TB, s string, bitSize int) float64 {
	t.Helper()
	p0, err := strconv.ParseFloat(s, bitSize)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
