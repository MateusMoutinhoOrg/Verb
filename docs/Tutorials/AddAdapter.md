# Add an Adapter

## Description
Covers creating a new opinionated implementation of the `Deps` contract under [adapters/](../../adapters/).

### Rules
- Each adapter lives in its own directory under [adapters/](../../adapters/) and uses a package named after that directory.
- The adapter is a struct carrying a `Deps deps.Deps` field, filled by one **factory** per field of the contract — the same factory pattern `sandbox/internal/` uses, and a binding rule of the project. See [RULES.md](/docs/References/RULES.md#factory-pattern).
- Fields are never filled by binding methods of the adapter. Methods may exist only as unexported helpers a closure calls.
- A single `New(...) deps.Deps` constructor calls **every** field factory, assigns its return value into the matching field, and returns the `deps.Deps` contract struct, never the concrete adapter type.
- Filling every field is the author's job — an unassigned field compiles and panics on first call. See [StructContracts.md](/docs/Explanations/StructContracts.md).
- An adapter lives outside the sandbox and is the only place OS-bound and third-party code is allowed. It may import [sandbox/contracts/deps](../../sandbox/contracts/deps/), but never [sandbox/internal/](../../sandbox/internal/) — see [SandboxIsolation.md](/docs/Explanations/SandboxIsolation.md).
- The adapter file must follow its specification — locate it in [Specs.md](/docs/References/Specs.md).

---

## Workflow
1. Create the adapter directory and its file, both named after the adapter (e.g., `adapters/memory/memory.go`).
2. Declare the package and the adapter struct — the **carrier**, leading with the `Deps` field its factories fill, followed by its configuration and state:
   ```go
   package responsefile

   import (
       "os"
       "strings"

       "github.com/MateusMoutinhoOrg/Verb/sandbox/contracts/deps"
   )

   // ResponseFileAdapter fills deps.Deps by reading the argv to parse from a
   // "response file" — one argument per line — instead of the real os.Args.
   type ResponseFileAdapter struct {
       // Deps is the contract this adapter fills; its factories assign into it.
       Deps     deps.Deps
       filePath string
   }
   ```
3. Write one `<Field>Factory` per field of the `Deps` contract, each returning a single closure that reads the adapter's state through the pointer:
   ```go
   // ArgsFactory returns the closure that fills deps.Deps.Args, reading the
   // response file and splitting it one argument per line.
   func ArgsFactory(f *ResponseFileAdapter) func() []string {
       return func() []string {
           raw, err := os.ReadFile(f.filePath)
           if err != nil {
               return nil
           }
           lines := strings.Split(strings.TrimSpace(string(raw)), "\n")
           args := make([]string, 0, len(lines))
           for _, line := range lines {
               if line = strings.TrimSpace(line); line != "" {
                   args = append(args, line)
               }
           }
           return args
       }
   }
   ```
   Reading `f.filePath` inside the closure — instead of reading and splitting the file when the factory runs — is what keeps the argv current if the file changes before `Args` is called.
4. Expose the `New` constructor: build the adapter instance, run every field factory over it, assign each return value into its matching field, and return its `Deps`:
   ```go
   // New creates a deps.Deps that parses the argv found in filePath, one
   // argument per line, instead of the real process argv.
   func New(filePath string) deps.Deps {
       adapter := &ResponseFileAdapter{filePath: filePath}
       adapter.Deps.Args = ArgsFactory(adapter)
       return adapter.Deps
   }
   ```
5. Compare the assignments in your `New` against `sandbox/contracts/deps/deps.go` field by field. A missing field will **not** fail the build.
6. Register the new directory and file in [Structure.md](/docs/References/Structure.md), and add a row for the adapter in [Adapters.md](/docs/References/Adapters.md).
7. If the adapter is public-facing, expose its `New` factory following [ExposePublicApi.md](/docs/Tutorials/ExposePublicApi.md).
8. If the adapter needs a runnable demonstration, add one following [AddSample.md](/docs/Tutorials/AddSample.md).
9. Build the project and exercise the adapter:
   ```bash
   go build ./...
   ```
