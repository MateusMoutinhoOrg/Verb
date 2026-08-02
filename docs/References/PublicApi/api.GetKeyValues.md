# `api.Lib.GetStringKeyValues` / `GetIntKeyValues` / `GetDoubleKeyValues` / `GetTimestampKeyValues`

**Type:** Fields

## Signatures

```go
GetStringKeyValues    func(prefixes []string, occurrence int) (string, error)
GetIntKeyValues       func(prefixes []string, occurrence int) (int, error)
GetDoubleKeyValues    func(prefixes []string, occurrence int) (float64, error)
GetTimestampKeyValues func(prefixes []string, occurrence int) (time.Time, error)
```

## Description

Finds the `occurrence`-th (0-based) argument that starts with one of the given `key=value` prefixes — the separator is part of the prefix, e.g. `[]string{"user=", "username="}` matches `"username=alice"` — then reads the text after the matched prefix as the value. The matched argument is marked used in [`Used`](/docs/References/PublicApi/api.Used.md).

- `GetStringKeyValues` returns the value's raw text.
- `GetIntKeyValues` additionally parses it with `strconv.Atoi`.
- `GetDoubleKeyValues` additionally parses it with `strconv.ParseFloat` into a `float64`.
- `GetTimestampKeyValues` additionally parses it with `time.Parse` using `time.RFC3339`.

All four return an error when `occurrence` is out of range for the number of matches (see [`GetKeyValuesSize`](/docs/References/PublicApi/api.Size.md)), or when the matched argument's value portion is empty (e.g. bare `"username="` with nothing after it). The typed variants additionally return a parse error when the value doesn't fit the target type.

## Parameters

| Parameter | Type | Description |
| :--- | :--- | :--- |
| `prefixes` | `[]string` | Every accepted `key=` prefix, separator included. |
| `occurrence` | `int` | Which match to read, 0-based, in argv order. |

## Returns

| Field | Type | Description |
| :--- | :--- | :--- |
| `GetStringKeyValues` | `(string, error)` | The raw value text. |
| `GetIntKeyValues` | `(int, error)` | The value parsed as a base-10 integer. |
| `GetDoubleKeyValues` | `(float64, error)` | The value parsed as a 64-bit float. |
| `GetTimestampKeyValues` | `(time.Time, error)` | The value parsed as an RFC 3339 timestamp. |

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

	size := l.GetKeyValuesSize([]string{"username=", "user="})
	for i := 0; i < size; i++ {
		v, err := l.GetStringKeyValues([]string{"username=", "user="}, i)
		if err != nil {
			continue
		}
		fmt.Println("username:", v)
	}
}
```
