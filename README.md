<a href="https://pkg.go.dev/github.com/krhubert/gotest"><img src="https://img.shields.io/badge/godoc-reference-%23007d9c.svg"></a>

# gotest

Wrappers for functions, in standard go packages, that return an error.
The main purpose is to reduce boilerplate for error handling in tests, and make tests more readable.

The package layout is exactly the same as the standard go packages, with the same function names, but with the suffix `test` added to the package name.

## Usage

```go
package main_test

import (
    "testing"

    "github.com/krhubert/gotest/go/ostest"
)

func TestOpenFile(t *testing.T) {
    f := ostest.OpenFile(t, "file.txt")
    defer f.Close()
}
```
