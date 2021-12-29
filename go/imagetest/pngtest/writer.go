package pngtest

import (
	"image"
	"image/png"
	"io"
	"testing"
)

// Encode writes the Image m to w in PNG format. Any Image may be
// encoded, but images that are not [image.NRGBA] might be encoded lossily.
func Encode(t testing.TB, w io.Writer, m image.Image) {
	t.Helper()
	err := png.Encode(w, m)
	if err != nil {
		t.Fatal(err)
	}
	return
}
