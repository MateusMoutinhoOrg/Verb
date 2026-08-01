//go:build ignore

// This file is an illustrative sample, not part of the build.
package main

import (
	verbadapter "github.com/MateusMoutinhoOrg/Verb/adapters/standard"
	verblib "github.com/MateusMoutinhoOrg/Verb/sandbox"
)

func main() {
	// 1. Build deps through an adapter (the opinionated layer).
	deps := verbadapter.New("data.json")

	// 2. Inject deps into the pure library.
	l := verblib.New(deps)

	// 3. Exercise the library — it never knows which adapter is behind it.
	//    Its functions are struct fields, called like any method.
	obj := l.NewExampleObject(1, "2")
	println(obj.ExampleObjectMethod())
}
