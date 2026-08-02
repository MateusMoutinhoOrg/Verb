# `memory.New`

**Type:** Function

## Signature

```go
func New(args []string) deps.Deps
```

## Description

Creates a [`deps.Deps`](/docs/References/PublicApi/deps.Deps.md) backed by the memory adapter: it hands back a fixed, in-memory `args` slice instead of a real process's argv. It exists for tests and scripted scenarios — hand it any `[]string` literal and get a `deps.Deps` that parses exactly that, with no dependency on the actual `os.Args` the test binary was invoked with. The factory returns the **contract struct**, never the concrete `MemoryAdapter`, so consumers stay decoupled from the implementation. For all shipped adapters, see [Adapters.md](/docs/References/Adapters.md).

## Parameters

| Name | Type | Description |
| :--- | :--- | :--- |
| `args` | `[]string` | The fixed argument vector to parse. |

## Returns

| Type | Description |
| :--- | :--- |
| [`deps.Deps`](/docs/References/PublicApi/deps.Deps.md) | A fully filled dependency contract, ready to be passed to `lib.New`. |

## Examples

```go
package main

import (
	"fmt"

	verbadapter "github.com/MateusMoutinhoOrg/Verb/adapters/memory"
	verblib "github.com/MateusMoutinhoOrg/Verb/sandbox"
)

func main() {
	d := verbadapter.New([]string{"-q", "input.txt"})
	l := verblib.New(d)

	fmt.Println(l.IsPresent([]string{"-q", "--quiet"})) // true
}
```
