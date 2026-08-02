# Library Initialization

## Description
Covers installing the library and initializing it with the standard adapter in a new program. Using the parser after initialization is covered by [ParseOption.md](/docs/Tutorials/ParseOption.md). For other ways to build the dependencies, see [DepsMechanic.md](/docs/Explanations/DepsMechanic.md).

---

## Workflow
1. Install the lib:
   ```bash
   go get github.com/MateusMoutinhoOrg/Verb@v0.0.7
   ```
2. Create a file called `main.go` with the following code:
   ```go
   package main

   // 1. Import the standard adapter and the lib
   import (
       "os"

       verbadapter "github.com/MateusMoutinhoOrg/Verb/adapters/standard"
       verblib "github.com/MateusMoutinhoOrg/Verb/sandbox"
   )

   func main() {
       // 2. Create deps via an adapter (the "opinionated" part: hands
       //    back the real process argv)
       deps := verbadapter.New(os.Args[1:])

       // 3. Inject deps into the pure library — an argv parser
       l := verblib.New(deps)

       // 4. Use the library — it never knows which adapter is behind the scenes
       quiet := l.IsPresent([]string{"-q", "--quiet"})
       println("quiet:", quiet)
   }
   ```
3. Run the code:
   ```bash
   go run main.go --quiet
   ```
