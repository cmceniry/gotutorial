tar
===

This exercise shows how to work with a tar file, and how to wrap it in gzip.
You'll be using the built in tar and gzip libraries:

  https://golang.org/pkg/archive/tar/
  https://golang.org/pkg/compress/gzip/

This will create two sample files - one just tar and one tar wrapped in
gzip. You can confirm the output of this by checking the files by hand:

  tar tvf sample.tar
  tar ztvf sample.tgz

This exercise will compare how to use an interface to create one function
that can write to straight files or to gzipped files.

Focus - Tutorial Section
------------------------
interface

Prerequisite - Tutorial Section
-------------------------------
gettingstarted
if
err
panic
for
struct
pointers
func

Setup
-----
None
