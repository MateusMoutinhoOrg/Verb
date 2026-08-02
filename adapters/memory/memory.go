package memory

import (
	"github.com/MateusMoutinhoOrg/Verb/sandbox/contracts/deps"
)

// MemoryAdapter fills deps.Deps with a fixed, in-memory argument slice
// instead of a real process's argv. It exists for tests and scripted
// scenarios: hand it any []string literal and get a deps.Deps that parses
// exactly that, with no dependency on the actual os.Args the test binary
// was invoked with.
type MemoryAdapter struct {
	// Deps is the contract this adapter fills; its factories assign into it.
	Deps deps.Deps
	args []string
}

// ArgsFactory returns the closure that fills deps.Deps.Args, handing back
// the fixed argument slice the adapter was constructed with.
func ArgsFactory(m *MemoryAdapter) func() []string {
	return func() []string {
		return m.args
	}
}

// New creates a deps.Deps backed by the memory adapter, ready for lib.New.
// args is the fixed argument vector to parse, useful for tests that want a
// known, repeatable argv instead of the real process arguments. It builds
// the adapter instance and runs every field factory over it. Adding a field
// to deps.Deps means adding its factory call here.
func New(args []string) deps.Deps {
	adapter := &MemoryAdapter{args: args}
	adapter.Deps.Args = ArgsFactory(adapter)
	return adapter.Deps
}
