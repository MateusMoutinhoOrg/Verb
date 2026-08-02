# Adapters

## Description
Lists every adapter shipped with the library — the opinionated `deps.Deps` implementations under `adapters/` — and when to use each one. Every adapter exposes a `New(...) deps.Deps` factory that runs one `<Field>Factory` per field of the contract and returns the filled contract struct, ready to be passed to [`lib.New`](/docs/References/PublicApi/lib.New.md) — the same [factory pattern](/docs/References/RULES.md#factory-pattern) the sandbox uses. Any single field can be replaced before injection — see [DepsMechanic.md](/docs/Explanations/DepsMechanic.md). To build a new adapter, follow [AddAdapter.md](/docs/Tutorials/AddAdapter.md).

---

## Available Adapters

| Adapter | Factory | Behavior | Use When |
|---------|---------|----------|----------|
| `standard` | [standard.New](/docs/References/PublicApi/standard.New.md) | Hands back the argv slice it was constructed with — pass `os.Args[1:]` for the real process command line | You want to parse the program's actual command-line arguments |
| `memory` | [memory.New](/docs/References/PublicApi/memory.New.md) | Hands back a fixed, caller-supplied `[]string` regardless of the real process argv | You're writing a test or script and want a known, repeatable argv |
