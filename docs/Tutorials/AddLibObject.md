# Add a Library Object

## Description
Covers adding an object created by the library: its struct in [sandbox/contracts/api/api.go](../../sandbox/contracts/api/api.go), the factories that fill its function fields under [sandbox/internal/](../../sandbox/internal/), and the `New` constructor that builds it.

### Rules
- The object **is** its api struct. There is no internal mirror type: [sandbox/internal/](../../sandbox/internal/) holds only the factories and the constructor.
- Every field factory's return value must be assigned in the package's `New(...) api.<Object>` constructor, which doubles as the factory aggregate — an unassigned field stays nil and panics on first call.
- Every api field must be exported: the factories fill them from another package, and consumers read them.
- Adding a directory or file to [sandbox/internal/](../../sandbox/internal/) requires updating [Structure.md](/docs/References/Structure.md).

---

## Workflow
1. Declare the object's struct in [sandbox/contracts/api/api.go](../../sandbox/contracts/api/api.go). This walkthrough adds a `Bucket`, a namespaced view over a set of keys:
   ```go
   type Bucket struct {
       Prefix  string
       FullKey func(name string) string
   }
   ```
2. Create the object's package and file, both named after the object (e.g. `sandbox/internal/bucket/bucket.go`) — no `internal_` prefix, the `internal/` parent already marks it private — holding its field factories and its `New` constructor:
   ```go
   package bucket

   import (
       "github.com/MateusMoutinhoOrg/Verb/sandbox/contracts/api"
   )

   // FullKeyFactory returns the closure that fills api.Bucket.FullKey,
   // namespacing a name under the bucket's prefix.
   func FullKeyFactory(b *api.Bucket) func(name string) string {
       return func(name string) string {
           return b.Prefix + ":" + name
       }
   }

   // New builds an api.Bucket and runs every bucket factory over it,
   // assigning each return value into its matching function field.
   func New(prefix string) api.Bucket {
       b := api.Bucket{
           Prefix: prefix,
       }
       b.FullKey = FullKeyFactory(&b)
       return b
   }
   ```
3. Declare the constructor as a field of the `Lib` api struct in [sandbox/contracts/api/api.go](../../sandbox/contracts/api/api.go), returning the object's api struct:
   ```go
   type Lib struct {
       NewBucket func(prefix string) Bucket
   }
   ```
4. Write the constructor's factory in [sandbox/internal/lib/](../../sandbox/internal/lib/):
   ```go
   // NewBucketFactory returns the closure that fills api.Lib.NewBucket,
   // creating a Bucket.
   func NewBucketFactory(l *api.Lib) func(prefix string) api.Bucket {
       return func(prefix string) api.Bucket {
           return bucket.New(prefix)
       }
   }
   ```
5. Assign `NewBucketFactory`'s return value in the lib package's `New` constructor, as described in [AddLibFunction.md](/docs/Tutorials/AddLibFunction.md).
6. Expose the object, its constructor, and its fields following [ExposePublicApi.md](/docs/Tutorials/ExposePublicApi.md).
7. Register the new directory and file in [Structure.md](/docs/References/Structure.md).
8. If the object needs a runnable demonstration, add one following [AddSample.md](/docs/Tutorials/AddSample.md).
9. Build the project, then call the new field once to confirm it is not nil — a missing assignment in `New` compiles cleanly:
   ```bash
   go build ./...
   ```
