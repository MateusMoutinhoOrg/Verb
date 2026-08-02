# `api.Lib`

**Type:** Struct

## Definition

```go
type Lib struct {
	Deps deps.Deps
	Args []string
	Used []bool

	IsPresent      func(flags []string) bool
	GetOptionsSize func(flags []string) int
	GetKeyValuesSize func(prefixes []string) int

	GetStringOption    func(flags []string, occurrence int) (string, error)
	GetIntOption       func(flags []string, occurrence int) (int, error)
	GetDoubleOption    func(flags []string, occurrence int) (float64, error)
	GetTimestampOption func(flags []string, occurrence int) (time.Time, error)

	GetStringArg    func(index int) (string, error)
	GetIntArg       func(index int) (int, error)
	GetDoubleArg    func(index int) (float64, error)
	GetTimestampArg func(index int) (time.Time, error)

	GetNextStringArg    func() (string, error)
	GetNextIntArg       func() (int, error)
	GetNextDoubleArg    func() (float64, error)
	GetNextTimestampArg func() (time.Time, error)

	GetStringKeyValues    func(prefixes []string, occurrence int) (string, error)
	GetIntKeyValues       func(prefixes []string, occurrence int) (int, error)
	GetDoubleKeyValues    func(prefixes []string, occurrence int) (float64, error)
	GetTimestampKeyValues func(prefixes []string, occurrence int) (time.Time, error)
}
```

## Description

The library entry point, returned by [`lib.New`](/docs/References/PublicApi/lib.New.md). It is an argv parser exposed as a struct of function fields: `lib.New` stores the injected [`deps.Deps`](/docs/References/PublicApi/deps.Deps.md) in `Deps`, snapshots the raw argument vector into `Args`, allocates the matching `Used` tracking slice, then runs the factories in `sandbox/internal/lib/`, each of which fills one function field with a closure reading `Args`/`Used` at call time. Calling a field reads exactly like calling a method — `l.IsPresent([]string{"-q"})`. See [StructContracts.md](/docs/Explanations/StructContracts.md).

Every `Get*`/`IsPresent` call marks the argument(s) it matched as read in `Used`; `GetNextStringArg` and its typed variants return the first argument still unread. This is the Unused Mechanic — see [UnnusedMechanic.md](/docs/Explanations/UnnusedMechanic.md).

`Deps` is exported because the library's own factories read it, but it is **read-only after construction**: the closures already captured the struct they were built over, so reassigning `Deps` here does not change behavior. Patch the `deps.Deps` value before calling `lib.New`.

## Fields

| Field | Description |
| :--- | :--- |
| [`Deps deps.Deps`](/docs/References/PublicApi/deps.Deps.md) | The dependency set injected by `lib.New`; read-only after construction. |
| [`Args []string`](/docs/References/PublicApi/api.Args.md) | Snapshot of the argument vector being parsed. |
| [`Used []bool`](/docs/References/PublicApi/api.Used.md) | Tracks which positions of `Args` have already been read. |
| [`IsPresent func(flags []string) bool`](/docs/References/PublicApi/api.IsPresent.md) | Reports whether any of the given flag spellings is present. |
| [`GetOptionsSize func(flags []string) int`](/docs/References/PublicApi/api.Size.md) | Counts how many arguments match one of the given flags. |
| [`GetKeyValuesSize func(prefixes []string) int`](/docs/References/PublicApi/api.Size.md) | Counts how many arguments start with one of the given `key=` prefixes. |
| [`GetStringOption` / `GetIntOption` / `GetDoubleOption` / `GetTimestampOption`](/docs/References/PublicApi/api.GetOption.md) | Read a flag's following value, typed. |
| [`GetStringArg` / `GetIntArg` / `GetDoubleArg` / `GetTimestampArg`](/docs/References/PublicApi/api.GetArg.md) | Read the argument at an absolute position, typed. |
| [`GetNextStringArg` / `GetNextIntArg` / `GetNextDoubleArg` / `GetNextTimestampArg`](/docs/References/PublicApi/api.GetNextArg.md) | Read the next not-yet-used argument, typed. |
| [`GetStringKeyValues` / `GetIntKeyValues` / `GetDoubleKeyValues` / `GetTimestampKeyValues`](/docs/References/PublicApi/api.GetKeyValues.md) | Read a `key=value` argument's value, typed. |
