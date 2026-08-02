//go:build ignore

// This file is an illustrative sample, not part of the build.
// It shows the factory pattern used inside sandbox/internal/, where the
// carrier is an api struct being filled.
package example_factories

import (
	"github.com/MateusMoutinhoOrg/Verb/sandbox/contracts/api"
)

// ExampleLibFunctionFactory returns the closure that fills
// api.Lib.ExampleLibFunction, closed over l, so it reads the struct's own
// state through l.Multiplier at call time.
func ExampleLibFunctionFactory(l *api.Lib) func() int {
	return func() int {
		return l.Multiplier + 1
	}
}

// New builds the api.Lib entry point and runs every lib factory over it,
// assigning each return value into its matching field. It is the factory
// aggregate — a field left unassigned here stays nil and panics on first
// call.
func New(multiplier int) api.Lib {
	l := api.Lib{Multiplier: multiplier}
	l.ExampleLibFunction = ExampleLibFunctionFactory(&l)
	return l
}
