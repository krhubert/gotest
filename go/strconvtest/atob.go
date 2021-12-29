package strconvtest

import (
	"strconv"
	"testing"
)

// ParseBool returns the boolean value represented by the string.
// It accepts 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False.
// Any other value returns an error.
func ParseBool(t testing.TB, str string) bool {
	t.Helper()
	p0, err := strconv.ParseBool(str)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
