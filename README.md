# BFI - Brainf*ck interpreter in Go

[![Build Status](https://travis-ci.org/ibraimgm/bfi.svg?branch=master)](https://travis-ci.org/ibraimgm/bfi)
[![codecov](https://codecov.io/gh/ibraimgm/bfi/branch/master/graph/badge.svg)](https://codecov.io/gh/ibraimgm/bfi)
[![Go Report Card](https://goreportcard.com/badge/github.com/ibraimgm/bfi)](https://goreportcard.com/report/github.com/ibraimgm/bfi)
[![BFI Docs](https://img.shields.io/badge/godoc-api-blue.svg)](https://godoc.org/github.com/ibraimgm/bfi)

This is my take on a simple brainf*ck interpreter in Go.
The interpreter is very simple and straightforward to use, and the main application (main.go) serves more asa demo on how to
use the interpreter API.

If you want, you can check how it works using one of the files available in the `samples` directory (or just writing your own brainf*ck code).

The interpreter API (and the command line) allows you to change the cell size and the tape length (try the `--help` option).
However, most brainf*ck code assumes a cell size of 8 bits and a tape of at least 3000 cells (if no option is specified, these are the defaults used
in the interpreter).

## License

See [LICENSE](LICENSE) for details.
