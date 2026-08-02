package lib

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/MateusMoutinhoOrg/Verb/sandbox/contracts/api"
)

// timestampLayout is the format every Timestamp getter parses values with.
const timestampLayout = time.RFC3339

// matchesFlag reports whether arg equals one of the given flag spellings.
func matchesFlag(arg string, flags []string) bool {
	for _, f := range flags {
		if arg == f {
			return true
		}
	}
	return false
}

// matchPrefix reports whether arg starts with one of the given key=value
// prefixes, returning the matched prefix.
func matchPrefix(arg string, prefixes []string) (string, bool) {
	for _, p := range prefixes {
		if strings.HasPrefix(arg, p) {
			return p, true
		}
	}
	return "", false
}

// findOptionIndex returns the index in l.Args of the occurrence-th argument
// matching one of flags, regardless of Used.
func findOptionIndex(l *api.Lib, flags []string, occurrence int) (int, error) {
	count := 0
	for i, a := range l.Args {
		if matchesFlag(a, flags) {
			if count == occurrence {
				return i, nil
			}
			count++
		}
	}
	return -1, fmt.Errorf("verb: option %v: occurrence %d not found (only %d present)", flags, occurrence, count)
}

// stringOptionValue implements GetStringOption: it locates the occurrence-th
// flag, marks it and its following value as used, and returns that value.
func stringOptionValue(l *api.Lib, flags []string, occurrence int) (string, error) {
	idx, err := findOptionIndex(l, flags, occurrence)
	if err != nil {
		return "", err
	}
	if idx+1 >= len(l.Args) {
		return "", fmt.Errorf("verb: option %v: flag %q has no following value", flags, l.Args[idx])
	}
	l.Used[idx] = true
	l.Used[idx+1] = true
	return l.Args[idx+1], nil
}

// stringArgValue implements GetStringArg: it validates index against l.Args,
// marks it used, and returns the argument at that absolute position.
func stringArgValue(l *api.Lib, index int) (string, error) {
	if index < 0 || index >= len(l.Args) {
		return "", fmt.Errorf("verb: arg index %d out of range (have %d arguments)", index, len(l.Args))
	}
	l.Used[index] = true
	return l.Args[index], nil
}

// nextStringArgValue implements GetNextStringArg: it finds the first
// not-yet-used argument in order, marks it used, and returns it.
func nextStringArgValue(l *api.Lib) (string, error) {
	for i, used := range l.Used {
		if !used {
			l.Used[i] = true
			return l.Args[i], nil
		}
	}
	return "", fmt.Errorf("verb: no unused arguments remaining")
}

// stringKeyValuesValue implements GetStringKeyValues: it locates the
// occurrence-th argument starting with one of prefixes, marks it used, and
// returns the text after the matched prefix.
func stringKeyValuesValue(l *api.Lib, prefixes []string, occurrence int) (string, error) {
	count := 0
	for i, a := range l.Args {
		prefix, ok := matchPrefix(a, prefixes)
		if !ok {
			continue
		}
		if count == occurrence {
			l.Used[i] = true
			value := a[len(prefix):]
			if value == "" {
				return "", fmt.Errorf("verb: key/value %v: occurrence %d has an empty value", prefixes, occurrence)
			}
			return value, nil
		}
		count++
	}
	return "", fmt.Errorf("verb: key/value %v: occurrence %d not found (only %d present)", prefixes, occurrence, count)
}

// parseInt parses s as a base-10 integer for the Int-typed getters.
func parseInt(s string) (int, error) {
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("verb: %q is not a valid integer: %w", s, err)
	}
	return v, nil
}

// parseDouble parses s as a 64-bit float for the Double-typed getters.
func parseDouble(s string) (float64, error) {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, fmt.Errorf("verb: %q is not a valid number: %w", s, err)
	}
	return v, nil
}

// parseTimestamp parses s as an RFC 3339 timestamp for the Timestamp-typed
// getters.
func parseTimestamp(s string) (time.Time, error) {
	v, err := time.Parse(timestampLayout, s)
	if err != nil {
		return time.Time{}, fmt.Errorf("verb: %q is not a valid RFC3339 timestamp: %w", s, err)
	}
	return v, nil
}

// IsPresentFactory fills api.Lib.IsPresent, matching any unused argument
// against flags and marking it used on a hit.
func IsPresentFactory(l *api.Lib) func(flags []string) bool {
	return func(flags []string) bool {
		for i, a := range l.Args {
			if l.Used[i] {
				continue
			}
			if matchesFlag(a, flags) {
				l.Used[i] = true
				return true
			}
		}
		return false
	}
}

