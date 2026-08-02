# Factories Specification

## Description
Defines the **factory pattern**, the single way any struct of function fields is filled inside `sandbox/internal/`. A factory takes a pointer to the struct that carries the state and returns exactly one field's value; the package's `New` constructor assigns that return value into the field. This spec describes the shape every factory shares; the per-tree specifications ([LibFunctions](/docs/References/Meta/LibFunctions/Specs.md), [LibObjects](/docs/References/Meta/LibObjects/Specs.md)) build on it and add what is specific to their tree.

### Rules

#### Carrier and Target
- Every factory takes **one** parameter: a pointer to the **carrier** — the api struct being filled (`*api.Lib`, `*api.<Object>`).
- A factory returns the field's type. Its only job is to build and return that value; assignment happens at the call site in `New`.

#### One Field, One Factory
- One factory per field, named `<Field>Factory` after the field it fills — `SetFactory` fills `Set`.
- The body returns a single value for that field, and touches no other field.
- Every field in this project's api structs is a function, so its factory returns a **closure**.
- The returned closure's signature must match the field's declaration in the contract exactly.

#### State Is Read Through the Pointer
- The closure reads state through the carrier pointer — `l.Multiplier`, `o.FirstProp` — so the value is resolved when the field is **called**, never captured when the factory ran.
- Copying a field into a local variable at factory time freezes it, and defeats the pattern.

#### `New` Is the Aggregate
- Every package exposing a filled struct declares a `New(...)` constructor that builds the carrier, calls **every** factory in the package exactly once and assigns its return value into the matching field, and returns the filled struct **by value** — `api.Lib` or `api.<Object>`. There is no separate `Factory` aggregate function.
- Completeness is **not** checked by the compiler. A field no factory's return value is assigned into stays nil and panics on the first call, so keeping `New` complete is the author's job. See [StructContracts.md](/docs/Explanations/StructContracts.md).

#### No Methods in Place of Factories
- A field is never filled by binding a method of the carrier. Methods may exist only as unexported helpers a closure calls.
- There is no internal mirror type and no `Api()` projection anywhere in the project.

## Structure
1. **Carrier struct**: the state holder — an api struct in [sandbox/contracts/api](/sandbox/contracts/api/).
2. **Field factories**: `func <Field>Factory(c *<Carrier>) <FieldType>`, one per field, each returning one closure.
3. **Doc comment**: one sentence per factory, naming the field it fills and what the returned value does.
4. **`New` constructor**: builds the carrier, calls every field factory exactly once and assigns its return value into the matching field, and returns the filled struct by value.

> **Note**: For a concrete example, refer to [sample.go](./sample.go).
