# Lab

Go test suite built on top of `go test`

## Defining a test

A `lab` test is a traditional `go test` with a special syntax to improve
the output.
The following is an example that can be used as a baseline:

```go
package apackage_test

import (
	"testing"

	"github.com/sunset-project/lab"
)

func TestSomething(t *testing.T) {
  // Starts a lab test session
  Context, Test, Assert := lab.StartSession(t)
  Context("A context", func() {
    Test("A test", func() {
      Assert(true)
      Assert.Panic(func() { panic("panic") })
      Assert.PanicMsg(
        func() { panic("panic") },
        func(data interface{}) {
          _, ok := data.(string)
          return ok
        }
      )
    })
  })
}
```

## Running tests

Define the following bash function

```bash
lab() {
  if [ $# -gt 0 ]; then
    go test -v -count=1 $@
    return $?
  else
    go test -v -count=1 './...'
    return $?
  fi
}
```

Tests can be executed from the project directory with

```bash
lab
```

To selectively execute a single test-file, the following command can be used:

```bash
lab a/test/file_test.go
```

## TODO

- IO Reporter
  - [x] Indentation should increase for each nested block
  - [x] Error printing should be colored
  - [x] Error printing should be indented
  - [ ] Failed assertion line should be printed in the output
  - [ ] Panic line should be printed in the output
  - [x] Output should be colored
  - [ ] Output coloring should be activated when it's a tty
- [ ] `lab` bash function should be an executable released as part of this
  package
- [ ] Configuration through environment variables
- [ ] Stack trace during panic should exclude non-relevant lines
- [ ] Remove `pkg/errors` dependency
- [ ] `Context.Skip`
- [ ] `Test.Skip`
