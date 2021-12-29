package coveragetest

import (
	"io"
	"runtime/coverage"
	"testing"
)

// WriteMetaDir writes a coverage meta-data file for the currently
// running program to the directory specified in 'dir'. An error will
// be returned if the operation can't be completed successfully (for
// example, if the currently running program was not built with
// "-cover", or if the directory does not exist).
func WriteMetaDir(t testing.TB, dir string) {
	t.Helper()
	err := coverage.WriteMetaDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	return
}

// WriteMeta writes the meta-data content (the payload that would
// normally be emitted to a meta-data file) for the currently running
// program to the writer 'w'. An error will be returned if the
// operation can't be completed successfully (for example, if the
// currently running program was not built with "-cover", or if a
// write fails).
func WriteMeta(t testing.TB, w io.Writer) {
	t.Helper()
	err := coverage.WriteMeta(w)
	if err != nil {
		t.Fatal(err)
	}
	return
}

// WriteCountersDir writes a coverage counter-data file for the
// currently running program to the directory specified in 'dir'. An
// error will be returned if the operation can't be completed
// successfully (for example, if the currently running program was not
// built with "-cover", or if the directory does not exist). The
// counter data written will be a snapshot taken at the point of the
// call.
func WriteCountersDir(t testing.TB, dir string) {
	t.Helper()
	err := coverage.WriteCountersDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	return
}

// WriteCounters writes coverage counter-data content for the
// currently running program to the writer 'w'. An error will be
// returned if the operation can't be completed successfully (for
// example, if the currently running program was not built with
// "-cover", or if a write fails). The counter data written will be a
// snapshot taken at the point of the invocation.
func WriteCounters(t testing.TB, w io.Writer) {
	t.Helper()
	err := coverage.WriteCounters(w)
	if err != nil {
		t.Fatal(err)
	}
	return
}

// ClearCounters clears/resets all coverage counter variables in the
// currently running program. It returns an error if the program in
// question was not built with the "-cover" flag. Clearing of coverage
// counters is also not supported for programs not using atomic
// counter mode (see more detailed comments below for the rationale
// here).
func ClearCounters(t testing.TB) {
	t.Helper()
	err := coverage.ClearCounters()
	if err != nil {
		t.Fatal(err)
	}
	return
}
