package lib

import (
	"github.com/MateusMoutinhoOrg/Verb/bootstrap/sandbox/contracts/api"
	"github.com/MateusMoutinhoOrg/Verb/bootstrap/sandbox/contracts/deps"
)

// TestFuncFactory returns the closure that fills api.Lib.TestFunc, which
// exercises the embedded library reached through the Deps: it checks for a
// verbose flag, reads the next unread positional argument, and prints both.
// The embedded library is never imported here — the adapter injects it as a
// verbdeps.Lib struct, so calling it is just calling a function field.
func TestFuncFactory(l *api.Lib) func() {
	return func() {
		parser := l.Deps.ArgvLib
		verbose := parser.IsPresent([]string{"-v", "--verbose"})

		arg, err := parser.GetNextStringArg()
		if err != nil {
			l.Deps.Println("verbose:", verbose, "next arg:", err)
			return
		}
		l.Deps.Println("verbose:", verbose, "next arg:", arg)
	}
}

// New builds the api.Lib entry point, storing the injected deps on it and
// running every lib factory over it to fill its function fields. Adding a
// function field to api.Lib means adding its factory call here.
func New(d deps.Deps) api.Lib {
	l := api.Lib{Deps: d}
	l.TestFunc = TestFuncFactory(&l)
	return l
}
