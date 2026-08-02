
The Unused Mechanic

Every argument the library parses starts out unread. Calling `IsPresent` or
any `Get*` function marks the argument(s) it matched as read (`api.Lib.Used`
flips to `true` at that index). `GetNextStringArg` (and its typed variants)
walks `Used` in order and returns the first argument still marked unread —
so once every flag and option a program expects has been read, whatever is
left over is exactly the positional arguments nobody explicitly asked for.

Example:
```bash
./cli -o teste.out --quiet teste.c
```
in a cli like that will have theses values:

```go

package main

import (
	"os"

	verbadapter "github.com/MateusMoutinhoOrg/Verb/adapters/standard"
	verblib "github.com/MateusMoutinhoOrg/Verb/sandbox"
)

func main() {
	// Build deps via an adapter (the real process argv) and inject.
	lib := verblib.New(verbadapter.New(os.Args[1:]))

	// marks -o and teste.out as used
	output, err := lib.GetStringOption([]string{"-o", "--o", "--output", "--out"}, 0)
	if err != nil {
		panic(err)
	}

	// marks -q, --q, --quiet as used
	quiet := lib.IsPresent([]string{"-q", "--q", "--quiet"})

	// gets teste.c since it's the first unused item
	file, err := lib.GetNextStringArg()
	if err != nil {
		panic(err)
	}

	println(output, quiet, file)
}

```
