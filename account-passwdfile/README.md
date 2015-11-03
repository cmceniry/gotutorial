passwdfile
==============

This exercise shows how to do some simple parsing of a formatted file,
and select part of it based on ignoring comments and on username
fields.

The goal is to display all passwd entries for users whose username is
some variant of daemon. There are several libraries necessary to do
this.

The bufio library will read in the file:

  https://golang.org/pkg/bufio

The strings library will help to remove comments and pull out the
appropriate field in passwd:

  https://golang.org/pkg/strings

The regexp library will match the daemon variant:

  https://golang.org/pkg/regexp

Focus - Tutorial Section
------------------------
slice
struct

Prerequisite - Tutorial Section
-------------------------------
gettingstarted
var
if
err
panic
for

Setup
-----
None
