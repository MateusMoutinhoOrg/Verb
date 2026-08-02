package main

import (
	"fmt"
	"os"

	verblib "github.com/MateusMoutinhoOrg/Verb/sandbox"
)

func main() {
	// Build the lib directly from the real process argv.
	lib := verblib.New(os.Args[1:])

	// IsPresent marks the matching flag as used and reports whether any of
	// its spellings occurred in the argv, e.g. `./app -f` or `./app --file`.
	forceExists := lib.IsPresent([]string{"-f", "--f", "-file", "--file"})

	fmt.Println("force:", forceExists)
}
