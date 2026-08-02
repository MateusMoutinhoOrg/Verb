# Parse an Option

## Description
Covers checking a boolean flag with `IsPresent` and reading a flag's value with `GetStringOption`, including handling a missing option and repeated occurrences. Installing and initializing the lib is covered by [LibInitialization.md](/docs/Tutorials/LibInitialization.md); draining the leftover positional arguments afterward is covered by [UseUnusedMechanic.md](/docs/Tutorials/UseUnusedMechanic.md).

---

## Workflow
1. Initialize the lib as shown in [LibInitialization.md](/docs/Tutorials/LibInitialization.md):
   ```go
   deps := verbadapter.New(os.Args[1:])
   l := verblib.New(deps)
   ```
2. Check a boolean flag with `IsPresent`, listing every spelling you accept:
   ```go
   quiet := l.IsPresent([]string{"-q", "--quiet"})
   ```
3. Read an option's value with `GetStringOption`, passing occurrence `0` for the first (or only) match. It returns an error when the flag is absent or has no following value — always check it:
   ```go
   output, err := l.GetStringOption([]string{"-o", "--output"}, 0)
   if err != nil {
       println("no output option given")
   } else {
       println("output:", output)
   }
   ```
4. For a repeatable option, size it first with `GetOptionsSize`, then loop over every occurrence:
   ```go
   size := l.GetOptionsSize([]string{"--tag"})
   for i := 0; i < size; i++ {
       tag, _ := l.GetStringOption([]string{"--tag"}, i)
       println("tag:", tag)
   }
   ```
5. For a typed value, use the matching typed getter instead of parsing the string yourself:
   ```go
   retries, err := l.GetIntOption([]string{"--retries"}, 0)
   if err != nil {
       println("invalid or missing --retries")
   }
   ```
6. Run the program:
   ```bash
   go run main.go --quiet -o out.txt --retries 3
   ```

---

## Full Code
```go
package main

import (
    "fmt"
    "os"

    verbadapter "github.com/MateusMoutinhoOrg/Verb/adapters/standard"
    verblib "github.com/MateusMoutinhoOrg/Verb/sandbox"
)

func main() {
    deps := verbadapter.New(os.Args[1:])
    l := verblib.New(deps)

    quiet := l.IsPresent([]string{"-q", "--quiet"})

    output, err := l.GetStringOption([]string{"-o", "--output"}, 0)
    if err != nil {
        output = "stdout"
    }

    retries, err := l.GetIntOption([]string{"--retries"}, 0)
    if err != nil {
        retries = 1
    }

    fmt.Println("quiet:", quiet, "output:", output, "retries:", retries)
}
```
