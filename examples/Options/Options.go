package main

import (
	"os"

	verbadapter "github.com/MateusMoutinhoOrg/Verb/adapters/standard"
	verblib "github.com/MateusMoutinhoOrg/Verb/sandbox"
)

func main() {
	// Build deps via an adapter (JSON file store + real clock) and inject.
	lib := verblib.New(verbadapter.New(os.Args))

	size := lib.GetOptionsSize([]string{"-username", "--username"})
	for i := 0; i < size; i++ {

		//will get all --username <value> , if a value is not present , will raize a error
		current_username := lib.GetStringOption([]string{"-username", "--username"}, i)
		println("username: ", current_username)
	}

	//marks --username and <value>  as used args

}
