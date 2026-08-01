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

					ExpectedValue: "commit",
					//will be triggered if has commit word in args
					// from  1 to 1 means that both argv[0] and argv[1] can be commit
					PlotageArea: []uint{1, 1},
				},
			},
			Entries: []verblib.Entry{
				verblib.FlagEntry{
					Name:        "message",
					Identifiers: []string{"-m", "--m", "-message", "--message"},
					Required:    true,
				},
			},
		},
	})
}

// test by
// go run examples/Triggers/Triggers.go  commit
