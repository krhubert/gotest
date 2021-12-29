package pngtest

import (
	"image"
	"image/png"
	"io"
	"testing"
)

// Decode reads a PNG image from r and returns it as an [image.Image].
// The type of Image returned depends on the PNG contents.
func Decode(t testing.TB, r io.Reader) image.Image {
	t.Helper()
	p0, err := png.Decode(r)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// DecodeConfig returns the color model and dimensions of a PNG image without
// decoding the entire image.
func DecodeConfig(t testing.TB, r io.Reader) image.Config {
	t.Helper()
	p0, err := png.DecodeConfig(r)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
