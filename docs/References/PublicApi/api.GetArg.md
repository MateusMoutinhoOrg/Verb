# `api.Lib.GetStringArg` / `GetIntArg` / `GetDoubleArg` / `GetTimestampArg`

**Type:** Fields

## Signatures

```go
GetStringArg    func(index int) (string, error)
GetIntArg       func(index int) (int, error)
GetDoubleArg    func(index int) (float64, error)
GetTimestampArg func(index int) (time.Time, error)
```

## Description

Reads the argument at an **absolute** index of [`Args`](/docs/References/PublicApi/api.Args.md) — the same numbering as the raw command line, where index `0` is the first argument after the program name (`./app test` → index `0` is `"test"`), independent of which arguments have already been read. The matched index is marked used in [`Used`](/docs/References/PublicApi/api.Used.md).

- `GetStringArg` returns the argument's raw text.
- `GetIntArg` additionally parses it with `strconv.Atoi`.
- `GetDoubleArg` additionally parses it with `strconv.ParseFloat` into a `float64`.
- `GetTimestampArg` additionally parses it with `time.Parse` using `time.RFC3339`.

All four return an error when `index` is negative or beyond the end of `Args`. The typed variants additionally return a parse error when the value doesn't fit the target type.

## Parameters

| Parameter | Type | Description |
| :--- | :--- | :--- |
| `index` | `int` | Absolute position in `Args`, 0-based. |

## Returns

| Field | Type | Description |
| :--- | :--- | :--- |
| `GetStringArg` | `(string, error)` | The raw argument text. |
| `GetIntArg` | `(int, error)` | The argument parsed as a base-10 integer. |
| `GetDoubleArg` | `(float64, error)` | The argument parsed as a 64-bit float. |
| `GetTimestampArg` | `(time.Time, error)` | The argument parsed as an RFC 3339 timestamp. |

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

	// ./app test -> first == "test"
	first, err := l.GetStringArg(0)
	if err != nil {
		panic(err)
	}
	fmt.Println("first arg:", first)
}
```
