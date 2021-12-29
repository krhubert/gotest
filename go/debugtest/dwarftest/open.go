package dwarftest

import (
	"debug/dwarf"
	"testing"
)

// New returns a new [Data] object initialized from the given parameters.
// Rather than calling this function directly, clients should typically use
// the DWARF method of the File type of the appropriate package [debug/elf],
// [debug/macho], or [debug/pe].
//
// The []byte arguments are the data from the corresponding debug section
// in the object file; for example, for an ELF object, abbrev is the contents of
// the ".debug_abbrev" section.
func New(t testing.TB, abbrev, aranges, frame, info, line, pubnames, ranges, str []byte) *dwarf.Data {
	t.Helper()
	p0, err := dwarf.New(abbrev, aranges, frame, info, line, pubnames, ranges, str)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
