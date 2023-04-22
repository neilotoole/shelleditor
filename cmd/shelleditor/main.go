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
