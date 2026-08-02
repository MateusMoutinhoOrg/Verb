# `api.Lib.Deps`

**Type:** Field

## Signature

```go
Deps deps.Deps
```

## Description

The injected dependency set the struct was built with. [`lib.New`](/docs/References/PublicApi/lib.New.md) writes it onto [`api.Lib`](/docs/References/PublicApi/api.Lib.md), and reads it exactly once — to call `Deps.Args()` and take the `Args` snapshot — before any function field is called. Every function field on `Lib` is a closure that reaches its behavior through `Args`/`Used` rather than through `Deps` directly, since parsing is pure computation once the argv is known. See [DepsMechanic.md](/docs/Explanations/DepsMechanic.md).

The field is exported because the library's own factories, which live in another package, must read it. It is **not** a customization point after construction: the snapshot in `Args` is already taken, so assigning to `Deps` on a struct you already hold changes nothing. To parse a different argv, patch the [`deps.Deps`](/docs/References/PublicApi/deps.Deps.md) value **before** passing it to `lib.New`.

## Returns

| Type | Description |
| :--- | :--- |
| [`deps.Deps`](/docs/References/PublicApi/deps.Deps.md) | The filled dependency contract; read-only after construction. |

## Examples

### Patch Before Injection, Not After

```go
package main

import (
	"fmt"

	verbadapter "github.com/MateusMoutinhoOrg/Verb/adapters/standard"
	verblib "github.com/MateusMoutinhoOrg/Verb/sandbox"
)

func main() {
	myDeps := verbadapter.New([]string{"-q"})

	// Correct: replace Args while it is still a plain deps value.
	myDeps.Args = func() []string { return []string{"--quiet", "input.txt"} }

	l := verblib.New(myDeps)

	// Wrong: New() already called Deps.Args() and snapshotted it into l.Args
	// — this assignment is inert.
	l.Deps.Args = func() []string { return nil }

	fmt.Println(l.IsPresent([]string{"-q", "--quiet"})) // true
}
```
