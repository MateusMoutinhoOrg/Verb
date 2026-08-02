package deps

// Deps is the dependency contract every adapter must satisfy. It is a
// struct of function fields, not an interface: an adapter fills every
// field with the behavior it provides, and the library calls those fields
// directly.
//
// An argv parser has exactly one OS-bound effect: reading the raw argument
// vector it must parse. Everything else — matching flags, tracking which
// items were already consumed, parsing typed values — is pure computation
// and lives in sandbox/internal/lib without needing a dependency.
type Deps struct {
	// Args returns the argument vector to parse, e.g. os.Args[1:]. It is
	// called exactly once, when lib.New builds the api.Lib, and its result
	// is snapshotted so argument positions stay stable for the whole
	// lifetime of the returned Lib.
	Args func() []string
}
