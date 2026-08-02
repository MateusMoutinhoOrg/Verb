package api

import (
	"time"
)

// Lib is the entry point handed back by lib.New: an argument-vector (argv)
// parser exposed as a struct of function fields, filled by the factories in
// sandbox/internal/lib. Because it is a struct and not an interface, a
// consumer that itself uses this pattern can copy the shape of Lib into its
// own deps contract and receive the whole parser as a single injected field.
//
// # The Unused Mechanic
//
// Every argument in Args starts out unread. Calling any Get* function or
// IsPresent — with one exception, see GetOptionsSize and GetKeyValuesSize
// below — marks the argument(s) it matched as read by flipping the matching
// index of Used to true. This lets a caller build a complete CLI: after every
// expected flag and option has been read, whatever is left over in Args is
// exactly the positional arguments the caller never explicitly asked for
// (e.g. a trailing filename), retrievable in order with GetNextStringArg and
// its typed variants. See docs/Explanations/UnnusedMechanic.md for a worked
// example.
//
// # Typed Getters
//
// Every family below (Option, Arg, NextArg, KeyValues) is exposed once per
// supported value type: String (the raw text, no parsing), Int (parsed with
// strconv.Atoi), Double (parsed with strconv.ParseFloat as a float64), and
// Timestamp (parsed with time.Parse using time.RFC3339, e.g.
// "2024-01-02T15:04:05Z"). A typed getter marks its matched argument(s) as
// used exactly like GetString*, even when parsing then fails — the argument
// was still found and read, only its value turned out malformed.
//
// # Errors
//
// Every Get* function returns an error instead of a bool or a panic when it
// cannot produce a value: the flag/argument/key was not present, the
// occurrence index was out of range, an option's value was missing, a
// key=value pair had an empty value, or a typed parse failed. The two
// *Size functions are the only Get-shaped fields that return a plain int,
// because counting matches can never fail. IsPresent returns a plain bool
// for the same reason — checking presence cannot fail either.
type Lib struct {
	// Args is the argument vector being parsed, the same slice lib.New was
	// called with. Every index-based function (GetStringArg, GetStringOption,
	// GetStringKeyValues, ...) refers to positions in this slice. It is
	// exported so sandbox/internal/lib can populate it from another package,
	// but callers should treat it as read-only: mutating it after
	// construction leaves Used out of sync and produces undefined matching
	// behavior.
	Args []string

	// Used tracks, index for index against Args, which arguments have
	// already been matched by a previous call. Used[i] is true once Args[i]
	// has been consumed by any Get* function or by IsPresent. It starts
	// entirely false and only grows more true over the Lib's lifetime — the
	// Unused Mechanic reads it to find the next never-consumed positional
	// argument. Exported for the same cross-package reason as Args; treat it
	// as read-only from outside sandbox/internal/lib.
	Used []bool

	// IsPresent reports whether any of the given flag spellings (e.g.
	// []string{"-q", "--quiet"}) occurs anywhere in the unread portion of
	// Args. On a match it marks that single argument as used and returns
	// true; if none of the flags is found it returns false and nothing is
	// marked. It never returns an error: "not present" is a valid, expected
	// outcome, not a failure.
	IsPresent func(flags []string) bool

	// GetOptionsSize counts how many arguments in Args equal one of the
	// given flag spellings (e.g. []string{"-o", "--output"}), regardless of
	// whether they have already been marked used. It never mutates Used —
	// call it before looping over occurrence indices 0..size-1 with
	// GetStringOption (or a typed variant) to read every occurrence of a
	// repeatable option.
	GetOptionsSize func(flags []string) int

	// GetKeyValuesSize counts how many arguments in Args start with one of
	// the given key=value prefixes (e.g. []string{"user=", "username="} —
	// the separator is part of the prefix), regardless of Used. Like
	// GetOptionsSize, it never mutates Used; pair it with
	// GetStringKeyValues (or a typed variant) to iterate every match.
	GetKeyValuesSize func(prefixes []string) int

	// GetStringOption finds the occurrence-th (0-based) argument that equals
	// one of the given flag spellings, then returns the argument immediately
	// following it as the option's value. It marks both the flag and its
	// value as used. It returns an error when occurrence is out of range for
	// the number of matches (see GetOptionsSize), or when the matched flag
	// is the last argument and has no following value to return.
	GetStringOption func(flags []string, occurrence int) (string, error)
	// GetIntOption behaves exactly like GetStringOption, additionally
	// parsing the option's value as a base-10 integer. It returns an error
	// on every failure GetStringOption can return, plus a parse error when
	// the value is not a valid integer.
	GetIntOption func(flags []string, occurrence int) (int, error)
	// GetDoubleOption behaves exactly like GetStringOption, additionally
	// parsing the option's value as a 64-bit floating-point number. It
	// returns an error on every failure GetStringOption can return, plus a
	// parse error when the value is not a valid number.
	GetDoubleOption func(flags []string, occurrence int) (float64, error)
	// GetTimestampOption behaves exactly like GetStringOption, additionally
	// parsing the option's value as an RFC 3339 timestamp (e.g.
	// "2024-01-02T15:04:05Z"). It returns an error on every failure
	// GetStringOption can return, plus a parse error when the value is not a
	// valid RFC 3339 timestamp.
	GetTimestampOption func(flags []string, occurrence int) (time.Time, error)

	// GetStringArg returns the argument at the given absolute index of Args
	// — the same index numbering as the raw command line (0 is the first
	// argument after the program name), independent of which arguments have
	// already been read. It marks that index as used. It returns an error
	// when index is negative or beyond the end of Args.
	GetStringArg func(index int) (string, error)
	// GetIntArg behaves exactly like GetStringArg, additionally parsing the
	// argument as a base-10 integer, returning a parse error when it is not
	// a valid integer.
	GetIntArg func(index int) (int, error)
	// GetDoubleArg behaves exactly like GetStringArg, additionally parsing
	// the argument as a 64-bit floating-point number, returning a parse
	// error when it is not a valid number.
	GetDoubleArg func(index int) (float64, error)
	// GetTimestampArg behaves exactly like GetStringArg, additionally
	// parsing the argument as an RFC 3339 timestamp, returning a parse error
	// when it is not a valid RFC 3339 timestamp.
	GetTimestampArg func(index int) (time.Time, error)

	// GetNextStringArg returns the first argument in Args, in order, whose
	// Used entry is still false, and marks it used. This is the core of the
	// Unused Mechanic: after every flag and option a program expects has
	// been read with the functions above, whatever remains unread is exactly
	// the leftover positional arguments — call this repeatedly to drain them
	// in order. It returns an error when every argument has already been
	// used.
	GetNextStringArg func() (string, error)
	// GetNextIntArg behaves exactly like GetNextStringArg, additionally
	// parsing the argument as a base-10 integer, returning a parse error
	// when it is not a valid integer.
	GetNextIntArg func() (int, error)
	// GetNextDoubleArg behaves exactly like GetNextStringArg, additionally
	// parsing the argument as a 64-bit floating-point number, returning a
	// parse error when it is not a valid number.
	GetNextDoubleArg func() (float64, error)
	// GetNextTimestampArg behaves exactly like GetNextStringArg,
	// additionally parsing the argument as an RFC 3339 timestamp, returning
	// a parse error when it is not a valid RFC 3339 timestamp.
	GetNextTimestampArg func() (time.Time, error)

	// GetStringKeyValues finds the occurrence-th (0-based) argument that
	// starts with one of the given key=value prefixes (the separator is
	// part of the prefix, e.g. "username="), then returns the text after
	// the matched prefix as the value. It marks that argument as used. It
	// returns an error when occurrence is out of range for the number of
	// matches (see GetKeyValuesSize), or when the matched argument's value
	// portion is empty (e.g. bare "username=" with nothing after it).
	GetStringKeyValues func(prefixes []string, occurrence int) (string, error)
	// GetIntKeyValues behaves exactly like GetStringKeyValues, additionally
	// parsing the value portion as a base-10 integer, returning a parse
	// error when it is not a valid integer.
	GetIntKeyValues func(prefixes []string, occurrence int) (int, error)
	// GetDoubleKeyValues behaves exactly like GetStringKeyValues,
	// additionally parsing the value portion as a 64-bit floating-point
	// number, returning a parse error when it is not a valid number.
	GetDoubleKeyValues func(prefixes []string, occurrence int) (float64, error)
	// GetTimestampKeyValues behaves exactly like GetStringKeyValues,
	// additionally parsing the value portion as an RFC 3339 timestamp,
	// returning a parse error when it is not a valid RFC 3339 timestamp.
	GetTimestampKeyValues func(prefixes []string, occurrence int) (time.Time, error)
}
