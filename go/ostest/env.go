package ostest

import (
	"os"
	"testing"
)

// Setenv sets the value of the environment variable named by the key.
// It returns an error, if any.
func Setenv(t testing.TB, key, value string) {
	t.Helper()
	err := os.Setenv(key, value)
	if err != nil {
		t.Fatal(err)
	}
	return
}

// Unsetenv unsets a single environment variable.
func Unsetenv(t testing.TB, key string) {
	t.Helper()
	err := os.Unsetenv(key)
	if err != nil {
		t.Fatal(err)
	}
	return
}
