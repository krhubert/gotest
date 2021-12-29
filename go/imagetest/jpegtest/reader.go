package jpegtest

import (
	"image"
	"image/jpeg"
	"io"
	"testing"
)

// Decode reads a JPEG image from r and returns it as an [image.Image].
func Decode(t testing.TB, r io.Reader) image.Image {
	t.Helper()
	p0, err := jpeg.Decode(r)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// DecodeConfig returns the color model and dimensions of a JPEG image without
// decoding the entire image.
func DecodeConfig(t testing.TB, r io.Reader) image.Config {
	t.Helper()
	p0, err := jpeg.DecodeConfig(r)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
