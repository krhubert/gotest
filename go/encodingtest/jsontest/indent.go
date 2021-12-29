package jsontest

import (
	"bytes"
	"encoding/json"
	"testing"
)

// Compact appends to dst the JSON-encoded src with
// insignificant space characters elided.
func Compact(t testing.TB, dst *bytes.Buffer, src []byte) {
	t.Helper()
	err := json.Compact(dst, src)
	if err != nil {
		t.Fatal(err)
	}
	return
}

// Indent appends to dst an indented form of the JSON-encoded src.
// Each element in a JSON object or array begins on a new,
// indented line beginning with prefix followed by one or more
// copies of indent according to the indentation nesting.
// The data appended to dst does not begin with the prefix nor
// any indentation, to make it easier to embed inside other formatted JSON data.
// Although leading space characters (space, tab, carriage return, newline)
// at the beginning of src are dropped, trailing space characters
// at the end of src are preserved and copied to dst.
// For example, if src has no trailing spaces, neither will dst;
// if src ends in a trailing newline, so will dst.
func Indent(t testing.TB, dst *bytes.Buffer, src []byte, prefix, indent string) {
	t.Helper()
	err := json.Indent(dst, src, prefix, indent)
	if err != nil {
		t.Fatal(err)
	}
	return
}
