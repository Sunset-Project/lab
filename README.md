```go
func TestSomething(test *testing.T) {
  t := lab.New(test)
	t.Context("Foo", func() {
    t.Context("Bar", func() {
      t.Context("Baz", func() {
        t.Test("Something", func() {
          t.Assert(1 == 1)
          t.Assert(2 == 2)
        })
      })
    })
  })
}


func TestSomething(test *testing.T) {
  Context, Test := lab.New(test)
	t.Context("Foo", func() {
    t.Context("Bar", func() {
      t.Context("Baz", func() {
        t.Test("Something", func() {
          t.Assert(1 == 1)
          t.Assert(2 == 2)
        })
      })
    })
  })
}
```

# Functions
- lab.BuildTest
- Assert
- Test
- Context
- Assert.Panic
- Assert.Ok - `_, ok`

# Objects

## Session
List of all ongoing tests.
Synchronizes the output for one test at a time
Global (synchronization point)

## Output
Formatted printing of test proses

## Test
Test contexts and data for a single test file (test function)
