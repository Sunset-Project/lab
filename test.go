package lab

import (
	"fmt"
	"testing"

	"github.com/Fire-Dragon-DoL/lab/reporting"
)

// Test represents a single test in `go test`
type Test struct {
	t         *testing.T
	assertion Assertion
	reporter  reporting.Reporter
}

// NewTest prepares a new test unit to test code
func NewTest(t *testing.T, reporter reporting.Reporter) *Test {
	test := &Test{}
	test.t = t
	test.assertion = Assertion(test.Assert)
	test.reporter = reporter
	return test
}

// Context opens a new context for this test unit
func (test *Test) Context(args ...interface{}) {
	var prose string
	var do func()

	switch len(args) {
	case 1:
		prose = args[0].(string)
	case 2:
		prose = args[0].(string)
		do = args[1].(func())
	default:
		panic(ArgumentError{"args", "invalid amount"})
	}

	// EnterContext
	test.reporter.ContextEntered(prose)
	defer func() {
		// LeaveContext
		test.reporter.ContextExited(prose)
	}()

	if do == nil {
		// SkipContext
		test.reporter.ContextSkipped(prose)
	} else {
		defer func() {
			err := recover()

			if err == nil {
				// SuccessContext
				test.reporter.ContextSucceeded(prose)
			} else {
				// Output error
				test.reporter.PanicInvoked(err)
				// FailContext
				test.reporter.ContextFailed(prose)
				test.t.FailNow()
			}
		}()

		do()
	}
}

// Test opens a new test section for this test unit
func (test *Test) Test(args ...interface{}) {
	var prose string
	var do func()

	switch len(args) {
	case 1:
		prose = args[0].(string)
	case 2:
		prose = args[0].(string)
		do = args[1].(func())
	default:
		panic(ArgumentError{"args", "invalid amount"})
	}

	// EnterTest
	test.reporter.TestStarted(prose)
	defer func() {
		// LeaveTest
		test.reporter.TestFinished(prose)
	}()

	if do == nil {
		// SkipTest
		test.reporter.TestSkipped(prose)
	} else {
		defer func() {
			err := recover()

			if err == nil {
				// SuccessTest
				test.reporter.TestPassed(prose)
			} else {
				// Output error
				test.reporter.PanicInvoked(err)
				// FailTest
				test.reporter.TestFailed(prose)
				test.t.FailNow()
			}
		}()

		do()
	}
}

// Assertion provides a new assertion context
func (test *Test) Assertion() Assertion {
	return test.assertion
}

// Assert tests the result is successful (true)
func (test *Test) Assert(args ...interface{}) {
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
		panic(ArgumentError{"args", "invalid amount"})
	}

	// Output Assert
	test.reporter.Asserted()

	if !assertOk {
		// if _, file, line, ok := runtime.Caller(skip); ok {
		// 	panic(&AssertionError{Msg: msg, File: file, Line: line})
		// }
		fmt.Printf("%+v", skip)

		panic(AssertionError{Msg: msg})
	}
}
