package pproftest

import (
	"io"
	"runtime/pprof"
	"testing"
)

// WriteHeapProfile is shorthand for [Lookup]("heap").WriteTo(w, 0).
// It is preserved for backwards compatibility.
func WriteHeapProfile(t testing.TB, w io.Writer) {
	t.Helper()
	err := pprof.WriteHeapProfile(w)
	if err != nil {
		t.Fatal(err)
	}
	return
}

// StartCPUProfile enables CPU profiling for the current process.
// While profiling, the profile will be buffered and written to w.
// StartCPUProfile returns an error if profiling is already enabled.
//
// On Unix-like systems, StartCPUProfile does not work by default for
// Go code built with -buildmode=c-archive or -buildmode=c-shared.
// StartCPUProfile relies on the SIGPROF signal, but that signal will
// be delivered to the main program's SIGPROF signal handler (if any)
// not to the one used by Go. To make it work, call [os/signal.Notify]
// for [syscall.SIGPROF], but note that doing so may break any profiling
// being done by the main program.
func StartCPUProfile(t testing.TB, w io.Writer) {
	t.Helper()
	err := pprof.StartCPUProfile(w)
	if err != nil {
		t.Fatal(err)
	}
	return
}
