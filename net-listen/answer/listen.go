package main

import (
	"io"
	"net"
	"os"
)

// HandleConnection does the work of processing the connections once
// they are established. This leaves the main thread to be able to
// continue to listen for new connections and hand off instead of
// blocking.
func HandleConnection(c net.Conn) {
	// Copy all of the input from the connection to stdout
	io.Copy(os.Stdout, c)
	// Kill the connection once we're done copying
	c.Close()
}

func main() {
	// Create a lsitener on localhost:4270/tcp
	l, err := net.Listen("tcp", ":4270")
	// Confirm that we listened without a problem
	if err != nil {
		panic(err)
	}

	// So main doesn't exit after the first connection, start a loop
	for {
		// Wait for a new connection. Bail out if any error crops up
		c, err := l.Accept()
		if err != nil {
			panic(err)
		}
		// Hand off to HandleConnection
		go HandleConnection(c)
	}
}
