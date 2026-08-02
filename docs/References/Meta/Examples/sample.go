//go:build ignore

// This file is an illustrative sample, not part of the build.
package main

import (
	verblib "github.com/MateusMoutinhoOrg/Verb/sandbox"
)

func main() {
	// 1. Build the library directly from the arguments it needs.
	l := verblib.New("data.json")

	// 2. Exercise the library — its functions are struct fields, called
	//    like any method.
	obj := l.NewExampleObject(1, "2")
	println(obj.ExampleObjectMethod())
}
