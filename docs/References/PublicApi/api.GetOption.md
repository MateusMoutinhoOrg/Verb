# `api.Lib.GetStringOption` / `GetIntOption` / `GetDoubleOption` / `GetTimestampOption`

**Type:** Fields

## Signatures

```go
GetStringOption    func(flags []string, occurrence int) (string, error)
GetIntOption       func(flags []string, occurrence int) (int, error)
GetDoubleOption    func(flags []string, occurrence int) (float64, error)
GetTimestampOption func(flags []string, occurrence int) (time.Time, error)
```

## Description

Finds the `occurrence`-th (0-based) argument that equals one of the given flag spellings, then reads the argument immediately following it as the option's value — the `-o out.txt` shape. Both the flag and its value are marked used in [`Used`](/docs/References/PublicApi/api.Used.md), even when a typed parse below fails, because the option was still found and read.

- `GetStringOption` returns the value's raw text.
- `GetIntOption` additionally parses it with `strconv.Atoi`.
- `GetDoubleOption` additionally parses it with `strconv.ParseFloat` into a `float64`.
- `GetTimestampOption` additionally parses it with `time.Parse` using `time.RFC3339` (e.g. `"2024-01-02T15:04:05Z"`).

All four return an error when `occurrence` is out of range for the number of matches (see [`GetOptionsSize`](/docs/References/PublicApi/api.Size.md)), or when the matched flag is the last argument and has no following value. The typed variants additionally return a parse error when the value doesn't fit the target type.

## Parameters

| Parameter | Type | Description |
| :--- | :--- | :--- |
| `flags` | `[]string` | Every accepted spelling of the option's flag. |
| `occurrence` | `int` | Which match to read, 0-based, in argv order. |

## Returns

| Field | Type | Description |
| :--- | :--- | :--- |
| `GetStringOption` | `(string, error)` | The raw value text. |
| `GetIntOption` | `(int, error)` | The value parsed as a base-10 integer. |
| `GetDoubleOption` | `(float64, error)` | The value parsed as a 64-bit float. |
| `GetTimestampOption` | `(time.Time, error)` | The value parsed as an RFC 3339 timestamp. |

## Examples

```go
package main

import (
	"fmt"
	"os"

	verblib "github.com/MateusMoutinhoOrg/Verb/sandbox"
)

func main() {
	l := verblib.New(os.Args[1:])

	// ./cli --retries 3 --timeout 1.5 --since 2024-01-02T15:04:05Z
	retries, err := l.GetIntOption([]string{"--retries"}, 0)
	if err != nil {
		panic(err)
	}
	timeout, err := l.GetDoubleOption([]string{"--timeout"}, 0)
	if err != nil {
		panic(err)
	}
	since, err := l.GetTimestampOption([]string{"--since"}, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(retries, timeout, since)
}
```
