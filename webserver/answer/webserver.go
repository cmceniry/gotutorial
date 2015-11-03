package main

import (
	"fmt"
	"net/http"
	"runtime"
)

// A simple function which follows the required signature
// to be used as a web server handler in net/http
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World\n\n")
}

func main() {
	//
	http.HandleFunc("/", hello)

	// Use the same function as above, but do it inline
	http.HandleFunc("/hey", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hey World\n\n")
	})

	// Make a more complex function that returns the stack trace
	http.HandleFunc("/go", func(w http.ResponseWriter, r *http.Request) {
		var buf []byte = make([]byte, 10240)
		n := runtime.Stack(buf, true)
		fmt.Fprintf(w, "%d\n", n)
		fmt.Fprintf(w, "%s\n", string(buf))
	})

	// Start up the web server and server forever
	http.ListenAndServe(":8000", nil)
}
