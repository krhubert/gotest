package flagtest

import (
	"flag"
	"testing"
)

// Set sets the value of the named command-line flag.
func Set(t testing.TB, name, value string) {
	t.Helper()
	err := flag.Set(name, value)
	if err != nil {
		t.Fatal(err)
	}
	return
}
