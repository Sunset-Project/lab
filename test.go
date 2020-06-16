package lab

import (
	"github.com/sunset-project/lab/reporting"
	"github.com/sunset-project/lab/trace"
)

// Test represents a single test in `go test`
type Test struct {
	controller TestController
	assertion  Assertion
	reporter   reporting.Reporter
}

// NewTest prepares a new test unit to test code
func NewTest(controller TestController, reporter reporting.Reporter) *Test {
	if controller == nil {
		panic(ArgumentError{"controller", "is nil"})
	}
	if reporter == nil {
		panic(ArgumentError{"reporter", "is nil"})
	}

	test := &Test{}
	test.controller = controller
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
				panicMsg := trace.NewMessage(err)
				test.reporter.PanicInvoked(panicMsg)
				// FailContext
				test.reporter.ContextFailed(prose)
				test.controller.FailNow()
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
				panicMsg := trace.NewMessage(err)
				test.reporter.PanicInvoked(panicMsg)
				// FailTest
				test.reporter.TestFailed(prose)
				test.controller.FailNow()
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

	switch len(args) {
	case 1:
		assertOk = args[0].(bool)
	case 2:
		assertOk = args[0].(bool)
		msg = args[1].(string)
	case 3:
		assertOk = args[0].(bool)
		msg = args[1].(string)
	default:
		panic(ArgumentError{"args", "invalid amount"})
	}

	// Output Assert
	test.reporter.Asserted()

	if !assertOk {
		panic(AssertionError{Msg: msg})
	}
}
