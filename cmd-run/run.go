package main

import (
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) != 2 {
		panic("Need to specify path")
	}

	// Setup a Command
	cmd := ...

	// Connect os.Stdout to the Command
	... = os.Stdout

	// Begin the Command
	err := ...
	if err != nil {
		panic(err)
	}

}
