# Examples Specification

## Description
Defines the required shape of a runnable example in `examples/<example>/<example>.go`. An example is a self-contained `package main` program that calls the lib's entry point to demonstrate real usage.

### Rules
- Each example lives in its own directory under `examples/` named after the feature it demonstrates (e.g. `examples/ExampleSample/`).
- The file is named after its directory (`<example>/<example>.go`) and declares `package main` with a `main` function.
- An example calls `lib.New(...)` directly with the arguments it needs (e.g. `os.Args[1:]`), which returns an `api.Lib`.
- An example may import `sandbox` (aliased `verblib`) and `sandbox/contracts/api` (aliased `verbtypes`); it must never import `sandbox/internal/` — Go's `internal/` rule rejects it.
- Examples live outside the sandbox and are the only place `os` (or another OS-bound package) and the library are named in the same file.
- Keep examples minimal and runnable via `go run ./examples/<example>/<example>.go`; add explanatory comments on the key steps.
- Adding, renaming, or deleting an example requires updating the Samples section of [README.md](/README.md) — see [AddSample.md](/docs/Tutorials/AddSample.md).

## Structure
1. **Package clause**: `package main`.
2. **Imports**: every import of this module is aliased with the `verb` prefix, so a reader sees at a glance which layer a call belongs to — the sandbox entry point as `verblib "github.com/MateusMoutinhoOrg/Verb/sandbox"`, and, when the example names an output type, `verbtypes "github.com/MateusMoutinhoOrg/Verb/sandbox/contracts/api"`.
3. **`main` function**: build the lib with `verblib.New(...)`, then exercise the returned `verbtypes.Lib`.

> **Note**: For a concrete example, refer to [sample.go](./sample.go).
