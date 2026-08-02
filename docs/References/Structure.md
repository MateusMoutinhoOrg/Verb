# Project Structure

This document maps the project **schema** — the kinds of files the project is built from — not every concrete file. A slot with a **Spec** name is governed by a specification; resolve the name through [Specs.md](/docs/References/Specs.md) to get its description and sample.

The project is split into two top-level trees, and the dependency flow between them is one-way:

```
sandbox/  ◀──  examples/
(closed)       (consumes the lib)
```

- **`/sandbox/`** is a **closed sandbox**: the pure library. Nothing inside it may import `examples/`, a third-party module, or any OS-bound standard-library package — every input it needs arrives as a plain function argument. See [SandboxIsolation.md](/docs/Explanations/SandboxIsolation.md).
- **`/examples/`** sits outside the sandbox and is the only place `os.Args` is read and handed to the library.

## Root

| File | Description | Spec |
|------|-------------|------|
| `README.md` | Project overview, quick start, Doc Index, and Samples section | Readme |
| `LICENSE` | License terms for the project | |
| `go.mod` | Go module definition and dependencies | |
| `.gitignore` | Intentionally untracked files to ignore | |

---

## `/sandbox/`
The closed sandbox — the pure library. It holds its own entry point, the contract everything is wired through, and the internal implementation. It reaches nothing outside itself. Its package is named `lib`, so consumers import it as `lib "…/sandbox"` and call `lib.New`.

| File | Description | Spec |
|------|-------------|------|
| `new.go` | The `New` constructor storing `args` on `api.Lib` and running the internal factories over it | |

### `/sandbox/contracts/`
The structs the rest of the project is wired through — the only part of the sandbox anything outside it may import. Contracts hold the project's **public types** and are structs of function fields, never interfaces; see [StructContracts.md](/docs/Explanations/StructContracts.md). Contracts import nothing from `sandbox/internal/`.

#### `/sandbox/contracts/api/`
The structs the library hands back to callers.

| File | Description | Spec |
|------|-------------|------|
| `api.go` | The `Lib` entry-point struct plus one struct per object the lib creates | Outputs |

### `/sandbox/internal/`
**Factories only** — no types. Each package here holds the functions that take a pointer to an [`api`](#sandboxcontractsapi) struct and return closures reading that struct's state, which the package's `New` constructor assigns into the matching function fields. Types never live here; they stay in `contracts/`. Go's `internal/` rule makes this tree unreachable from outside `sandbox/`.

#### `/sandbox/internal/lib/`
The entry-point implementation. The `internal/` parent already marks it private, so the package carries no `internal_` prefix.

| File | Description | Spec |
|------|-------------|------|
| `lib.go` | One `<Field>Factory(l *api.Lib)` per lib function, each returning a closure, plus the `New(args []string) api.Lib` constructor that assigns every factory's return value and runs them all | LibFunctions |

#### `/sandbox/internal/<object>/`
One package per object the library creates, named after the object itself.

| File | Description | Spec |
|------|-------------|------|
| `<object>.go` | The object's `<Field>Factory` functions, each returning a closure, plus the `New(…) api.<Object>` constructor that assigns every factory's return value | LibObjects |

---

## `/examples/`
Outside the sandbox. Runnable examples demonstrating how to use the library — the only place `os.Args` is read.

### `/examples/<example>/`

| File | Description | Spec |
|------|-------------|------|
| `<example>.go` | Self-contained `package main` calling `lib.New` with the real process argv | Examples |

**Run an example:**
```sh
go run ./examples/<example>/<example>.go
```

---

## `/docs/`
Documentation of the project.

### `/docs/References/`
Listable material — structures, rules, specifications, and the public API index.

| File | Description | Spec |
|------|-------------|------|
| `RULES.md` | Rules to follow when contributing to this project | Rules |
| `Structure.md` | The project's schema and the purpose of each component | Structure |
| `Specs.md` | Index of every specification and the files each one governs | |
| `PublicApi.md` | Index of all public-facing components, linking to their detail pages | ReferenceDocs |
| `TemplateFileActions.md` | The action each template file takes when forking or adapting a library | ReferenceDocs |
| `<Name>.md` | Any other reference page the library needs | ReferenceDocs |

#### `/docs/References/Meta/`
The specifications describing how each kind of file in the project must be shaped. Never browse this directory — locate a specification by reading `Specs.md`.

| File | Description | Spec |
|------|-------------|------|
| `<Spec>/Specs.md` | The required shape of the artifact the specification governs | |
| `<Spec>/sample.<ext>` | Concrete reference implementation of the specification | |

#### `/docs/References/PublicApi/`
Detailed documentation for each individual public API entry.

| File | Description | Spec |
|------|-------------|------|
| `<pkg>.<Symbol>.md` | One detail page per public struct, function, or field | ReferenceDocs |

### `/docs/Explanations/`
Explanations of the project's mechanics and features.

| File | Description | Spec |
|------|-------------|------|
| `<Topic>.md` | One page per mechanic the library needs explained | ExplanationDocs |

### `/docs/Tutorials/`
Workflow guides explaining how to use, extend, and maintain the project. Each file covers a single goal.

| File | Description | Spec |
|------|-------------|------|
| `<Goal>.md` | One page per workflow the library's maintainers repeat | TutorialDocs |
