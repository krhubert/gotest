package bufiotest

import (
	"bufio"
	"testing"
)

// ScanBytes is a split function for a [Scanner] that returns each byte as a token.
func ScanBytes(t testing.TB, data []byte, atEOF bool) (advance int, token []byte) {
	t.Helper()
	p0, p1, err := bufio.ScanBytes(data, atEOF)
	if err != nil {
		t.Fatal(err)
	}
	return p0, p1
}

// ScanRunes is a split function for a [Scanner] that returns each
// UTF-8-encoded rune as a token. The sequence of runes returned is
// equivalent to that from a range loop over the input as a string, which
// means that erroneous UTF-8 encodings translate to U+FFFD = "\xef\xbf\xbd".
// Because of the Scan interface, this makes it impossible for the client to
// distinguish correctly encoded replacement runes from encoding errors.
func ScanRunes(t testing.TB, data []byte, atEOF bool) (advance int, token []byte) {
	t.Helper()
	p0, p1, err := bufio.ScanRunes(data, atEOF)
	if err != nil {
		t.Fatal(err)
	}
	return p0, p1
}

// ScanLines is a split function for a [Scanner] that returns each line of
// text, stripped of any trailing end-of-line marker. The returned line may
// be empty. The end-of-line marker is one optional carriage return followed
// by one mandatory newline. In regular expression notation, it is `\r?\n`.
// The last non-empty line of input will be returned even if it has no
// newline.
func ScanLines(t testing.TB, data []byte, atEOF bool) (advance int, token []byte) {
	t.Helper()
	p0, p1, err := bufio.ScanLines(data, atEOF)
	if err != nil {
		t.Fatal(err)
	}
	return p0, p1
}

// ScanWords is a split function for a [Scanner] that returns each
// space-separated word of text, with surrounding spaces deleted. It will
// never return an empty string. The definition of space is set by
// unicode.IsSpace.
func ScanWords(t testing.TB, data []byte, atEOF bool) (advance int, token []byte) {
	t.Helper()
	p0, p1, err := bufio.ScanWords(data, atEOF)
	if err != nil {
		t.Fatal(err)
	}
	return p0, p1
}
