package exectest

import (
	"os/exec"
	"testing"
)

// LookPath searches for an executable named file in the
// directories named by the PATH environment variable.
// If file contains a slash, it is tried directly and the PATH is not consulted.
// The result may be an absolute path or a path relative to the current directory.
func LookPath(t testing.TB, file string) string {
	t.Helper()
	p0, err := exec.LookPath(file)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
