package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var r = regexp.MustCompile("(?i)daemon")

func main() {
	f, err := os.Open("/etc/passwd")
	if err != nil {
		panic(err)
	}
	// Create a scanner var which uses the bufio Scanner type
	// to wrap the file f
	scanner := ...
	// Iterate on the scanner
	for ... {
		// Pull out the current scan, and split it apart
		commentless := strings.SplitAfter(..., "#")
		splitline := strings.Split(..., ":")
		// If the length doesn't match 7 - a valid passwd format
		// skip to the next line
		if ... {
			continue
		}
		// Check to see if gecos name field (5th field) of the
		// password file matches our regexp and print it out
		if r.MatchString(...) {
			fmt.Println(...)
		}
	}
	// Check for errors
	if err = ...; err != nil {
		panic(err)
	}

}
