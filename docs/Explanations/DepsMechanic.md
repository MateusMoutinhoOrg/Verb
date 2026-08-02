# Dependency Mechanics

## Description
Explains how the library receives its dependencies. `Deps` is a **struct of function fields**: the library declares the behaviors it needs as fields, and the caller supplies a value with every field filled — a pre-built adapter, an adapter with one field replaced, or a struct literal written by hand.

---

## The Contract

`sandbox/contracts/deps` declares what the library needs; `sandbox/contracts/api` declares what it hands back. Nothing else crosses the boundary.

```go
// sandbox/contracts/deps/deps.go — what the library needs
type Deps struct {
	Args func() []string
}

// sandbox/contracts/api/api.go — what the library returns
type Lib struct {
	Deps      deps.Deps
	Args      []string
	Used      []bool
	IsPresent func(flags []string) bool
	// ...GetStringOption, GetStringArg, GetNextStringArg, GetStringKeyValues,
	// and their Int/Double/Timestamp variants
}
```

`lib.New(deps.Deps) api.Lib` is the single wiring point. Because both sides are plain structs, `sandbox` never imports an adapter, and callers never import `sandbox/internal/`. For why the contracts are structs rather than interfaces, see [StructContracts.md](/docs/Explanations/StructContracts.md).

`Deps` is the *only* door in the sandbox wall: since nothing under `sandbox/` may import an adapter, a third-party module, or an OS-bound standard-library package, every effect the library performs has to be a field on this struct. An argv parser needs exactly one such effect — reading the raw argument vector — so `Deps` has exactly one field. That constraint is what this page's mechanic exists to serve — see [SandboxIsolation.md](/docs/Explanations/SandboxIsolation.md).

---

## Using an Adapter

The simplest setup: an adapter builds a ready-to-use `deps.Deps`.

```go
package main

import (
	"os"

	verbadapter "github.com/MateusMoutinhoOrg/Verb/adapters/standard"
	verblib "github.com/MateusMoutinhoOrg/Verb/sandbox"
)

func main() {
	// The standard adapter fills deps.Deps.Args with the real process argv
	myDeps := verbadapter.New(os.Args[1:])
	l := verblib.New(myDeps)

	if l.IsPresent([]string{"-q", "--quiet"}) {
		println("quiet mode")
	}
}
```

---

## Overwriting a Single Behavior

To keep an adapter but change one behavior, take the `deps.Deps` it returns and assign the field you want to replace. With a single-field contract this mostly means replacing `Args` wholesale — e.g. testing against a fixed argv instead of the real one, without switching adapters at all.

```go
package main

import (
	"os"

	verbadapter "github.com/MateusMoutinhoOrg/Verb/adapters/standard"
	verblib "github.com/MateusMoutinhoOrg/Verb/sandbox"
)

func main() {
	// 1. Get the default implementation from an adapter
	myDeps := verbadapter.New(os.Args[1:])

	// 2. Replace Args with a fixed slice for a dry run, ignoring the real argv
	myDeps.Args = func() []string { return []string{"--quiet", "input.txt"} }

	// 3. Inject — the lib sees a normal deps.Deps
	l := verblib.New(myDeps)
	println(l.IsPresent([]string{"-q", "--quiet"})) // true
}
```

> **Careful:** patch the `deps.Deps` value **before** calling `lib.New`. `lib.New` calls `Deps.Args()` exactly once to take its snapshot, so replacing `Args` on the `api.Lib` returned by `verblib.New` has no effect — the snapshot is already taken.

---

## Writing Custom Deps

For complete control, build the `deps.Deps` yourself. No adapter is involved: it is a struct literal, so there is no type to declare and no method set to satisfy.

```go
package main

import (
	verblib "github.com/MateusMoutinhoOrg/Verb/sandbox"
	verbdeps "github.com/MateusMoutinhoOrg/Verb/sandbox/contracts/deps"
)

func main() {
	// 1. Build your own implementation, hard-coding the argv to parse
	myDeps := verbdeps.Deps{
		Args: func() []string { return []string{"-o", "out.txt", "in.txt"} },
	}

	// 2. Inject it into the library
	l := verblib.New(myDeps)

	// 3. Use the library normally
	output, err := l.GetStringOption([]string{"-o", "--output"}, 0)
	if err != nil {
		panic(err)
	}
	println(output)
}
```

> **Careful:** the compiler cannot tell you a field is missing. A `Deps` built by hand with an unfilled field holds a nil function that panics on first call — fill every field.

---

## Propagation

`lib.New` delegates to the constructor in `sandbox/internal/lib/`, which snapshots `Deps.Args()`, stores the `Deps` and the snapshot on the `api.Lib` struct itself, and then runs the factories over it:

```go
// sandbox/new.go
func New(d deps.Deps) api.Lib {
	return internallib.New(d)
}

// sandbox/internal/lib/lib.go
func New(d deps.Deps) api.Lib {
	args := d.Args()
	l := api.Lib{Deps: d, Args: args, Used: make([]bool, len(args))}
	l.IsPresent = IsPresentFactory(&l)
	l.GetStringOption = GetStringOptionFactory(&l)
	// ...every other field factory
	return l
}
```

The carrier is the **closure**. Each factory returns a closure that reads and mutates `l.Args`/`l.Used` when the field is called, and `New` assigns it:

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
```

Because every factory closes over the **same** `*api.Lib` that `New` built, a call to `IsPresent` marking an index used is visible to every later call to `GetNextStringArg` on the same `Lib` — that shared, mutable `Used` slice is what makes the Unused Mechanic work. See [UnnusedMechanic.md](/docs/Explanations/UnnusedMechanic.md).

```
standard.New() ──▶ deps.Deps ──▶ lib.New(deps) ──▶ api.Lib
                                                     │
                                          snapshots Deps.Args() once
                                                     ▼
                                       Args []string, Used []bool
```

To add a new behavior to the contract, follow [AddDependency.md](/docs/Tutorials/AddDependency.md).
