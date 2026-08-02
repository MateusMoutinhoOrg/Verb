# `api.Lib.Args`

**Type:** Field

## Signature

```go
Args []string
```

## Description

The snapshot of the argument vector being parsed, taken once from `Deps.Args()` when [`lib.New`](/docs/References/PublicApi/lib.New.md) built this `Lib`. Every index-based function field (`GetStringArg`, `GetStringOption`, `GetStringKeyValues`, and their typed variants) refers to positions in this slice — index `0` is the first argument after the program name.

Exported so `sandbox/internal/lib` can populate it from another package; treat it as read-only from outside the library. Mutating it after construction desynchronizes it from [`Used`](/docs/References/PublicApi/api.Used.md) and produces undefined matching behavior.

## Returns

| Type | Description |
| :--- | :--- |
| `[]string` | The argv this `Lib` parses, fixed for its whole lifetime. |
