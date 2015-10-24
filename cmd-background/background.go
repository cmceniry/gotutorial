package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// Setup a Command
	cmd := ...

	// Begin the Command in the background
	err := ...
	if err != nil {
		panic(err)
	}

	// Watch for it to end
	fmt.Println("Watching for sleep to end")
	err = ...
	if err != nil {
		panic(err)
	}
	fmt.Println("Complete")
}
