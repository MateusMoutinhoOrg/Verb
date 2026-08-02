package main

import (
	"fmt"
	"os"

	verbadapter "github.com/MateusMoutinhoOrg/Verb/adapters/standard"
	verblib "github.com/MateusMoutinhoOrg/Verb/sandbox"
)

func main() {
	// Build deps via an adapter (the real process argv) and inject.
	lib := verblib.New(verbadapter.New(os.Args[1:]))

	// GetStringArg reads by absolute position: `./app test` -> index 0 is
	// "test", regardless of what else is on the command line.
	first, err := lib.GetStringArg(0)
	if err != nil {
		fmt.Println("arg 0 error:", err)
		return
	}
	fmt.Println("first arg:", first)
}
