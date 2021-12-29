package tracetest

import (
	"io"
	"runtime/trace"
	"testing"
)

// Start enables tracing for the current program.
// While tracing, the trace will be buffered and written to w.
// Start returns an error if tracing is already enabled.
func Start(t testing.TB, w io.Writer) {
	t.Helper()
	err := trace.Start(w)
	if err != nil {
		t.Fatal(err)
	}
	return
}
