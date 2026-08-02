# Verb

[![Go Reference](https://pkg.go.dev/badge/github.com/MateusMoutinhoOrg/Verb.svg)](https://pkg.go.dev/github.com/MateusMoutinhoOrg/Verb)
[![Release](https://img.shields.io/github/v/release/MateusMoutinhoOrg/Verb)](https://github.com/MateusMoutinhoOrg/Verb/releases/latest)
[![Go Version](https://img.shields.io/badge/go-%3E%3D1.22-blue)](go.mod)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)

An OS-independent Go library template demonstrating a **struct-of-functions** public API with a clean separation between the pure library and its consumers.

---

## Overview

Verb is a structured Go template that showcases how to build libraries exposed as plain data instead of interfaces. It uses a pattern in which:

- **`/sandbox/contracts/`** defines the `api` structs the library hands back.
- **`/sandbox/internal/`** contains the pure library logic as factories filling the `api` contract structs.
- **`/sandbox/`** is the entry point: it takes plain arguments and returns an `api.Lib`.

This design ensures the library remains portable, testable, and easy to extend without modifying its core.

---

## Quick Start

**1. Install the library:**
```bash
go get github.com/MateusMoutinhoOrg/Verb
```

**2. Create a `main.go` file:**
```go
package main

import (
    verblib "github.com/MateusMoutinhoOrg/Verb/sandbox"
)

func main() {
    // 1. Build the library from plain arguments
    l := verblib.New(3)

    // 2. Use the library
    obj := l.NewExampleObject(1, "2")
    println(obj.ExampleObjectMethod())
}
```

**3. Run:**
```bash
go run main.go
```

---

> [!IMPORTANT]
> **Must Read before contributing.** The following documents are **required reading** for every developer. Do not open a pull request or make changes without first reading them:
>
> | Document | Why it's required |
> |----------|-------------------|
> | [Rules](/docs/References/RULES.md) | The contribution rules and guidelines that **must** be followed for any change to be accepted. |
> | [Structure](/docs/References/Structure.md) | The project's directory layout and the purpose of each component — needed to know **where** changes belong. |
> | [Specs](/docs/References/Specs.md) | The index of every specification — needed to know **how** the file you are about to touch must be shaped. |

## Library Usage

For consuming the lib as a user: install it, run a first program, and understand its core mechanic.

| Doc | Description | Type |
| --- | --- | --- |
| [/docs/Tutorials/LibInitialization.md](/docs/Tutorials/LibInitialization.md) | Install the lib, call `lib.New`, and run a first program | Tutorial |
| [/docs/Tutorials/RunSample.md](/docs/Tutorials/RunSample.md) | Browse and run the executable samples in the examples/ directory | Tutorial |
| [/docs/Explanations/StructContracts.md](/docs/Explanations/StructContracts.md) | How the struct-of-functions public API works | Explanation |

---

## Samples

Creating and running the example programs under `examples/`.

| Doc | Description | Type |
| --- | --- | --- |
| [/docs/Tutorials/RunSample.md](/docs/Tutorials/RunSample.md) | Browse and run the executable samples in the examples/ directory | Tutorial |
| [/docs/Tutorials/AddSample.md](/docs/Tutorials/AddSample.md) | Create a runnable sample in examples/ and register it in the README | Tutorial |

### Available Samples

| Sample | Description |
|----------|-------------|
| [ExampleSample](/examples/ExampleSample/ExampleSample.go) | How to use the library |

---

## Project Rules & Structure

The binding conventions every change to this repo must follow.

| Doc | Description | Type |
| --- | --- | --- |
| [/docs/References/RULES.md](/docs/References/RULES.md) | The binding contribution rules and their required companion updates | Reference |
| [/docs/References/Structure.md](/docs/References/Structure.md) | The project's directory layout and the purpose of each component | Reference |
| [/docs/References/Specs.md](/docs/References/Specs.md) | Lists every specification and the files each one governs | Reference |

---

## License

This project is licensed under the [MIT License](./LICENSE).
