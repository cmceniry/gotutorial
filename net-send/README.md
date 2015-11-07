net-send
========

This is a simple exercise which uses the network calls to send data to
a foreign endpoint. You'll be using the net functions:

  http://golang.org/pkg/net/

This will go over a bit of proper error handling. In addition, since
we're using network connections, we'll want to make sure we clean up
regardless of what happens.

Focus - Tutorial Section
------------------------
if
err
panic
defer

Prerequisite - Tutorial Section
-------------------------------
gettingstarted

Setup
-----
You'll need to set up a remote endpoint to connect to - on
localhost:4270/tcp. The easiest way to do this is with nc:

  nc -l 4270

Start that in one terminal, and work in the other.
