# Struct-of-Functions Output

## Description
Explains how the library exposes its behavior as a struct of function fields instead of an interface, and how those fields get filled.

---

## The Output Struct

`api.Example` is a plain struct whose fields are functions — not an interface. A **factory** fills each field with a closure that reads the struct's other fields at call time:

```go
type Example struct {
    Count int
    Double func() int // each field is one exposed behavior
}
```

Because a field is a plain function, calling it reads exactly like calling a method.

---

## Filling the Fields

A factory takes a pointer to the struct being built and returns the closure for one field; the package's `New` constructor assigns every factory's return value:

```go
func DoubleFactory(e *Example) func() int {
    return func() int { return e.Count * 2 }
}

func New(count int) Example {
    e := Example{Count: count}
    e.Double = DoubleFactory(&e)
    return e
}
```

The closure reads `e.Count` through the pointer, so it always sees the struct's current state, not a value captured when the factory ran.
