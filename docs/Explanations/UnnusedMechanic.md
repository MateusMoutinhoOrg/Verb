
The Unnused Mechanic 
the unnused mechanic marks every iten of argvlist as readed,
its ussefull to make complete command.

Example:
```bash
./cli  -o teste.out --quiet teste.c 
```
in a cli like that will have theses values: 

```go

package main

import (
	"os"

	verbadapter "github.com/MateusMoutinhoOrg/Verb/adapters/standard"
	verblib "github.com/MateusMoutinhoOrg/Verb/sandbox"
)

func main() {
	// Build deps via an adapter (JSON file store + real clock) and inject.
	lib := verblib.New(verbadapter.New(os.Args))

	// marks -o and teste.out as used 
	output = lib.getStringOption([]string{"-o","--o","--output","--out"},0)

    // marks -q and --q and --quiet and --quiet as used 
    quiet := lib.isPresent([]string{"-q","--q","--quiet","--quiet"})
    
   // gets teste.c since is the first unnused iten
    file := lib.getNextStringArg()



}

```


