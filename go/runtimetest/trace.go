package runtimetest

import (
	"runtime"
	"testing"
)

// StartTrace enables tracing for the current process.
// While tracing, the data will be buffered and available via [ReadTrace].
// StartTrace returns an error if tracing is already enabled.
// Most clients should use the [runtime/trace] package or the [testing] package's
// -test.trace flag instead of calling StartTrace directly.
func StartTrace(t testing.TB) {
	t.Helper()
	err := runtime.StartTrace()
	if err != nil {
		t.Fatal(err)
	}
	return
}
