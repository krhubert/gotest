package exectest

import (
	"os/exec"
	"testing"
)

// LookPath searches for an executable named file in the
// directories named by the path environment variable.
// If file begins with "/", "#", "./", or "../", it is tried
// directly and the path is not consulted.
// On success, the result is an absolute path.
//
// In older versions of Go, LookPath could return a path relative to the current directory.
// As of Go 1.19, LookPath will instead return that path along with an error satisfying
// [errors.Is](err, [ErrDot]). See the package documentation for more details.
func LookPath(t testing.TB, file string) string {
	t.Helper()
	p0, err := exec.LookPath(file)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
