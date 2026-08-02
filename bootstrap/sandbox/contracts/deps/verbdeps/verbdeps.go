package verbdeps

// This package is the bootstrap library's *copy* of the embedded Verb argv
// parser's public api. The sandbox may not import the embedded library —
// that would be a third-party import — so it restates the shape it needs
// here, field for field. The adapter, which lives outside the sandbox, is
// what fills these structs from the real library.
//
// Copying is cheap precisely because the embedded library exposes structs
// of function fields instead of interfaces: an adapter assigns the real
// library's fields straight into the copy. Only the fields this bootstrap
// demo actually exercises are mirrored — a real consumer would mirror every
// field of the embedded api.Lib it intends to call.

// Lib mirrors the fields of the embedded library's api.Lib that this
// bootstrap demo exercises: presence checks and reading the next unread
// positional argument.
type Lib struct {
	// IsPresent reports whether one of the given flag spellings is present
	// in the embedded parser's argv, marking it used on a hit.
	IsPresent func(flags []string) bool
	// GetNextStringArg returns the next not-yet-used argument in the
	// embedded parser's argv, or an error when none remain.
	GetNextStringArg func() (string, error)
}
