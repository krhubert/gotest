package quicktest

import (
	"testing"
	"testing/quick"
)

// Check looks for an input to f, any function that returns bool,
// such that f returns false. It calls f repeatedly, with arbitrary
// values for each argument. If f returns false on a given input,
// Check returns that input as a *[CheckError].
// For example:
//
//	func TestOddMultipleOfThree(t *testing.T) {
//		f := func(x int) bool {
//			y := OddMultipleOfThree(x)
//			return y%2 == 1 && y%3 == 0
//		}
//		if err := quick.Check(f, nil); err != nil {
//			t.Error(err)
//		}
//	}
func Check(t testing.TB, f any, config *quick.Config) {
	t.Helper()
	err := quick.Check(f, config)
	if err != nil {
		t.Fatal(err)
	}
	return
}

// CheckEqual looks for an input on which f and g return different results.
// It calls f and g repeatedly with arbitrary values for each argument.
// If f and g return different answers, CheckEqual returns a *[CheckEqualError]
// describing the input and the outputs.
func CheckEqual(t testing.TB, f, g any, config *quick.Config) {
	t.Helper()
	err := quick.CheckEqual(f, g, config)
	if err != nil {
		t.Fatal(err)
	}
	return
}
