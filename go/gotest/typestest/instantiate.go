package typestest

import (
	"go/types"
	"testing"
)

// Instantiate instantiates the type orig with the given type arguments targs.
// orig must be an *Alias, *Named, or *Signature type. If there is no error,
// the resulting Type is an instantiated type of the same kind (*Alias, *Named
// or *Signature, respectively).
//
// Methods attached to a *Named type are also instantiated, and associated with
// a new *Func that has the same position as the original method, but nil function
// scope.
//
// If ctxt is non-nil, it may be used to de-duplicate the instance against
// previous instances with the same identity. As a special case, generic
// *Signature origin types are only considered identical if they are pointer
// equivalent, so that instantiating distinct (but possibly identical)
// signatures will yield different instances. The use of a shared context does
// not guarantee that identical instances are deduplicated in all cases.
//
// If validate is set, Instantiate verifies that the number of type arguments
// and parameters match, and that the type arguments satisfy their respective
// type constraints. If verification fails, the resulting error may wrap an
// *ArgumentError indicating which type argument did not satisfy its type parameter
// constraint, and why.
//
// If validate is not set, Instantiate does not verify the type argument count
// or whether the type arguments satisfy their constraints. Instantiate is
// guaranteed to not return an error, but may panic. Specifically, for
// *Signature types, Instantiate will panic immediately if the type argument
// count is incorrect; for *Named types, a panic may occur later inside the
// *Named API.
func Instantiate(t testing.TB, ctxt *types.Context, orig types.Type, targs []types.Type, validate bool) types.Type {
	t.Helper()
	p0, err := types.Instantiate(ctxt, orig, targs, validate)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
