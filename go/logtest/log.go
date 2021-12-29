package logtest

import (
	"log"
	"testing"
)

// Output writes the output for a logging event. The string s contains
// the text to print after the prefix specified by the flags of the
// Logger. A newline is appended if the last character of s is not
// already a newline. Calldepth is the count of the number of
// frames to skip when computing the file name and line number
// if [Llongfile] or [Lshortfile] is set; a value of 1 will print the details
// for the caller of Output.
func Output(t testing.TB, calldepth int, s string) {
	t.Helper()
	err := log.Output(calldepth, s)
	if err != nil {
		t.Fatal(err)
	}
	return
}
