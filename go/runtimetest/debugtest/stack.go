package debugtest

import (
	"os"
	"runtime/debug"
	"testing"
)

// SetCrashOutput configures a single additional file where unhandled
// panics and other fatal errors are printed, in addition to standard error.
// There is only one additional file: calling SetCrashOutput again overrides
// any earlier call.
// SetCrashOutput duplicates f's file descriptor, so the caller may safely
// close f as soon as SetCrashOutput returns.
// To disable this additional crash output, call SetCrashOutput(nil).
// If called concurrently with a crash, some in-progress output may be written
// to the old file even after an overriding SetCrashOutput returns.
func SetCrashOutput(t testing.TB, f *os.File, opts debug.CrashOptions) {
	t.Helper()
	err := debug.SetCrashOutput(f, opts)
	if err != nil {
		t.Fatal(err)
	}
	return
}
