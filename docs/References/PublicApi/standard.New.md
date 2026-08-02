# `standard.New`

**Type:** Function

## Signature

```go
func New(args []string) deps.Deps
```

## Description

Creates a [`deps.Deps`](/docs/References/PublicApi/deps.Deps.md) backed by the standard adapter, which fills the contract using the Go standard library only: it hands back whatever `args` slice it was constructed with. The intended call is `standard.New(os.Args[1:])`, reading the process's real command line — the adapter's job is to be the one place that reads `os.Args`, so the sandbox never has to. The factory returns the **contract struct**, never the concrete `StandardAdapter`, so consumers stay decoupled from the implementation — the single field is filled by a factory whose closure reads the adapter instance, which is how the adapter's state travels with the deps. For all shipped adapters, see [Adapters.md](/docs/References/Adapters.md).

## Parameters

| Name | Type | Description |
| :--- | :--- | :--- |
| `args` | `[]string` | The argument vector to parse — pass `os.Args[1:]` for the real command line. |

## Returns

| Type | Description |
| :--- | :--- |
| [`deps.Deps`](/docs/References/PublicApi/deps.Deps.md) | A fully filled dependency contract, ready to be passed to `lib.New`. |

## Examples

```go
package main

import (
	"fmt"
	"os"

	verbadapter "github.com/MateusMoutinhoOrg/Verb/adapters/standard"
	verblib "github.com/MateusMoutinhoOrg/Verb/sandbox"
)

func main() {
	d := verbadapter.New(os.Args[1:])
	l := verblib.New(d)

	if l.IsPresent([]string{"-q", "--quiet"}) {
		fmt.Println("quiet mode")
	}
}
```
