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
	cmd := exec.Command("/bin/ls", "-l", os.Args[1])

	// Connect os.Stdout to the Command
	cmd.Stdout = os.Stdout

	// Start the Command
	err := cmd.Run()
	if err != nil {
		panic(err)
	}

}
