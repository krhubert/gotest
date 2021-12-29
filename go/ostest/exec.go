package ostest

import (
	"os"
	"testing"
)

// FindProcess looks for a running process by its pid.
//
// The [Process] it returns can be used to obtain information
// about the underlying operating system process.
//
// On Unix systems, FindProcess always succeeds and returns a Process
// for the given pid, regardless of whether the process exists. To test whether
// the process actually exists, see whether p.Signal(syscall.Signal(0)) reports
// an error.
func FindProcess(t testing.TB, pid int) *os.Process {
	t.Helper()
	p0, err := os.FindProcess(pid)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// StartProcess starts a new process with the program, arguments and attributes
// specified by name, argv and attr. The argv slice will become [os.Args] in the
// new process, so it normally starts with the program name.
//
// If the calling goroutine has locked the operating system thread
// with [runtime.LockOSThread] and modified any inheritable OS-level
// thread state (for example, Linux or Plan 9 name spaces), the new
// process will inherit the caller's thread state.
//
// StartProcess is a low-level interface. The [os/exec] package provides
// higher-level interfaces.
//
// If there is an error, it will be of type [*PathError].
func StartProcess(t testing.TB, name string, argv []string, attr *os.ProcAttr) *os.Process {
	t.Helper()
	p0, err := os.StartProcess(name, argv, attr)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
