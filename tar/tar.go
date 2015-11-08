package main

import (
  "archive/tar"
  "compress/gzip"
  "os"
  ...
  "fmt"
)

// MakeTar creates a test tarball on top of any underlying
// construct that can be written to
func MakeTar(w ...) {
  scratch := []byte("this is a test\n")

  // Create the writer for tar to use
  t := ...
  defer t.Close()
  // Initialize some header values
  th1 := ...{Name: "tar-first", Mode: 0700, Size: 75}
  th2 := ...{Name: "tar-second", Mode: 0700, Size: 90}

  // Write the first file
  ...(th1)
  for x := 0; x < 5; x += 1 {
    t.Write(scratch)
  }
  // Write the second file
  ...(th2)
  for x := 0; x < 6; x += 1 {
    t.Write(scratch)
  }
  t.Close()
}

func main() {
  // Setup the raw file structure to write to
  f, err := os.OpenFile("sample.tar", os.O_RDWR|os.O_CREATE, 0700)
  if err != nil {
    panic(err)
  }
  defer f.Close()

  // Show the type for the raw file
  fmt.Printf("%#v\n", f)

  // Write to raw file
  MakeTar(f)

  // Setup the raw file structure to write the gzip data to
  f, err = os.OpenFile("sample.tgz", os.O_RDWR|os.O_CREATE, 0700)
  if err != nil {
    panic(err)
  }
  defer f.Close()

  // Wrap the file in a gzip writer
  g := gzip.NewWriter(f)
  defer g.Close()

  // Show the type for the gzipped file
  fmt.Printf("%#v\n", g)

  // Write to the gzip interface to write to the file
  MakeTar(g)
}
