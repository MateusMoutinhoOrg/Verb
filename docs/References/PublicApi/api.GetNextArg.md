# `api.Lib.GetNextStringArg` / `GetNextIntArg` / `GetNextDoubleArg` / `GetNextTimestampArg`

**Type:** Fields

## Signatures

```go
GetNextStringArg    func() (string, error)
GetNextIntArg       func() (int, error)
GetNextDoubleArg    func() (float64, error)
GetNextTimestampArg func() (time.Time, error)
```

## Description

The core of the Unused Mechanic (see [UnnusedMechanic.md](/docs/Explanations/UnnusedMechanic.md)): returns the first argument in [`Args`](/docs/References/PublicApi/api.Args.md), in order, whose entry in [`Used`](/docs/References/PublicApi/api.Used.md) is still `false`, and marks it used. After every flag and option a program expects has been read with `IsPresent`, [`GetStringOption`](/docs/References/PublicApi/api.GetOption.md), or [`GetStringKeyValues`](/docs/References/PublicApi/api.GetKeyValues.md) (or their typed variants), whatever remains unread is exactly the leftover positional arguments — call one of these repeatedly to drain them in order.

- `GetNextStringArg` returns the argument's raw text.
- `GetNextIntArg` additionally parses it with `strconv.Atoi`.
- `GetNextDoubleArg` additionally parses it with `strconv.ParseFloat` into a `float64`.
- `GetNextTimestampArg` additionally parses it with `time.Parse` using `time.RFC3339`.

All four return an error when every argument has already been used. The typed variants additionally return a parse error when the value doesn't fit the target type.

## Parameters

_None._

## Returns

| Field | Type | Description |
| :--- | :--- | :--- |
| `GetNextStringArg` | `(string, error)` | The raw text of the next unread argument. |
| `GetNextIntArg` | `(int, error)` | The next unread argument parsed as a base-10 integer. |
| `GetNextDoubleArg` | `(float64, error)` | The next unread argument parsed as a 64-bit float. |
| `GetNextTimestampArg` | `(time.Time, error)` | The next unread argument parsed as an RFC 3339 timestamp. |

## Examples

```go
package main

import (
	"fmt"
	"os"

	verblib "github.com/MateusMoutinhoOrg/Verb/sandbox"
)

func main() {
	// ./cli -o teste.out teste.c
	l := verblib.New(os.Args[1:])

	output, err := l.GetStringOption([]string{"-o", "--output"}, 0)
	if err != nil {
		panic(err)
	}

	// -o and teste.out are already used; this returns "teste.c".
	file, err := l.GetNextStringArg()
	if err != nil {
		panic(err)
	}

	fmt.Println(output, file)
}
```
