package main

import (
	"fmt"
	"os"

	verblib "github.com/MateusMoutinhoOrg/Verb/sandbox"
)

func main() {
	// Build the lib directly from the real process argv.
	lib := verblib.New(os.Args[1:])

	// GetKeyValuesSize counts every "username=<value>" / "user=<value>"
	// argument without consuming anything.
	size := lib.GetKeyValuesSize([]string{"username=", "user="})
	for i := 0; i < size; i++ {
		// GetStringKeyValues marks the i-th matching argument as used and
		// returns the text after "=", or an error if the value is empty.
		currentUsername, err := lib.GetStringKeyValues([]string{"username=", "user="}, i)
		if err != nil {
			fmt.Println("username error:", err)
			continue
		}
		fmt.Println("username:", currentUsername)
	}
}
