# shelleditor
Invoke the shell `EDITOR` like `kubectl edit` does.

This code is lifted from [kubectl](https://github.com/kubernetes/kubectl/tree/master/pkg/cmd/util/editor).
Thanks lads. 

Why not import the `kubectl` code directly? It has tons of dependencies that are not needed
for this simple task. The codebase has been edited to import fewer packages,
and those that are imported are mostly copied to the `/pkg` dir.

## Usage

Use the normal mechanism:

```shell
go get -u github.com/neilotoole/shelleditor
```

## Example program

There's an example program in [`cmd/shelleditor`](https://github.com/neilotoole/shelleditor/blob/master/cmd/shelleditor/main.go).

```shell
$ go install github.com/neilotoole/shelleditor/cmd/shelleditor
$ shelleditor hello.txt
```

It's very simple:

```go
package main

import (
	"fmt"
	"os"

	"golang.org/x/exp/slog"

	"github.com/neilotoole/shelleditor"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: shelleditor PATH")
		os.Exit(1)
	}

	// Set logger... you can usually ignore this. When not
	// set, log output is discarded.
	shelleditor.SetLogger(slog.Default())

	ed := shelleditor.NewDefaultEditor("EDITOR")
	if err := ed.Launch(os.Args[1]); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
```

## Alternatives

- [th/go-editor](https://github.com/tj/go-editor)
