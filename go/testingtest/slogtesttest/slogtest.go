package slogtesttest

import (
	"log/slog"
	"testing"
	"testing/slogtest"
)

// TestHandler tests a [slog.Handler].
// If TestHandler finds any misbehaviors, it returns an error for each,
// combined into a single error with [errors.Join].
//
// TestHandler installs the given Handler in a [slog.Logger] and
// makes several calls to the Logger's output methods.
// The Handler should be enabled for levels Info and above.
//
// The results function is invoked after all such calls.
// It should return a slice of map[string]any, one for each call to a Logger output method.
// The keys and values of the map should correspond to the keys and values of the Handler's
// output. Each group in the output should be represented as its own nested map[string]any.
// The standard keys [slog.TimeKey], [slog.LevelKey] and [slog.MessageKey] should be used.
//
// If the Handler outputs JSON, then calling [encoding/json.Unmarshal] with a `map[string]any`
// will create the right data structure.
//
// If a Handler intentionally drops an attribute that is checked by a test,
// then the results function should check for its absence and add it to the map it returns.
func TestHandler(t testing.TB, h slog.Handler, results func() []map[string]any) {
	t.Helper()
	err := slogtest.TestHandler(h, results)
	if err != nil {
		t.Fatal(err)
	}
	return
}
