# Struct Contracts

## Description
Explains why the library's public contract — everything in `sandbox/contracts/api` — is a **struct of function fields** instead of an interface, what that buys, and what it costs.

---

## The Shape

A contract is a struct whose fields are functions. The library declares the shape; its own factories decide the behavior.

```go
// sandbox/contracts/api/api.go — what the library hands back
type Lib struct {
	Args      []string
	Used      []bool
	IsPresent func(flags []string) bool
	// ...GetStringOption, GetStringArg, GetNextStringArg, GetStringKeyValues
}
```

Callers use it exactly as they would an interface — `l.IsPresent([]string{"-q"})` reads the same whether `IsPresent` is a method or a field holding a function. The difference only shows up at the wiring point.

---

## Factories Fill the Fields

There is no internal type and no method set. `sandbox/internal/` holds **factories**: functions that take a pointer to the api struct and return a closure for one of its function fields; the caller assigns the result.

```go
// sandbox/internal/lib/lib.go
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

// New is the package's constructor and its factory aggregate in one — the
// one place that must stay complete.
func New(args []string) api.Lib {
	l := api.Lib{Args: args, Used: make([]bool, len(args))}
	l.IsPresent = IsPresentFactory(&l)
	l.GetStringOption = GetStringOptionFactory(&l)
	// ...every other field factory
	return l
}
```

```go
// sandbox/new.go — the whole wiring
func New(args []string) api.Lib {
	return internallib.New(args)
}
```

The closure captures `l`, so `l.Args`/`l.Used` are read and mutated when `IsPresent` is *called*, not when the factory ran. That is what carries the parsed state into behavior the caller can hold, while `sandbox/internal/` stays unreachable from outside.

Two properties follow, and both are load-bearing:

- **One field, one factory.** A factory does nothing but return the closure; the package's `New` constructor is the only place that has to assign them all.
- **State is read through the pointer, never copied into the closure.** Capturing `args`/`Used` by value at factory time would freeze them at construction; reading `l.Args`/`l.Used` keeps the struct authoritative.

Binding a method into a field would work in Go, but the project does not do it: one shape for filling struct contracts means one place to look for completeness — the `New` at the bottom of the file. The rule is binding; see [RULES.md](/docs/References/RULES.md#factory-pattern) and the [Factories](/docs/References/Meta/Factories/Specs.md) specification.

---

## Consuming a Library That Uses This Pattern

The payoff shows when a caller wants to hold this library's whole entry point as a single field of its own contract: because `api.Lib` is a struct of function fields rather than an interface, a caller can copy its shape into a struct it declares itself and assign the real fields straight across, with no bridging type needed:

```go
type Deps struct {
	IsPresent        func(flags []string) bool
	GetNextStringArg func() (string, error)
}

parser := verblib.New(os.Args[1:])
d := Deps{
	IsPresent:        parser.IsPresent,        // identical signature: assigned straight across
	GetNextStringArg: parser.GetNextStringArg, // identical signature: assigned straight across
}
```

With interfaces, the same wiring needs a bridging type declaring every method, even the ones that pass straight through.

---

## What It Costs

The compiler no longer checks completeness. An interface implementation that misses a method fails to build; a struct contract with an unfilled field compiles fine and panics on the first call with a nil-pointer dereference.

That moves one guarantee from the compiler to the author: a package's `New` constructor must call **every** field factory of its api struct, or the field it skips stays nil and panics on first call.

In exchange, there is no partial-implementation ambiguity at the call site: a filled contract is a value that can be copied and passed on.
