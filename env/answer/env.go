package main

import (
  "os"
  "fmt"
)

func main() {
  // Obtain the entire environment as a slice
  env := os.Environ()
  // Iterate over env to output
  //   #     varname=value
  for i, _ := range env {
    fmt.Printf("%-5d %s\n", i, env[i])
  }
}
