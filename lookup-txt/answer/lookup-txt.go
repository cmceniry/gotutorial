package main

import (
	"fmt"
	"net"
)

// MustResolveOneIPByName takes in a string and returns the single
// string value for the IP which is being resolved.
// If it fails, or returns other values, panic
func MustResolveOneTXTByName(name string) string {
	// Do the resolve
	t, err := net.LookupTXT(name)
	// Check the initial error and panic if need be
	if err != nil {
		panic(err)
	}
	// Check for multiple return values
	if len(t) != 1 {
		panic(fmt.Errorf("Invalid response: %v", t))
	}
	// Return the single value
	return t[0]
}

// ResolveOneIPByName takes in a string and returns a string and error
// indicator. The error indicator is used if the resolve fails or returns
// multiple values.
func ResolveOneTXTByName(name string) (string, error) {
	// Do the resolve
	t, err := net.LookupTXT(name)
	// Check the initial error and return that if present
	if err != nil {
		return "", err
	}
	// Check for multiple IPs. If so, send back an error w/ "Multiple IPs returned"
	if len(t) != 1 {
		err = fmt.Errorf("Multiple TXT records returned")
		return "", err
	}
	// Return the single value
	return t[0], nil
}

func main() {
 	t1 := MustResolveOneTXTByName("example.com")
	fmt.Println(t1)
	t2, err := ResolveOneTXTByName("example.com")
	if err != nil {
		fmt.Println("ResolveOne returned error: ", err)
	} else {
		fmt.Println(t2)
	}
}
