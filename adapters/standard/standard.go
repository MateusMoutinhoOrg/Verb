package standard

import (
	"github.com/MateusMoutinhoOrg/Verb/sandbox/contracts/deps"
)

// StandardAdapter fills deps.Deps using the Go standard library only. It
// hands back whatever argument slice it was constructed with — typically
// os.Args[1:], read by the caller in main() and passed in here, since
// reading process arguments is the adapter's job, never the sandbox's.
type StandardAdapter struct {
	// Deps is the contract this adapter fills; its factories assign into it.
	Deps deps.Deps
	args []string
}

// ArgsFactory returns the closure that fills deps.Deps.Args, handing back
// the argument slice the adapter was constructed with.
func ArgsFactory(s *StandardAdapter) func() []string {
	return func() []string {
		return s.args
	}
}

// New creates a deps.Deps backed by the standard adapter, ready for lib.New.
// args is the argument vector to parse — pass os.Args[1:] to parse the
// process's real command line. It builds the adapter instance and runs
// every field factory over it. Adding a field to deps.Deps means adding its
// factory call here.
func New(args []string) deps.Deps {
	adapter := &StandardAdapter{args: args}
	adapter.Deps.Args = ArgsFactory(adapter)
	return adapter.Deps
}
