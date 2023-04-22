package main

import (
	"fmt"
	"os"

	"github.com/neilotoole/shelleditor"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: shelleditor PATH")
		os.Exit(1)
	}

	ed := shelleditor.NewDefaultEditor([]string{"EDITOR"})
	if err := ed.Launch(os.Args[1]); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
