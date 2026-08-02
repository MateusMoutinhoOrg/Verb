# `api.Lib.Used`

**Type:** Field

## Signature

```go
Used []bool
```

## Description

Tracks, index for index against [`Args`](/docs/References/PublicApi/api.Args.md), which arguments have already been matched by a previous call. `Used[i]` becomes `true` once `Args[i]` is consumed by any `Get*` function or by `IsPresent`; it starts entirely `false` and only grows more `true` over the `Lib`'s lifetime. `GetNextStringArg` and its typed variants scan this slice to find the next never-consumed positional argument — see [UnnusedMechanic.md](/docs/Explanations/UnnusedMechanic.md).

Exported for the same cross-package reason as `Args`; treat it as read-only from outside `sandbox/internal/lib`.

## Returns

| Type | Description |
| :--- | :--- |
| `[]bool` | Read-state for every position in `Args`. |
