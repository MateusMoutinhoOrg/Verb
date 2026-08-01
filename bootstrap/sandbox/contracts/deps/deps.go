package deps

import "github.com/MateusMoutinhoOrg/Verb/bootstrap/sandbox/contracts/deps/verbdeps"

// Deps is the dependency contract every bootstrap adapter must satisfy.
// Like every contract in this template it is a struct of fields, not an
// interface.
//
// CacheLib shows what changes when the dependency is another library built
// with this same pattern: the whole library arrives as one struct field,
// with no getter method and no bridging type around it.
type Deps struct {
	// Println writes a line to the library's output.
	Println func(a ...any)
	// CacheLib is the embedded Verb cache library, already initialized by
	// the adapter.
	CacheLib verbdeps.Lib
}
