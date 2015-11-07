package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Create a data struct type with two string fields: Name and Value
type Section struct {
	...
	...
}

func main() {
	// To allow for arbitrary subsections, create map with a string
	// key type and has the above Section type for its element type
	var cnf ...

	// Read in the entire JSON file
	data, err := ioutil.ReadFile("sample1.json")
	if err != nil {
		panic(err)
	}

	// Pull the data out using the json functions. For now, accept the
	// use of & below
	err = ...(data, &cnf)
	if err != nil {
		panic(err)
	}

	// Print the config values out in an ini style format by
	// iterating through each section name and section data
	for sname, section := range cnf {
		fmt.Printf("[%s]\n", sname)
		fmt.Printf("name = %s\n", ...)
		fmt.Printf("value = %s\n", ...)
	}
}
