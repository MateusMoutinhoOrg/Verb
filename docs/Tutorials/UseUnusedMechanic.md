# Use the Unused Mechanic

## Description
Covers draining the leftover positional arguments with `GetNextStringArg` after every expected flag and option has been read. Reading flags and options first is covered by [ParseOption.md](/docs/Tutorials/ParseOption.md); the mechanic itself is explained in [UnnusedMechanic.md](/docs/Explanations/UnnusedMechanic.md).

---

## Workflow
1. Read every flag and option your program expects, as shown in [ParseOption.md](/docs/Tutorials/ParseOption.md):
   ```go
   quiet := l.IsPresent([]string{"-q", "--quiet"})
   output, _ := l.GetStringOption([]string{"-o", "--output"}, 0)
   ```
2. Read the first leftover positional argument with `GetNextStringArg`. It returns an error when nothing is left unread — always check it:
   ```go
   input, err := l.GetNextStringArg()
   if err != nil {
       println("missing input file")
       return
   }
   println("input:", input)
   ```
3. Call it again to drain further positional arguments, in the order they appeared on the command line:
   ```go
   for {
       arg, err := l.GetNextStringArg()
       if err != nil {
           break
       }
       println("extra arg:", arg)
   }
   ```
4. For a typed positional argument, use the matching typed getter instead of parsing the string yourself:
   ```go
   port, err := l.GetNextIntArg()
   ```
5. Run the program:
   ```bash
   go run main.go -o out.txt input.txt extra.txt
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

    output, _ := l.GetStringOption([]string{"-o", "--output"}, 0)

    input, err := l.GetNextStringArg()
    if err != nil {
        fmt.Println("missing input file")
        return
    }

    fmt.Println("output:", output, "input:", input)

    for {
        arg, err := l.GetNextStringArg()
        if err != nil {
            break
        }
        fmt.Println("extra arg:", arg)
    }
}
```
