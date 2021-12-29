package giftest

import (
	"image"
	"image/gif"
	"io"
	"testing"
)

// Decode reads a GIF image from r and returns the first embedded
// image as an [image.Image].
func Decode(t testing.TB, r io.Reader) image.Image {
	t.Helper()
	p0, err := gif.Decode(r)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// DecodeAll reads a GIF image from r and returns the sequential frames
// and timing information.
func DecodeAll(t testing.TB, r io.Reader) *gif.GIF {
	t.Helper()
	p0, err := gif.DecodeAll(r)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// DecodeConfig returns the global color model and dimensions of a GIF image
// without decoding the entire image.
func DecodeConfig(t testing.TB, r io.Reader) image.Config {
	t.Helper()
	p0, err := gif.DecodeConfig(r)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
