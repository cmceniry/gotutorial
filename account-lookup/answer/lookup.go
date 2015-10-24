package main

import (
	"fmt"
	"os/user"
)

func main() {
	fmt.Println("Current User")
	// Lookup current user, and print out info
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Println(u)

	fmt.Println("root User by name")
	// Lookup root user by name, and print out info
	username, err := user.Lookup("root")
	if err != nil {
		panic(err)
	}
	fmt.Println(username)

	fmt.Println("root User by UID 0")
	// Lookup root user by user name, and print out info
	// Careful here - as the type is not what you expect it to be
	u, err = user.LookupId("0")
	if err != nil {
		panic(err)
	}
	fmt.Println(u)

}
