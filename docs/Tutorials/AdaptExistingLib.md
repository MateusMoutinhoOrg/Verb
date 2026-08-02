# Adapt a Pre-Existing Library

## Description
Covers converting a library that already exists into this project's struct-of-functions structure. To start a new library from scratch, follow [ForkTemplate.md](/docs/Tutorials/ForkTemplate.md) instead.

### Rules
- Read [RULES.md](/docs/References/RULES.md) and [Structure.md](/docs/References/Structure.md) before starting.
- Keep the separation defined in [Structure.md](/docs/References/Structure.md): public contract structs in `sandbox/contracts/`, internal types and logic in `sandbox/internal/`, and the entry point in `sandbox/`. Contracts are structs of function fields, never interfaces — see [StructContracts.md](/docs/Explanations/StructContracts.md).
- Every file of the template has one action — **Copy**, **Create**, **Rewrite**, or **Delete**. Take it from [TemplateFileActions.md](/docs/References/TemplateFileActions.md); the steps below follow that order.
- The pre-existing package layout does **not** survive: all library logic ends up in `sandbox/internal/`. Code left in its original packages, or still calling `os`/`net`/third-party APIs directly, is not adapted — any such effect must arrive as a plain argument passed in from `examples/`, never read inside `sandbox/`.
- Every public type the library returns becomes a contract struct in `sandbox/contracts/api`, whose function fields are filled by factories in `sandbox/internal/`. A type still declared in `sandbox/internal/` and handed back to callers is not adapted.
- Every file created or rewritten — code and `.md` alike — must follow its specification, located through [Specs.md](/docs/References/Specs.md).
- The adaptation is not complete until the final checklist in the last workflow step passes.

---

## Workflow
1. Recreate this project's directory layout inside the library being converted, using [Structure.md](/docs/References/Structure.md) as reference.
2. Copy every **[Copy](/docs/References/TemplateFileActions.md#copy)** file into the library unchanged — the specifications, rules, tutorials, explanations, and [sandbox/new.go](../../sandbox/new.go).
3. Rewrite `sandbox/contracts/api/api.go` with the `Lib` struct and one struct per type the library hands back, following [AddLibFunction.md](/docs/Tutorials/AddLibFunction.md) and [AddLibObject.md](/docs/Tutorials/AddLibObject.md).
4. Rewrite the existing library code into `sandbox/internal/`: move each source file in, turn each public function into a `<Field>Factory(l *api.Lib)` that returns a closure for the matching api field, assign every factory's return value from the package's `New` constructor, and replace **every** OS-bound or third-party call with a plain argument passed into `lib.New`, following [AddLibFunction.md](/docs/Tutorials/AddLibFunction.md) and [AddLibObject.md](/docs/Tutorials/AddLibObject.md). Do not keep the code in its original packages, leave methods on internal types, or leave direct calls in place.
5. Create the samples in `examples/` demonstrating the converted entry points, following [AddSample.md](/docs/Tutorials/AddSample.md).
6. Create the detail pages in `docs/References/PublicApi/` and rewrite `docs/References/PublicApi.md`, following [ExposePublicApi.md](/docs/Tutorials/ExposePublicApi.md).
7. Delete every **[Delete](/docs/References/TemplateFileActions.md#delete)** file carried over from the template, plus the pre-existing code the converted lib replaced. For `.md` files, follow [DeleteDocument.md](/docs/Tutorials/DeleteDocument.md).
8. Rewrite `docs/References/Structure.md` to describe the library's actual layout.
9. Create the tutorials specific to the converted library — one page per workflow its maintainers will repeat (e.g. adding a domain object, extending a feature, releasing) — following [AddDocument.md](/docs/Tutorials/AddDocument.md) and the [TutorialDocs specification](/docs/References/Meta/TutorialDocs/Specs.md). The template tutorials copied in step 2 cover the structure only; they do not document the library's own use cases.
10. Create any reference page the library needs beyond the public API — following [AddDocument.md](/docs/Tutorials/AddDocument.md) and the [ReferenceDocs specification](/docs/References/Meta/ReferenceDocs/Specs.md).
11. Rewrite the `README.md`: overview, quick start, badges, Doc Index, and Samples section.
12. Verify the result:
```bash
go build ./...
```
Then confirm every item below — the adaptation is only done when all pass:
- All library logic lives in `sandbox/internal/`; no file there imports `os`, `net`, or a third-party implementation directly — every such input arrives as a plain argument to `lib.New` or a lib function, read from `examples/`.
- `sandbox/contracts/api/api.go` declares every public object as a struct, and every one of its function fields is filled by a factory registered in that package's `New` constructor.
- `sandbox/new.go` is the only wiring point, and it imports nothing outside `sandbox/`.
- Tutorials and reference pages specific to this library exist under `docs/Tutorials/` and `docs/References/`.
- Every created or rewritten file matches its specification from [Specs.md](/docs/References/Specs.md).
- The `README.md` Doc Index lists every `.md` file and the Samples section lists every sample.
