package main

import (
	"fmt"
	"os"

	verblib "github.com/MateusMoutinhoOrg/Verb/sandbox"
)

func main() {
	// Build the lib directly from the real process argv.
	lib := verblib.New(os.Args[1:])

	// GetStringArg(0) reads the first argument by absolute position,
	// e.g. `./app test` -> first == "test".
	first, err := lib.GetStringArg(0)
	if err != nil {
		fmt.Println("arg 0 error:", err)
		return
	}
	fmt.Println("first arg:", first)

	// GetNextStringArg is the Unused Mechanic: it returns the next argument
	// that has not yet been read by any Get*/IsPresent call — see
	// docs/Explanations/UnnusedMechanic.md.
	second, err := lib.GetNextStringArg()
	if err != nil {
		fmt.Println("next arg error:", err)
		return
	}
	fmt.Println("next unused arg:", second)
}
