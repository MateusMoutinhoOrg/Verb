# Public API

## Description
Index of every public-facing entry of the library. Callers hold **structs of function fields** declared in `sandbox/contracts/api` and `sandbox/contracts/deps`; the **factories** that fill those fields live in `sandbox/internal/` and are unreachable from outside `sandbox/`. See [StructContracts.md](/docs/Explanations/StructContracts.md).

---

## Structs

### [api.Lib](/docs/References/PublicApi/api.Lib.md)
The library entry point — an argv parser with a flag/option/positional-argument mechanic. Returned by `lib.New`; exposes all library functions as fields.

### [deps.Deps](/docs/References/PublicApi/deps.Deps.md)
The dependency contract every adapter must fill: a single function returning the argument vector to parse.

---

## Functions

### [lib.New](/docs/References/PublicApi/lib.New.md)
Injects a `deps.Deps` into the library and returns an `api.Lib`.

### [standard.New](/docs/References/PublicApi/standard.New.md)
Creates a `deps.Deps` using the standard library adapter, wrapping the real process argv.

### [memory.New](/docs/References/PublicApi/memory.New.md)
Creates a `deps.Deps` using the memory adapter, wrapping a fixed argv for tests.

---

## Fields

### [api.Lib.Deps](/docs/References/PublicApi/api.Deps.md)
The injected dependency set the struct was built with; read-only after construction.

### [api.Lib.Args](/docs/References/PublicApi/api.Args.md)
Snapshot of the argument vector being parsed.

### [api.Lib.Used](/docs/References/PublicApi/api.Used.md)
Tracks which positions of `Args` have already been read — the Unused Mechanic's bookkeeping.

### [api.Lib.IsPresent](/docs/References/PublicApi/api.IsPresent.md)
Reports whether any spelling of a flag is present, marking it used on a hit.

### [api.Lib.GetOptionsSize](/docs/References/PublicApi/api.Size.md)
Counts how many arguments match one of the given flags, without mutating `Used`.

### [api.Lib.GetKeyValuesSize](/docs/References/PublicApi/api.Size.md)
Counts how many arguments start with one of the given `key=` prefixes, without mutating `Used`.

### [api.Lib.GetStringOption](/docs/References/PublicApi/api.GetOption.md)
Reads the value following a matched flag as raw text.

### [api.Lib.GetIntOption](/docs/References/PublicApi/api.GetOption.md)
Reads the value following a matched flag as a base-10 integer.

### [api.Lib.GetDoubleOption](/docs/References/PublicApi/api.GetOption.md)
Reads the value following a matched flag as a 64-bit float.

### [api.Lib.GetTimestampOption](/docs/References/PublicApi/api.GetOption.md)
Reads the value following a matched flag as an RFC 3339 timestamp.

### [api.Lib.GetStringArg](/docs/References/PublicApi/api.GetArg.md)
Reads the argument at an absolute index as raw text.

### [api.Lib.GetIntArg](/docs/References/PublicApi/api.GetArg.md)
Reads the argument at an absolute index as a base-10 integer.

### [api.Lib.GetDoubleArg](/docs/References/PublicApi/api.GetArg.md)
Reads the argument at an absolute index as a 64-bit float.

### [api.Lib.GetTimestampArg](/docs/References/PublicApi/api.GetArg.md)
Reads the argument at an absolute index as an RFC 3339 timestamp.

### [api.Lib.GetNextStringArg](/docs/References/PublicApi/api.GetNextArg.md)
Reads the next not-yet-used argument as raw text.

### [api.Lib.GetNextIntArg](/docs/References/PublicApi/api.GetNextArg.md)
Reads the next not-yet-used argument as a base-10 integer.

### [api.Lib.GetNextDoubleArg](/docs/References/PublicApi/api.GetNextArg.md)
Reads the next not-yet-used argument as a 64-bit float.

### [api.Lib.GetNextTimestampArg](/docs/References/PublicApi/api.GetNextArg.md)
Reads the next not-yet-used argument as an RFC 3339 timestamp.

### [api.Lib.GetStringKeyValues](/docs/References/PublicApi/api.GetKeyValues.md)
Reads a matched `key=value` argument's value as raw text.

### [api.Lib.GetIntKeyValues](/docs/References/PublicApi/api.GetKeyValues.md)
Reads a matched `key=value` argument's value as a base-10 integer.

### [api.Lib.GetDoubleKeyValues](/docs/References/PublicApi/api.GetKeyValues.md)
Reads a matched `key=value` argument's value as a 64-bit float.

### [api.Lib.GetTimestampKeyValues](/docs/References/PublicApi/api.GetKeyValues.md)
Reads a matched `key=value` argument's value as an RFC 3339 timestamp.
