package lab

import (
	"runtime"
	"testing"
)

// TestUnit represents a single test in `go test`
type TestUnit struct {
	t         *testing.T
	assertion Assertion
	reporter  Reporter
}

// NewTestUnit prepares a new test unit to test code
func NewTestUnit(t *testing.T) *TestUnit {
	unit := &TestUnit{}
	unit.t = t
	unit.assertion = Assertion(unit.Assert)
	unit.reporter = DefaultReporter()
	return unit
}

// Context opens a new context for this test unit
func (unit *TestUnit) Context(args ...interface{}) {
	var prose string
	var do func()

	switch len(args) {
	case 1:
		prose = args[0].(string)
	case 2:
		prose = args[0].(string)
		do = args[1].(func())
	default:
		panic(&ArgumentError{"args", "invalid amount"})
	}

	// EnterContext
	unit.reporter.ContextEntered(prose)
	defer func() {
		// LeaveContext
		unit.reporter.ContextExited(prose)
	}()

	if do == nil {
		// SkipContext
		unit.reporter.ContextSkipped(prose)
	} else {
		defer func() {
			err := recover()

			if err == nil {
				// SuccessContext
				unit.reporter.ContextSucceeded(prose)
			} else {
				// Output error
				unit.reporter.PanicInvoked(err)
				// FailContext
				unit.reporter.ContextFailed(prose)
			}
		}()

		do()
	}
}

// Test opens a new test section for this test unit
func (unit *TestUnit) Test(args ...interface{}) {
	var prose string
	var do func()

	switch len(args) {
	case 1:
		prose = args[0].(string)
	case 2:
		prose = args[0].(string)
		do = args[1].(func())
	default:
		panic(&ArgumentError{"args", "invalid amount"})
	}

	// EnterTest
	unit.reporter.TestStarted(prose)
	defer func() {
		// LeaveTest
		unit.reporter.TestFinished(prose)
	}()

	if do == nil {
		// SkipTest
		unit.reporter.TestSkipped(prose)
	} else {
		defer func() {
			err := recover()

			if err == nil {
				// SuccessTest
				unit.reporter.TestPassed(prose)
			} else {
				// Output error
				unit.reporter.PanicInvoked(err)
				// FailTest
				unit.reporter.TestFailed(prose)
			}
		}()

		do()
	}
}

// Assertion provides a new assertion context
func (unit *TestUnit) Assertion() Assertion {
	return unit.assertion
}

// Assert tests the result is successful (true)
func (unit *TestUnit) Assert(args ...interface{}) {
	assertOk := false
	msg := "Assertion failed"
	skip := 1

	switch len(args) {
	case 1:
		assertOk = args[0].(bool)
	case 2:
		assertOk = args[0].(bool)
		msg = args[1].(string)
	case 3:
		assertOk = args[0].(bool)
		msg = args[1].(string)
		skip = args[2].(int)
	default:
		panic(&ArgumentError{"args", "invalid amount"})
	}

	// Output Assert
	unit.reporter.Asserted()

	if !assertOk {
		if _, file, line, ok := runtime.Caller(skip); ok {
			panic(&AssertionError{Msg: msg, File: file, Line: line})
		}

		panic(&AssertionError{Msg: msg})
	}
}
