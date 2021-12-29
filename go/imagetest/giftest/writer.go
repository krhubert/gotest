package giftest

import (
	"image"
	"image/gif"
	"io"
	"testing"
)

// EncodeAll writes the images in g to w in GIF format with the
// given loop count and delay between frames.
func EncodeAll(t testing.TB, w io.Writer, g *gif.GIF) {
	t.Helper()
	err := gif.EncodeAll(w, g)
	if err != nil {
		t.Fatal(err)
	}
	return
}

// Encode writes the Image m to w in GIF format.
func Encode(t testing.TB, w io.Writer, m image.Image, o *gif.Options) {
	t.Helper()
	err := gif.Encode(w, m, o)
	if err != nil {
		t.Fatal(err)
	}
	return
}
