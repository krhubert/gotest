package jpegtest

import (
	"image"
	"image/jpeg"
	"io"
	"testing"
)

// Encode writes the Image m to w in JPEG 4:2:0 baseline format with the given
// options. Default parameters are used if a nil *[Options] is passed.
func Encode(t testing.TB, w io.Writer, m image.Image, o *jpeg.Options) {
	t.Helper()
	err := jpeg.Encode(w, m, o)
	if err != nil {
		t.Fatal(err)
	}
	return
}
