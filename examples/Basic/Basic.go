package main

import (
	"os"

	verbadapter "github.com/MateusMoutinhoOrg/Verb/adapters/standard"
	verblib "github.com/MateusMoutinhoOrg/Verb/sandbox"
)

func main() {
	// Build deps via an adapter (JSON file store + real clock) and inject.
	lib := verblib.New(verbadapter.New(os.Args))

	lib.setActions([]verblib.Action{
		verblib.Action{
			Triggers: []verblib.Trigger{

				verblib.ArgTrigger{

					ExpectedValue:"commit",
					ExpectedPosition: 1,
				}
			},

		},
	})
}
