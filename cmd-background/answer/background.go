package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// Setup a Command
	cmd := exec.Command("/bin/sleep", "5")

	// Begin the Command in the background
	err := cmd.Start()
	if err != nil {
		panic(err)
	}

	// Watch for it to end
	fmt.Println("Watching for sleep to end")
	err = cmd.Wait()
	if err != nil {
		panic(err)
	}
	fmt.Println("Complete")
}
