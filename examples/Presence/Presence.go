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

	// IsPresent marks the matching flag as used and reports whether any of
	// its spellings occurred in the argv, e.g. `./app -f` or `./app --file`.
	forceExists := lib.IsPresent([]string{"-f", "--f", "-file", "--file"})

	fmt.Println("force:", forceExists)
}
