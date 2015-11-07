package main

import (
	"fmt"
	"net"
)

func main() {
	// Open up the connection to the remove side - localhost:4270
	s, err := net.Dial("tcp", "localhost:4270")

	// Regardless of what happens, we want to clean up
	defer s.Close()

	// Check to see if there were any errors setting up
	if err != nil {
		// If so, crash and display the error
		panic(err)
	}

	// Now send a test string over "This is a test\n"
	fmt.Fprintf(s, "This is a test\n")
}
