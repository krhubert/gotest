package imagetest

import (
	"image"
	"io"
	"testing"
)

// Decode decodes an image that has been encoded in a registered format.
// The string returned is the format name used during format registration.
// Format registration is typically done by an init function in the codec-
// specific package.
func Decode(t testing.TB, r io.Reader) (image.Image, string) {
	t.Helper()
	p0, p1, err := image.Decode(r)
	if err != nil {
		t.Fatal(err)
	}
	return p0, p1
}

// DecodeConfig decodes the color model and dimensions of an image that has
// been encoded in a registered format. The string returned is the format name
// used during format registration. Format registration is typically done by
// an init function in the codec-specific package.
func DecodeConfig(t testing.TB, r io.Reader) (image.Config, string) {
	t.Helper()
	p0, p1, err := image.DecodeConfig(r)
	if err != nil {
		t.Fatal(err)
	}
	return p0, p1
}
