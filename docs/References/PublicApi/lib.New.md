# `lib.New`

**Type:** Function

## Signature

```go
func New(d deps.Deps) api.Lib
```

## Description

Injects a filled dependency contract into the library and returns the [`api.Lib`](/docs/References/PublicApi/api.Lib.md) entry point. It stores the deps on the struct's `Deps` field, calls `d.Args()` once to snapshot the argv into `Args`, allocates the matching `Used` tracking slice, then runs the factories in `sandbox/internal/lib/` over it, each filling one function field with a closure that reads `Args`/`Used` at call time. This is the only wiring point: `sandbox` never imports an adapter, so the caller chooses which implementation to pass. The package is named `lib` and lives at `sandbox/`, so importers alias it: `verblib "github.com/MateusMoutinhoOrg/Verb/sandbox"` — matching the `verbadapter` / `verblib` / `verbtypes` alias convention used by every consumer of this module.

## Parameters

| Parameter | Type | Description |
| :--- | :--- | :--- |
| `d` | [`deps.Deps`](/docs/References/PublicApi/deps.Deps.md) | A dependency contract with every field filled, usually built by an adapter. |

## Returns

| Type | Description |
| :--- | :--- |
| [`api.Lib`](/docs/References/PublicApi/api.Lib.md) | A ready-to-use parser instance holding a snapshot of the injected argv. |

## Examples

### Basic Initialization

```go
package main

import (
	"log"
	"os"

	verbadapter "github.com/MateusMoutinhoOrg/Verb/adapters/standard"
	verblib "github.com/MateusMoutinhoOrg/Verb/sandbox"
)

func main() {
	// 1. Build the deps through an adapter, wrapping the real process argv
	deps := verbadapter.New(os.Args[1:])

	// 2. Inject them into the library
	l := verblib.New(deps)

	// The library instance 'l' is now ready for use.
	log.Println("Library successfully initialized:", l.IsPresent != nil)
}
```
