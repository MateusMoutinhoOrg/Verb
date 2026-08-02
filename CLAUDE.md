# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this is

Verb is a **Go library template** demonstrating a struct-of-functions public API instead of interfaces. The value is in the *structure and documentation conventions*, not the (deliberately trivial) example code — an OS-independent argv parser. Most "features" here are `Example*`-prefixed placeholders meant to be replaced when the template is adapted to a real library.

## Commands

```bash
go build ./...                                   # build everything
go run ./examples/Presence/Presence.go           # run an example
go test ./...                                    # run tests (none exist yet)
```

There is no lint config, CI, or test suite in the repo yet. When adding lib functions, `go build ./...` is the primary verification step.

## Architecture

Two top-level trees, wired through **structs of function fields** (never interfaces), with a strict one-way dependency flow:

```
sandbox/  ◀──  examples/
(closed)       (consumes the lib)

lib.New(args)  ──▶  api.Lib  ──▶  api.Entry
(entry point)       (output structs, filled by sandbox/internal/ factories)
```

Contracts are structs whose fields hold functions, and **every** one of them is filled by **factories** — `func <Field>Factory(carrier *T) <FieldType>` bodies that return one closure reading the carrier at call time, with the assignment made explicitly by the caller. The carrier is the `api` struct, and `New` assigns the result (`l.Set = SetFactory(&l)`, reading `l.Args`/`l.Used` or other struct state inside the closure). No methods bound into fields, no internal mirror type, no `Api()` projection. This is a binding rule — see `docs/References/RULES.md#factory-pattern`, the `Factories` spec (`docs/References/Meta/Factories/`), and `docs/Explanations/StructContracts.md`.

One trade-off, not caught by the compiler: **completeness is unchecked** — a field no factory fills is nil and panics on first call, so every factory must be called from its package's `New` constructor.

`sandbox/` is a **closed sandbox**: nothing in it may import `examples/`, a third-party module, or an OS-bound stdlib package (`os`, `net`, `syscall`, …). Every input the library needs from the outside world arrives as a plain function argument, e.g. `lib.New(args []string)`. This is a binding rule — see `docs/References/RULES.md` and `docs/Explanations/SandboxIsolation.md`.

- **`sandbox/new.go`** — package `lib`, the only wiring point consumers touch: `New(args []string) api.Lib`. Never imports anything OS-bound. Importers alias it: `verblib "github.com/MateusMoutinhoOrg/Verb/sandbox"`.
- **`sandbox/contracts/api/api.go`** — the output structs the lib hands back (`Lib`, `Entry`, …). Every field must be exported, or `sandbox/internal/` cannot fill it. **Every type in the project is declared here**, never in `internal/`. Types only — no function bodies.
- **`sandbox/internal/lib/`** — the lib's factories: one `<Field>Factory(l *api.Lib)` per function field of `api.Lib`, each returning the field's closure, plus `New(args []string) api.Lib` assigning every factory's return value into the matching field (`sandbox/new.go` just delegates to it). Go's `internal/` rule keeps it unreachable from `examples/` and consumers.
- **`sandbox/internal/<object>/`** — one package per object the lib creates, holding that object's `<Field>Factory` functions plus a `New(…) api.<Object>` constructor that runs them all. There is no separate `Factory` aggregate — `New` is the aggregate. **Factories only, no type declarations.** Packages here take no `internal_` prefix — the `internal/` parent already says it.
- **`examples/<name>/<name>.go`** — outside the sandbox; self-contained `package main` programs calling `lib.New` directly, e.g. with `os.Args[1:]`.

## Critical: this repo is documentation-driven

Changes are governed by required-reading docs, and several actions **must** update companion files in the same commit. Each tutorial in `docs/Tutorials/` covers exactly one goal — read the one matching your change:

| If you... | Read | And keep in sync |
|-----------|------|------------------|
| write or edit any `<Field>Factory` | `docs/References/Meta/Factories/Specs.md` | the `New` constructor that must call it |
| add/rename/delete any file or dir | `docs/References/Structure.md` | `docs/References/Structure.md` |
| add/rename/delete a `.md` file | `docs/Tutorials/AddDocument.md`, `RenameDocument.md`, `DeleteDocument.md` | Doc Index in `README.md` |
| add a lib function/object | `docs/Tutorials/AddLibFunction.md`, `AddLibObject.md` | `docs/References/PublicApi.md` (+ detail page in `docs/References/PublicApi/`, see `ExposePublicApi.md`) |
| need an OS/third-party input inside `sandbox/` | `docs/Explanations/SandboxIsolation.md` | pass it as a plain argument to `lib.New` or the relevant lib function, read from `examples/` |
| add/rename/delete a sample | `docs/Tutorials/AddSample.md` | Samples section in `README.md` |
| fork or adapt the template into a real library | `docs/Tutorials/ForkTemplate.md`, `AdaptExistingLib.md` | `docs/References/TemplateFileActions.md` (the per-file copy/create/rewrite/delete list both tutorials follow) |

`docs/References/RULES.md` is the binding rule set and `docs/References/Specs.md` is the index of every file specification; `AGENTS.md` points here. A new public lib function or object must be declared in `sandbox/contracts/api/api.go`, given a factory in `sandbox/internal/`, **and** called from that package's `New` constructor, or callers get a nil field — that's the most common footgun, and `go build` will not catch it.

## Conventions

- Code that consumes the library from outside it (`examples/`, third-party callers) aliases every import with the `verb` prefix: `verblib` (`sandbox`), `verbtypes` (`sandbox/contracts/api`). Files belonging to the library itself — `sandbox/` — keep the plain package names. See the Import Aliases rule in `docs/References/RULES.md`.
- Module path is `github.com/MateusMoutinhoOrg/Verb`; renaming it is a documented procedure — see `docs/Tutorials/RenameModule.md`.
- Public-facing lib API entries each get a detail page under `docs/References/PublicApi/` named `<pkg>.<Symbol>.md`.
- `docs/References/Meta/` holds the specifications: one directory per kind of file, each pairing a `Specs.md` (how the file must be shaped) with a `sample`. Never browse it — always locate a spec through `docs/References/Specs.md`.
