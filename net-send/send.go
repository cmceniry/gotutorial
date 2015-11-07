package main

import (
	"fmt"
	"net"
)

func main() {
	// Open up the connection to the remove side - localhost:4270
	s, err := ...

	// Regardless of what happens, we want to clean up
  ...

	// Check to see if there were any errors setting up
	if ... {
		// If so, crash and display the error
		...
	}

	// Now send a test string over "This is a test\n"
	fmt.Fprintf(s, "This is a test\n")
}
