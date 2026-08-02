# `lib.New`

**Type:** Function

## Signature

```go
func New(args []string) api.Lib
```

## Description

Builds and returns the [`api.Lib`](/docs/References/PublicApi/api.Lib.md) entry point from the argument vector to parse. It stores `args` on the struct's `Args` field, allocates the matching `Used` tracking slice, then runs the factories in `sandbox/internal/lib/` over it, each filling one function field with a closure that reads `Args`/`Used` at call time. The package is named `lib` and lives at `sandbox/`, so importers alias it: `verblib "github.com/MateusMoutinhoOrg/Verb/sandbox"` — matching the `verblib` / `verbtypes` alias convention used by every consumer of this module.

## Parameters

| Parameter | Type | Description |
| :--- | :--- | :--- |
| `args` | `[]string` | The argument vector to parse — pass `os.Args[1:]` for the real command line. |

## Returns

| Type | Description |
| :--- | :--- |
| [`api.Lib`](/docs/References/PublicApi/api.Lib.md) | A ready-to-use parser instance holding the given argv. |

## Examples

### Basic Initialization

```go
package main

import (
	"log"
	"os"

	verblib "github.com/MateusMoutinhoOrg/Verb/sandbox"
)

func main() {
	// Build the library directly from the real process argv.
	l := verblib.New(os.Args[1:])

	// The library instance 'l' is now ready for use.
	log.Println("Library successfully initialized:", l.IsPresent != nil)
}
```
