package plugintest

import (
	"plugin"
	"testing"
)

// Open opens a Go plugin.
// If a path has already been opened, then the existing *[Plugin] is returned.
// It is safe for concurrent use by multiple goroutines.
func Open(t testing.TB, path string) *plugin.Plugin {
	t.Helper()
	p0, err := plugin.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