// GetOptionsSizeFactory fills api.Lib.GetOptionsSize, counting arguments
// matching flags without mutating Used.
func GetOptionsSizeFactory(l *api.Lib) func(flags []string) int {
	return func(flags []string) int {
		count := 0
		for _, a := range l.Args {
			if matchesFlag(a, flags) {
				count++
			}
		}
		return count
	}
}

// GetKeyValuesSizeFactory fills api.Lib.GetKeyValuesSize, counting arguments
// starting with one of prefixes without mutating Used.
func GetKeyValuesSizeFactory(l *api.Lib) func(prefixes []string) int {
	return func(prefixes []string) int {
		count := 0
		for _, a := range l.Args {
			if _, ok := matchPrefix(a, prefixes); ok {
				count++
			}
		}
		return count
	}
}

// GetStringOptionFactory fills api.Lib.GetStringOption.
func GetStringOptionFactory(l *api.Lib) func(flags []string, occurrence int) (string, error) {
	return func(flags []string, occurrence int) (string, error) {
		return stringOptionValue(l, flags, occurrence)
	}
}

// GetIntOptionFactory fills api.Lib.GetIntOption, parsing the option value
// found by stringOptionValue as an integer.
func GetIntOptionFactory(l *api.Lib) func(flags []string, occurrence int) (int, error) {
	return func(flags []string, occurrence int) (int, error) {
		s, err := stringOptionValue(l, flags, occurrence)
		if err != nil {
			return 0, err
		}
		return parseInt(s)
	}
}

// GetDoubleOptionFactory fills api.Lib.GetDoubleOption, parsing the option
// value found by stringOptionValue as a float64.
func GetDoubleOptionFactory(l *api.Lib) func(flags []string, occurrence int) (float64, error) {
	return func(flags []string, occurrence int) (float64, error) {
		s, err := stringOptionValue(l, flags, occurrence)
		if err != nil {
			return 0, err
		}
		return parseDouble(s)
	}
}

// GetTimestampOptionFactory fills api.Lib.GetTimestampOption, parsing the
// option value found by stringOptionValue as an RFC 3339 timestamp.
func GetTimestampOptionFactory(l *api.Lib) func(flags []string, occurrence int) (time.Time, error) {
	return func(flags []string, occurrence int) (time.Time, error) {
		s, err := stringOptionValue(l, flags, occurrence)
		if err != nil {
			return time.Time{}, err
		}
		return parseTimestamp(s)
	}
}

// GetStringArgFactory fills api.Lib.GetStringArg.
func GetStringArgFactory(l *api.Lib) func(index int) (string, error) {
	return func(index int) (string, error) {
		return stringArgValue(l, index)
	}
}

// GetIntArgFactory fills api.Lib.GetIntArg, parsing the argument found by
// stringArgValue as an integer.
func GetIntArgFactory(l *api.Lib) func(index int) (int, error) {
	return func(index int) (int, error) {
		s, err := stringArgValue(l, index)
		if err != nil {
			return 0, err
		}
		return parseInt(s)
	}
}

// GetDoubleArgFactory fills api.Lib.GetDoubleArg, parsing the argument found
// by stringArgValue as a float64.
func GetDoubleArgFactory(l *api.Lib) func(index int) (float64, error) {
	return func(index int) (float64, error) {
		s, err := stringArgValue(l, index)
		if err != nil {
			return 0, err
		}
		return parseDouble(s)
	}
}

// GetTimestampArgFactory fills api.Lib.GetTimestampArg, parsing the argument
// found by stringArgValue as an RFC 3339 timestamp.
func GetTimestampArgFactory(l *api.Lib) func(index int) (time.Time, error) {
	return func(index int) (time.Time, error) {
		s, err := stringArgValue(l, index)
		if err != nil {
			return time.Time{}, err
		}
		return parseTimestamp(s)
	}
}

// GetNextStringArgFactory fills api.Lib.GetNextStringArg.
func GetNextStringArgFactory(l *api.Lib) func() (string, error) {
	return func() (string, error) {
		return nextStringArgValue(l)
	}
}

// GetNextIntArgFactory fills api.Lib.GetNextIntArg, parsing the argument
// found by nextStringArgValue as an integer.
func GetNextIntArgFactory(l *api.Lib) func() (int, error) {
	return func() (int, error) {
		s, err := nextStringArgValue(l)
		if err != nil {
			return 0, err
		}
		return parseInt(s)
	}
}

