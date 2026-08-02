# Outputs Specification

## Description
Defines the required shape of the output contracts in `sandbox/contracts/api/api.go` — the structs the library hands back to callers. These structs are also what the library works *on*: the factories in `sandbox/internal/` receive a pointer to one and fill its function fields with closures. There is no separate internal mirror type.

### Rules
- `api.go` must declare one **struct** per object the library hands back, including the `Lib` entry point returned by `lib.New`. Contracts are structs, never interfaces; see [StructContracts.md](/docs/Explanations/StructContracts.md).
- Behavior is exposed as **function fields** (`Name func(...) ...`), each filled by a factory in `sandbox/internal/`. Values fixed at construction time are plain data fields.
- `api.go` declares **types only** — never a function body. Every implementation lives in `sandbox/internal/`; see the [LibFunctions](/docs/References/Meta/LibFunctions/Specs.md) specification.
- Every field must be **exported**: `sandbox/internal/` fills them from another package, and consumers read them.
- A function field returning another library object must return that object's **api struct**; there is no internal type to return.
- `api.go` must not import anything from `examples/`, `sandbox/internal/`, or `sandbox` (the entry point) — the contract stays free of implementations.
- Exported structs must have a doc comment and be listed in [PublicApi.md](/docs/References/PublicApi.md).

## Structure
1. **Package clause**: `package api`.
2. **One struct per output object**: the object's plain data fields and the function fields its factories fill.
3. **`Lib` struct**: the entry point, declaring the data and function fields the library exposes.

> **Note**: For a concrete example, refer to [sample.go](./sample.go).
