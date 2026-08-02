# `deps.Deps`

**Type:** Struct

## Definition

```go
type Deps struct {
	Args func() []string
}
```

## Description

The dependency contract every adapter must fill. An argv parser has exactly one OS-bound effect: reading the raw argument vector, so `Deps` has exactly one field. `Args` is called exactly once, when [`lib.New`](/docs/References/PublicApi/lib.New.md) builds the [`api.Lib`](/docs/References/PublicApi/api.Lib.md); its result is snapshotted so argument positions stay stable for the returned `Lib`'s whole lifetime. A filled `Deps` is built by an adapter — see [`standard.New`](/docs/References/PublicApi/standard.New.md) — and passed to `lib.New`.

Because it is a struct and not an interface, a value returned by an adapter can be patched field by field before injection, and a custom contract needs no type declaration at all. The trade-off: the compiler cannot detect a field you forgot to fill — it stays nil and panics on first call. See [StructContracts.md](/docs/Explanations/StructContracts.md) and [DepsMechanic.md](/docs/Explanations/DepsMechanic.md).

## Fields

| Field | Description |
| :--- | :--- |
| `Args func() []string` | Returns the argument vector to parse, e.g. `os.Args[1:]`. |

## Examples

```go
package main

import (
	"fmt"

	verblib "github.com/MateusMoutinhoOrg/Verb/sandbox"
	verbdeps "github.com/MateusMoutinhoOrg/Verb/sandbox/contracts/deps"
)

func main() {
	// A deps.Deps built without any adapter, parsing a fixed argv.
	d := verbdeps.Deps{
		Args: func() []string { return []string{"-o", "out.txt", "in.txt"} },
	}

	l := verblib.New(d)

	output, _ := l.GetStringOption([]string{"-o", "--output"}, 0)
	fmt.Println(output) // out.txt
}
```