// GetNextDoubleArgFactory fills api.Lib.GetNextDoubleArg, parsing the
// argument found by nextStringArgValue as a float64.
func GetNextDoubleArgFactory(l *api.Lib) func() (float64, error) {
	return func() (float64, error) {
		s, err := nextStringArgValue(l)
		if err != nil {
			return 0, err
		}
		return parseDouble(s)
	}
}

// GetNextTimestampArgFactory fills api.Lib.GetNextTimestampArg, parsing the
// argument found by nextStringArgValue as an RFC 3339 timestamp.
func GetNextTimestampArgFactory(l *api.Lib) func() (time.Time, error) {
	return func() (time.Time, error) {
		s, err := nextStringArgValue(l)
		if err != nil {
			return time.Time{}, err
		}
		return parseTimestamp(s)
	}
}

// GetStringKeyValuesFactory fills api.Lib.GetStringKeyValues.
func GetStringKeyValuesFactory(l *api.Lib) func(prefixes []string, occurrence int) (string, error) {
	return func(prefixes []string, occurrence int) (string, error) {
		return stringKeyValuesValue(l, prefixes, occurrence)
	}
}

// GetIntKeyValuesFactory fills api.Lib.GetIntKeyValues, parsing the value
// found by stringKeyValuesValue as an integer.
func GetIntKeyValuesFactory(l *api.Lib) func(prefixes []string, occurrence int) (int, error) {
	return func(prefixes []string, occurrence int) (int, error) {
		s, err := stringKeyValuesValue(l, prefixes, occurrence)
		if err != nil {
			return 0, err
		}
		return parseInt(s)
	}
}

// GetDoubleKeyValuesFactory fills api.Lib.GetDoubleKeyValues, parsing the
// value found by stringKeyValuesValue as a float64.
func GetDoubleKeyValuesFactory(l *api.Lib) func(prefixes []string, occurrence int) (float64, error) {
	return func(prefixes []string, occurrence int) (float64, error) {
		s, err := stringKeyValuesValue(l, prefixes, occurrence)
		if err != nil {
			return 0, err
		}
		return parseDouble(s)
	}
}

// GetTimestampKeyValuesFactory fills api.Lib.GetTimestampKeyValues, parsing
// the value found by stringKeyValuesValue as an RFC 3339 timestamp.
func GetTimestampKeyValuesFactory(l *api.Lib) func(prefixes []string, occurrence int) (time.Time, error) {
	return func(prefixes []string, occurrence int) (time.Time, error) {
		s, err := stringKeyValuesValue(l, prefixes, occurrence)
		if err != nil {
			return time.Time{}, err
		}
		return parseTimestamp(s)
	}
}

// New builds the api.Lib entry point: it stores args on Args, allocates a
// same-length Used tracking slice, and runs every lib factory over it to
// fill its function fields. Adding a function field to api.Lib means adding
// its factory call here.
func New(args []string) api.Lib {
	l := api.Lib{
		Args: args,
		Used: make([]bool, len(args)),
	}

	l.IsPresent = IsPresentFactory(&l)
	l.GetOptionsSize = GetOptionsSizeFactory(&l)
	l.GetKeyValuesSize = GetKeyValuesSizeFactory(&l)

	l.GetStringOption = GetStringOptionFactory(&l)
	l.GetIntOption = GetIntOptionFactory(&l)
	l.GetDoubleOption = GetDoubleOptionFactory(&l)
	l.GetTimestampOption = GetTimestampOptionFactory(&l)

	l.GetStringArg = GetStringArgFactory(&l)
	l.GetIntArg = GetIntArgFactory(&l)
	l.GetDoubleArg = GetDoubleArgFactory(&l)
	l.GetTimestampArg = GetTimestampArgFactory(&l)

	l.GetNextStringArg = GetNextStringArgFactory(&l)
	l.GetNextIntArg = GetNextIntArgFactory(&l)
	l.GetNextDoubleArg = GetNextDoubleArgFactory(&l)
	l.GetNextTimestampArg = GetNextTimestampArgFactory(&l)

	l.GetStringKeyValues = GetStringKeyValuesFactory(&l)
	l.GetIntKeyValues = GetIntKeyValuesFactory(&l)
	l.GetDoubleKeyValues = GetDoubleKeyValuesFactory(&l)
	l.GetTimestampKeyValues = GetTimestampKeyValuesFactory(&l)

	return l
}
