package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Perform the get for: https://api.ipify.org?format=text
	...
	if err != nil {
		panic(err)
	}
	// Inside of the get, the http library opens the response Body. We
	// want to make sure that we release the open Body when we're done
	// no matter what happens with out execution (though in this simple
  // example it's not as big of a deal)
	...

	// Access the body and read it all in, and print it out
	...
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
