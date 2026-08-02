//go:build ignore

// This file is an illustrative sample, not part of the build.
package lib

import (
	"github.com/MateusMoutinhoOrg/Verb/sandbox/contracts/api"
)

// ExampleFunctionFactory returns the closure that fills
// api.Lib.ExampleFunction, closed over l, so it reads the struct's own
// state through l.Multiplier at call time.
func ExampleFunctionFactory(l *api.Lib) func(i int) int {
	return func(i int) int {
		return l.Multiplier*i + 10
	}
}

// New builds the api.Lib entry point, storing multiplier on it and running
// every lib factory over it, assigning each return value into its matching
// function field. Adding a function field to api.Lib means adding its
// factory call and assignment here — an unlisted field stays nil and
// panics on first call.
func New(multiplier int) api.Lib {
	l := api.Lib{Multiplier: multiplier}
	l.ExampleFunction = ExampleFunctionFactory(&l)
	l.NewExampleObject = NewExampleObjectFactory(&l)
	return l
}
