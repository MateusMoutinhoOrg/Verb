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

	// GetOptionsSize counts every "-username <value>" / "--username <value>"
	// occurrence without consuming anything.
	size := lib.GetOptionsSize([]string{"-username", "--username"})
	for i := 0; i < size; i++ {
		// GetStringOption marks the i-th "--username" and its following
		// value as used, returning the value or an error if it is missing.
		currentUsername, err := lib.GetStringOption([]string{"-username", "--username"}, i)
		if err != nil {
			fmt.Println("username error:", err)
			continue
		}
		fmt.Println("username:", currentUsername)
	}
}
