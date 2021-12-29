package ostest

import (
	"os"
	"testing"
)

// Getgroups returns a list of the numeric ids of groups that the caller belongs to.
//
// On Windows, it returns [syscall.EWINDOWS]. See the [os/user] package
// for a possible alternative.
func Getgroups(t testing.TB) []int {
	t.Helper()
	p0, err := os.Getgroups()
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
