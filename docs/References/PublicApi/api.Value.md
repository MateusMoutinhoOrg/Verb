# `api.Entry.Value`

**Type:** Field

## Signature

```go
Value string
```

## Description

The cached value the entry was created with. A plain data field: it is written by the entry constructor when [`api.Lib.Get`](/docs/References/PublicApi/api.Get.md) builds the `Entry`, before any factory runs, and reading it touches no dependency.

## Type

| Type | Description |
| :--- | :--- |
| `string` | The stored value. |

## Examples

```go
package main

import (
	"fmt"

	verbadapter "github.com/MateusMoutinhoOrg/Verb/adapters/standard"
	verblib "github.com/MateusMoutinhoOrg/Verb/sandbox"
)

func main() {
	l := verblib.New(verbadapter.New("cache.json"))
	l.Set("k", "v", 60)

	entry, _ := l.Get("k")
	fmt.Println(entry.Value) // v
}
```
