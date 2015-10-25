package main

import (
	"log"
	"log/syslog"
)

func main() {
	// Create a new syslog container to send to notice/local0
	// and include the file shortname
	// Setup to log to notice/local0
	// Log using file short name
	...
	if err != nil {
		panic(err)
	}

	// Log "GOTUTORIAL" message
	...
}
