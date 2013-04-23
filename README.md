## Emacs flycheck-mode for the Go programming language

The `goflycheck` program is a wrapper around the `go` tool to provide
Emacs flychcek style syntax checking for golang source files within
multi-file packages and _test.go files.  Support for os/arch specific
*cgo* files is included thanks to the standard *go/build* package.

Based on the [goflymake][1] tool by [dougm][2].

[1]: https://github.com/dougm/goflymake
[2]: https://github.com/dougm

### Setup

 1. If needed, update your **$PATH** to include go installed binaries, for example:

    `export PATH=$PATH:$GOPATH/bin`

 2. Install goflymake:

    `go get -u github.com/ptrv/goflycheck`

### Emacs setup

 1. Install go-mode.el if you haven't already

 2. Add these lines to your **.emacs** or similar:

        (add-to-list 'load-path "~/gocode/src/github.com/ptrv/goflycheck")
        (require 'go-flycheck)
