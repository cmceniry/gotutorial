HTTP Client
===========

Go has a strong builtin HTTP library. You can interact with it on the
small scale with straight HTTP method calls, or you can use an client
which sets up an recurring environment (think cookies) and options.

This exercise is two parts:

  1. Get your internet facing IP by fetching from an api
  2. Attempt to access something inaccessible, but recover with a short
    timeout

You'll be using the net/http functions:

  https://golang.org/pkg/log/syslog/

With how net/http presents the response body, you'll also be using:

  https://golang.org/pkg/io/ioutil/

Tutorial Sections
-----------------
gettingstarted
var
if
err
panic
struct
defer


Setup
-----
None
