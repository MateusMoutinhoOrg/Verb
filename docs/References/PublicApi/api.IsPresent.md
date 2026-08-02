# `api.Lib.IsPresent`

**Type:** Field

## Signature

```go
IsPresent func(flags []string) bool
```

## Description

Reports whether any of the given flag spellings (e.g. `[]string{"-q", "--quiet"}`) occurs anywhere in the unread portion of [`Args`](/docs/References/PublicApi/api.Args.md). On a match it marks that single argument as used in [`Used`](/docs/References/PublicApi/api.Used.md) and returns `true`; if none of the flags is found it returns `false` and nothing is marked.

It never returns an error: "not present" is a valid, expected outcome for a boolean flag, not a failure.

## Parameters

| Parameter | Type | Description |
| :--- | :--- | :--- |
| `flags` | `[]string` | Every accepted spelling of the flag, e.g. `-f`, `--file`. |

## Returns

| Type | Description |
| :--- | :--- |
| `bool` | `true` if one of the flags was found (and consumed), `false` otherwise. |

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
	l := verblib.New(verbadapter.New(os.Args[1:]))

	quiet := l.IsPresent([]string{"-q", "--quiet"})
	fmt.Println("quiet:", quiet)
}
```
