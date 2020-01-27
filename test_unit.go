package lab

import (
	"runtime"
	"testing"
)

// TestUnit represents a single test in `go test`
type TestUnit struct {
	t         *testing.T
	assertion Assertion
}

// NewTestUnit prepares a new test unit to test code
func NewTestUnit(t *testing.T) *TestUnit {
	unit := &TestUnit{}
	unit.t = t
	unit.assertion = Assertion(unit.Assert)
	return unit
}

// Context opens a new context for this test unit
func (unit *TestUnit) Context(args ...interface{}) {
}

// Test opens a new test section for this test unit
func (unit *TestUnit) Test(args ...interface{}) {
}

// Assertion provides a new assertion context
func (unit *TestUnit) Assertion() Assertion {
	return unit.assertion
}

// Assert tests the result is successful (true)
func (unit *TestUnit) Assert(args ...interface{}) {
	assertOk := false
	msg := "Assertion failed"
	ok := false
	skip := 1

	switch len(args) {
	case 1:
		assertOk, ok = args[0].(bool)
		if !ok {
			panic(&ArgumentError{"args[0]", "is not a bool"})
		}
	case 2:
		assertOk, ok = args[0].(bool)
		if !ok {
			panic(&ArgumentError{"args[0]", "is not a bool"})
		}

		msg, ok = args[1].(string)
		if !ok {
			panic(&ArgumentError{"args[1]", "is not a string"})
		}
	case 3:
		assertOk, ok = args[0].(bool)
		if !ok {
			panic(&ArgumentError{"args[0]", "is not a bool"})
		}

		msg, ok = args[1].(string)
		if !ok {
			panic(&ArgumentError{"args[1]", "is not a string"})
		}

		skip, ok = args[2].(int)
		if !ok {
			panic(&ArgumentError{"args[2]", "is not an int"})
		}
	default:
		panic(&ArgumentError{"args", "invalid amount"})
	}

	if !assertOk {
		if _, file, line, ok := runtime.Caller(skip); ok {
			panic(&AssertionError{Msg: msg, File: file, Line: line})
		}

		panic(&AssertionError{Msg: msg})
	}
}
