package mimetest

import (
	"mime"
	"testing"
)

// ExtensionsByType returns the extensions known to be associated with the MIME
// type typ. The returned extensions will each begin with a leading dot, as in
// ".html". When typ has no associated extensions, ExtensionsByType returns an
// nil slice.
func ExtensionsByType(t testing.TB, typ string) []string {
	t.Helper()
	p0, err := mime.ExtensionsByType(typ)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}

// AddExtensionType sets the MIME type associated with
// the extension ext to typ. The extension should begin with
// a leading dot, as in ".html".
func AddExtensionType(t testing.TB, ext, typ string) {
	t.Helper()
	err := mime.AddExtensionType(ext, typ)
	if err != nil {
		t.Fatal(err)
	}
	return
}
