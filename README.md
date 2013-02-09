Simple HTTP Server
==================

If you need a quick "hello server" like

    python -m SimpleHTTPServer

and be able to serve more requests...


1. Compile
----------

    go build

2. Run it
---------

    ./simplehttpserver

It does not log anything for the moment (feel free to contribute). See the
default options below.

3. Modify the behavior
----------------------

    ./simplehttpserver -h
    Usage of ./simplehttpserver:
      -bind=":8000": Network address for listening
      -text="ok": Text to reply
