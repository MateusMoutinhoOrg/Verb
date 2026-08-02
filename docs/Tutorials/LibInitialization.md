# Library Initialization

## Description
Covers installing the library and initializing it in a new program. Using the parser after initialization is covered by [ParseOption.md](/docs/Tutorials/ParseOption.md). For why the entry point takes a plain argument instead of an injected dependency, see [SandboxIsolation.md](/docs/Explanations/SandboxIsolation.md).

---

## Workflow
1. Install the lib:
   ```bash
   go get github.com/MateusMoutinhoOrg/Verb@v0.0.1
   ```
2. Create a file called `main.go` with the following code:
   ```go
   package main

   // 1. Import the lib
   import (
       "os"

       verblib "github.com/MateusMoutinhoOrg/Verb/sandbox"
   )

   func main() {
       // 2. Build the library directly from the real process argv
       l := verblib.New(os.Args[1:])

       // 3. Use the library
       quiet := l.IsPresent([]string{"-q", "--quiet"})
       println("quiet:", quiet)
   }
   ```
3. Run the code:
   ```bash
   go run main.go --quiet
   ```
