# Add a Library Function

## Description
Covers adding a function to the library: declaring it as a field of the `Lib` struct in [sandbox/contracts/api/api.go](../../sandbox/contracts/api/api.go), writing the factory that fills it in [sandbox/internal/lib/](../../sandbox/internal/lib/), and registering that factory in the package's `New` constructor.

### Rules
- A function is only usable once its factory's return value is assigned from the package's `New(args []string) api.Lib` constructor, which doubles as the factory aggregate — an unassigned field stays nil and panics on first call. The compiler does not catch this.
- One factory per field, named `<Field>Factory`, taking a single `*api.Lib` and returning one closure.
- State is read as `l.<field>` **inside** the closure, never captured at factory time — that is what keeps the struct authoritative.
- `sandbox/` is a closed sandbox: library code must never import [examples/](../../examples/), a third-party module, or an OS-bound standard-library package (`os`, `net`, `syscall`, …). See [SandboxIsolation.md](/docs/Explanations/SandboxIsolation.md).
- Adding a file to [sandbox/internal/](../../sandbox/internal/) requires updating [Structure.md](/docs/References/Structure.md).

---

## Workflow
1. Declare the function as a field of the `Lib` struct in [sandbox/contracts/api/api.go](../../sandbox/contracts/api/api.go):
   ```go
   type Lib struct {
       Args           []string
       Used           []bool
       IsPresent      func(flags []string) bool
       GetOptionsSize func(flags []string) int
       HasOption      func(flags []string) bool // new function
   }
   ```
2. Write its factory in a new or existing file in [sandbox/internal/lib/](../../sandbox/internal/lib/), with the identical signature, returning the closure:
   ```go
   // HasOptionFactory returns the closure that fills api.Lib.HasOption,
   // reporting whether an option occurs at least once, without consuming it.
   func HasOptionFactory(l *api.Lib) func(flags []string) bool {
       return func(flags []string) bool {
           return l.GetOptionsSize(flags) > 0
       }
   }
   ```
   > Calling another field from inside a closure (`l.GetOptionsSize` above) is fine: by the time `HasOption` runs, `New` has already filled every field.
3. Assign the factory's return value in the package's `New` constructor — without this line the field stays nil and the function panics when called:
   ```go
   func New(args []string) api.Lib {
       l := api.Lib{Args: args, Used: make([]bool, len(args))}
       l.IsPresent = IsPresentFactory(&l)
       l.GetOptionsSize = GetOptionsSizeFactory(&l)
       l.HasOption = HasOptionFactory(&l) // register the new function
       return l
   }
   ```
4. If the function returns a new object, create it following [AddLibObject.md](/docs/Tutorials/AddLibObject.md) and return the object's `api` struct.
5. Expose the function following [ExposePublicApi.md](/docs/Tutorials/ExposePublicApi.md).
6. If a new file was created, register it in [Structure.md](/docs/References/Structure.md).
7. If the function needs a runnable demonstration, add one following [AddSample.md](/docs/Tutorials/AddSample.md).
8. Build the project — a signature mismatch fails here, but a forgotten assignment in `New` does not:
   ```bash
   go build ./...
   ```
9. Call the new field once — from a sample or a test — to confirm it is not nil. That is the only check that catches a missing assignment.
