# Sandbox Isolation

## Description
Explains why the library lives in `sandbox/` and what "closed sandbox" means in practice: the library reaches nothing outside itself — no third-party module, no OS-bound standard-library package — so everything it can do is exactly what its plain function inputs allow.

---

## The Two Trees

The project is split into two top-level directories, and the arrow only points one way:

```
sandbox/  ◀──  examples/
(closed)       (consumes the lib)
```

- `sandbox/` is the library. It is closed: it imports only itself and OS-independent standard-library packages (`time`, `strings`, `strconv`, `errors`, …).
- `examples/` is outside the wall — it is the only place `os` is read and handed to the library.

The split is what makes the library OS-independent: nothing inside `sandbox/` can be affected by which operating system, filesystem, or network the program runs on, because it has no way to reach any of them — every input it needs arrives as a plain function argument from its caller.

---

## What the Wall Forbids

A file under `sandbox/` may not import:

| Forbidden | Why |
|-----------|-----|
| `examples/…` | Samples are consumers of the library, never part of it. |
| Any third-party module | A dependency the caller cannot replace is a dependency the caller cannot test around. |
| OS-bound stdlib (`os`, `net`, `os/exec`, `syscall`, …) | Reading the real process argv is the caller's job — the library only ever receives the resulting `[]string`. |

An argv parser needs exactly one input from the outside world — the argument vector to parse — and it arrives as a plain parameter to `lib.New`:

```go
// sandbox/new.go — no os, no net, no third party
func New(args []string) api.Lib {
	return lib.New(args)
}
```

Everything downstream of that call — matching flags, tracking used positions, parsing typed values with `strconv`/`time` — is pure computation over `l.Args`/`l.Used` and needs no further door in the wall.

---

## What the Wall Forbids in the Other Direction

The wall is not only about what the sandbox imports — it also limits what the outside may reach into. `sandbox/internal/` holds the **factories** that fill the `api.Lib` fields, and it is protected by Go's `internal/` rule: only packages rooted at `sandbox/` can import it. A consumer that tries gets a compile error, not a convention warning.

Note what this does *not* hide. Because the factories fill the api struct directly, the struct the caller holds is the same one the library works on — including its `Args` and `Used` fields. What stays unreachable is the logic: a consumer can read `l.Args`/`l.Used`, but cannot call, replace, or re-run the factories that turned them into behavior.

So the outside world sees exactly two things:

| Package | Who imports it | For what |
|---------|----------------|----------|
| `sandbox` (package `lib`) | consumers, examples | `lib.New(args) api.Lib` — the single wiring point |
| `sandbox/contracts/api` | consumers, examples | the struct handed back |

Everything else in `sandbox/` is unreachable, which is why the factories can be renamed, split, or restructured without breaking a single consumer.

---

## Why the Entry Point Lives Inside

`sandbox/new.go` is the one file in the sandbox that consumers import directly, and it stays inside the wall because it obeys the same rule — it names no OS-bound package:

```go
// sandbox/new.go
func New(args []string) api.Lib {
	return lib.New(args)
}
```

It accepts a plain `[]string` and returns a contract struct. The caller decides what argv to hand it — usually the real process arguments — so the sandbox never has to read them itself:

```go
import (
	"os"

	verblib "github.com/MateusMoutinhoOrg/Verb/sandbox"
)

// This line is in examples/, outside the wall — the only place os.Args is read.
l := verblib.New(os.Args[1:])
```

For why the api contract is a struct rather than an interface, see [StructContracts.md](/docs/Explanations/StructContracts.md).

---

## If a Future Dependency Needs a Door

Should this library ever need an OS-bound or third-party effect beyond parsing a given `[]string` — reading a config file, the current time, and so on — that behavior would have to arrive as a function argument alongside `args`, never called directly from inside `sandbox/`. This template does not currently need one, so no such mechanism exists; add it only when a real effect requires it.
