# `api.Lib.GetOptionsSize` / `api.Lib.GetKeyValuesSize`

**Type:** Fields

## Signatures

```go
GetOptionsSize   func(flags []string) int
GetKeyValuesSize func(prefixes []string) int
```

## Description

Both fields count matches in [`Args`](/docs/References/PublicApi/api.Args.md) without mutating [`Used`](/docs/References/PublicApi/api.Used.md) — they are the only `Get`-shaped fields on `Lib` that neither consume input nor can fail, so they return a plain `int` instead of participating in the Unused Mechanic or returning an error.

- `GetOptionsSize` counts how many arguments equal one of the given flag spellings (e.g. `[]string{"-o", "--output"}`), regardless of whether they were already read.
- `GetKeyValuesSize` counts how many arguments start with one of the given `key=value` prefixes (the separator is part of the prefix, e.g. `[]string{"user=", "username="}`), regardless of whether they were already read.

Call one of these before looping over occurrence indices `0..size-1` with [`GetStringOption`/typed variants](/docs/References/PublicApi/api.GetOption.md) or [`GetStringKeyValues`/typed variants](/docs/References/PublicApi/api.GetKeyValues.md) to read every occurrence of a repeatable option or key.

## Parameters

| Field | Parameter | Type | Description |
| :--- | :--- | :--- | :--- |
| `GetOptionsSize` | `flags` | `[]string` | Every accepted spelling of the option's flag. |
| `GetKeyValuesSize` | `prefixes` | `[]string` | Every accepted `key=` prefix, separator included. |

## Returns

| Type | Description |
| :--- | :--- |
| `int` | The number of matching arguments currently in `Args`. |

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

	size := l.GetOptionsSize([]string{"-username", "--username"})
	for i := 0; i < size; i++ {
		v, err := l.GetStringOption([]string{"-username", "--username"}, i)
		if err != nil {
			continue
		}
		fmt.Println("username:", v)
	}
}
```
