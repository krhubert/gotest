package debugtest

import (
	"runtime/debug"
	"testing"
)

func ParseBuildInfo(t testing.TB, data string) (bi *debug.BuildInfo) {
	t.Helper()
	p0, err := debug.ParseBuildInfo(data)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
