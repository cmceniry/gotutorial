package main

import (
	"log"
)

func main() {
	// Println "Log1" with log with the default options
	log.Println("Log1")

	// Log "Log2" with a "GOTUTORIAL: " prefix
	log.SetPrefix("GOTUTORIAL: ")
	log.Println("Log2")

	// Log "Log3" with "GOTUTORIAL: " prefix
	// and include source file line with the default date and time
	log.SetFlags(log.LstdFlags | log.Llongfile)
	log.Println("Log3")

	// Panic via logging
	log.Panic()
}
