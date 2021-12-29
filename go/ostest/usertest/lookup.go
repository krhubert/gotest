package usertest

import (
	"os/user"
	"testing"
)

// Current returns the current user.
//
// The first call will cache the current user information.
// Subsequent calls will return the cached value and will not reflect
// changes to the current user.
func Current(t testing.TB) *user.User {
	t.Helper()
	p0, err := user.Current()
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// Lookup looks up a user by username. If the user cannot be found, the
// returned error is of type [UnknownUserError].
func Lookup(t testing.TB, username string) *user.User {
	t.Helper()
	p0, err := user.Lookup(username)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// LookupId looks up a user by userid. If the user cannot be found, the
// returned error is of type [UnknownUserIdError].
func LookupId(t testing.TB, uid string) *user.User {
	t.Helper()
	p0, err := user.LookupId(uid)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// LookupGroup looks up a group by name. If the group cannot be found, the
// returned error is of type [UnknownGroupError].
func LookupGroup(t testing.TB, name string) *user.Group {
	t.Helper()
	p0, err := user.LookupGroup(name)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// LookupGroupId looks up a group by groupid. If the group cannot be found, the
// returned error is of type [UnknownGroupIdError].
func LookupGroupId(t testing.TB, gid string) *user.Group {
	t.Helper()
	p0, err := user.LookupGroupId(gid)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
