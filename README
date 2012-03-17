ABOUT
=====

mc_fcsh is a wrapper server/client for fcsh written in Go

The fcsh tool is much faster than the standard Flex command line compiler
(mxmlc).  It's a bit rubbish for command line use, so mc_fcsh aims to solve
that problem.

It's probably not ready for public consumption.  Feel free to try it, but
don't expect wonders.

USAGE
=====

Run the server dj_fcsh in the background somewhere.
Use the command line mc_fcsh like mxmlc.

The truly brave can submit build requests via POST to http://localhost:7950/compile

WHY GO?
=======
* I wanted to learn it.
* It's quite good this sort of thing.
* It's cross-platform.

LIMITATIONS
===========
* Expects first argument to be you mxmlc file.
* Will not handle compiles with multiple referenced files well.
* No clean shutdown support.
* Nearly completely untested (woo!)

CREDITS
=======

Written by Jonathan Whiting

Concept is stolen from https://github.com/Draknek/fcsh-wrap/ which is great, but
lacks windows support.
