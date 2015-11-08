package main

import (
  "fmt"
  "io"
  "io/ioutil"
  "strings"
)

func main() {
  data, err := ioutil.ReadFile("sample.dat")
  if err != nil {
    panic(err)
  }

  // The data format is a string followed by 5 integers
  dformat := "%s %d %d %d %d %d"
  // temporary locations for holding data
  tmpname := ""
  tmpdata := [5]int{}

  // Split the data by lines, and process it
  for _, line := range strings.Split(string(data), "\n") {
    // Sscanf the line and load it into the tmpvalues
    n, err := fmt.Sscanf(line, dformat, &tmpname, &tmpdata[0], &tmpdata[1], &tmpdata[2], &tmpdata[3], &tmpdata[4])
    // panic if there are io errors (which aren't EOF)
    if err != nil && err != io.EOF {
      panic(err)
    }
    // If we don't scan enough fields, assume that we're at the end of
    // file and break out
    if n != 6 {
      break
    }

    // Compute the average and display with the dataset name
    total := 0
    for _, v := range tmpdata {
      total += v
    }
    fmt.Printf("%s : avg=%d\n", tmpname, total/5)
  }

}
