package standard

import (
	"fmt"

	"github.com/MateusMoutinhoOrg/Verb/bootstrap/sandbox/contracts/deps"
	"github.com/MateusMoutinhoOrg/Verb/bootstrap/sandbox/contracts/deps/verbdeps"

	verbadapter "github.com/MateusMoutinhoOrg/Verb/adapters/standard"
	verblib "github.com/MateusMoutinhoOrg/Verb/sandbox"
)

// StandardAdapter fills deps.Deps using the Go standard library plus the
// embedded Verb argv-parser library, which ArgvLibFactory wires up once.
type StandardAdapter struct {
	// Deps is the contract this adapter fills; its factories assign into it.
	Deps deps.Deps
	// argv is the argument vector the embedded parser will parse.
	argv []string
}

// PrintlnFactory returns the closure that fills deps.Deps.Println, writing a
// line to standard output.
func PrintlnFactory(s *StandardAdapter) func(a ...any) {
	return func(a ...any) {
		fmt.Println(a...)
	}
}

// ArgvLibFactory returns the value that fills deps.Deps.ArgvLib, by
// initializing the embedded Verb argv-parser library with the embedded
// library's own standard adapter, then copying the fields this bootstrap
// demo needs onto the verbdeps.Lib the bootstrap sandbox declares. It
// returns a value instead of a closure, because ArgvLib is a struct field,
// not a function field.
//
// Both structs use identical signatures for the fields mirrored here, so
// each is assigned straight across — no bridging type needed, because both
// sides are structs of function fields rather than interfaces.
//
// Only code outside the sandbox may import the embedded library, so the
// copy happens here.
func ArgvLibFactory(s *StandardAdapter) verbdeps.Lib {
	inner := verblib.New(verbadapter.New(s.argv))
	return verbdeps.Lib{
		IsPresent:        inner.IsPresent,
		GetNextStringArg: inner.GetNextStringArg,
	}
}

// New creates a deps.Deps backed by the standard adapter, ready for lib.New.
// argv is the argument vector the embedded parser will parse — pass
// os.Args[1:] to parse the process's real command line. It builds the
// adapter instance and runs every field factory over it. Adding a field to
// deps.Deps means adding its factory call here.
func New(argv []string) deps.Deps {
	adapter := &StandardAdapter{argv: argv}
	adapter.Deps.Println = PrintlnFactory(adapter)
	adapter.Deps.ArgvLib = ArgvLibFactory(adapter)
	return adapter.Deps
}
