# Contribution Rules

Rules to follow when contributing to this project. Every file must also be shaped by the specification that governs it — locate it in [Specs.md](/docs/References/Specs.md).

---

## Tutorials Guide
Before making anything, read the [README.md](/README.md) and search for a tutorial about what you want to do. If there is one, follow it; if there isn't, you need to create one following the spec defined in [TutorialDocs](./Meta/TutorialDocs/).


## Specification Compliance

Before creating or editing any file, read [Specs.md](/docs/References/Specs.md) and check whether the file matches an **Applies To** entry. If it does, create or edit it following the specification that entry points to — reproduce the shape it requires, using its `sample` as reference.

---

## Sandbox Isolation

[sandbox/](/sandbox/) is a closed sandbox. No file inside it may import [examples/](/examples/), a third-party module, or an OS-bound standard-library package (`os`, `net`, `os/exec`, `syscall`, …). Every input the library needs from the outside world arrives as a plain function argument, e.g. `lib.New(args []string)`. The mechanic is explained in [SandboxIsolation.md](/docs/Explanations/SandboxIsolation.md).

The `api` contract is a **struct of function fields**, never an interface — see [sandbox/contracts/api](/sandbox/contracts/api/). Every type in the project is declared there; [sandbox/internal/](/sandbox/internal/) declares no types at all. See [StructContracts.md](/docs/Explanations/StructContracts.md).

---

## Factory Pattern

Every struct of function fields in this project — an `api` struct inside the sandbox — is filled by **factories**, never by methods bound into fields and never by an internal mirror type. When you write or edit any file holding `<Field>Factory` functions, follow the [Factories](./Meta/Factories/Specs.md) specification on top of the one governing that file's tree.

A factory takes a pointer to the **carrier** — the api struct holding the state the closure reads — and returns exactly one field's value; the caller assigns it:

```go
// sandbox/internal/lib/ — the carrier is the api struct being filled
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

func New(args []string) api.Lib {
	l := api.Lib{Args: args, Used: make([]bool, len(args))}
	l.IsPresent = IsPresentFactory(&l)
	return l
}
```

Three rules follow, and none of them is checked by the compiler:

- Every api struct's closures reach state through the carrier pointer — `l.Args`, `l.Used` — read inside the closure, never captured at factory time.
- Every field factory must be called and its return value assigned from its package's `New(...)` constructor, which is the factory aggregate — there is no separate `Factory` function. A field no factory fills stays nil and panics on first call.
- A `New` constructor returns the filled **contract struct** by value — `api.Lib` or `api.<Object>` — never a private carrier type.

Nothing outside the sandbox may reach into it beyond its two public packages: `sandbox` (package `lib`) and `sandbox/contracts/api`.

---

## Import Aliases

Any file that **consumes** the library from outside it — [examples/](/examples/) and third-party consumers — imports it under `verb`-prefixed aliases, so each call site says which layer it belongs to:

| Import | Alias |
|--------|-------|
| `sandbox` | `verblib` |
| `sandbox/contracts/api` | `verbtypes` |

```go
import (
	verblib "github.com/MateusMoutinhoOrg/Verb/sandbox"
	verbtypes "github.com/MateusMoutinhoOrg/Verb/sandbox/contracts/api"
)
```

Files that belong to the library — everything under `sandbox/` — keep the plain package names (`api`, `lib`): there the prefix would be noise, since the import is already local.

---

## File Changes

Before creating, deleting, or renaming any file or directory, read [Structure.md](/docs/References/Structure.md) and check whether the change affects the project structure. If it does, update [Structure.md](/docs/References/Structure.md) in the same commit.

---

## Specification Changes

When you create, delete, or rename a specification inside [Meta/](./Meta), you MUST adapt all the files that match the spec's Applies To rule, and update the index in [Specs.md](/docs/References/Specs.md).

---

## Documentation Changes

When you create, delete, or rename a `.md` file, update the Doc Index of [README.md](/README.md).

---

## Sample Changes

When you create, delete, or rename a sample (any file inside [examples/](/examples/)), update the Samples section of [README.md](/README.md).
