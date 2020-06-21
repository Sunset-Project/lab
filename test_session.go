package lab

import (
	"github.com/sunset-project/lab/asserting"
	"github.com/sunset-project/lab/reporting"
	"github.com/sunset-project/lab/trace"
)

// TestSession represent the execution of a set of tests within one `go test` function. This type should be instantiated only through `NewTestSession`
type TestSession struct {
	controller TestController
	reporter   reporting.Reporter
}

// NewTestSession prepares a new session to test code
func NewTestSession(controller TestController, reporter reporting.Reporter) *TestSession {
	if controller == nil {
		panic(ArgumentError{"controller", "is nil"})
	}
	if reporter == nil {
		panic(ArgumentError{"reporter", "is nil"})
	}

	test := &TestSession{}
	test.controller = controller
	test.reporter = reporter

	return test
}

// Context opens a new context for this test unit
func (test *TestSession) Context(args ...interface{}) {
	var prose string
	var do func()
	panicked := true

	switch len(args) {
	case 1:
		prose = args[0].(string)
	case 2:
		prose = args[0].(string)
		do = args[1].(func())
	default:
		panic(ArgumentError{"args", "invalid amount"})
	}

	if do == nil {
		// SkipContext
		test.reporter.ContextSkipped(prose)
	} else {
		// EnterContext
		test.reporter.ContextEntered(prose)
		blockSucceeded := true
		defer func() {
			// LeaveContext
			test.reporter.ContextExited(prose, blockSucceeded)
		}()

		defer func() {
			err := recover()

			if panicked {
				// Output error
				if !test.controller.Failed() {
					panicMsg := trace.NewMessage(err)
					test.reporter.PanicInvoked(panicMsg)
				}
				// FailContext
				test.reporter.ContextFailed(prose)
				blockSucceeded = false
				test.controller.FailNow()
			} else {
				// SuccessContext
				test.reporter.ContextSucceeded(prose)
			}
		}()

		do()
	}
	panicked = false
}

// Test opens a new test section for this test unit
func (test *TestSession) Test(args ...interface{}) {
	var prose string
	var do func()
	panicked := true

	switch len(args) {
	case 1:
		prose = args[0].(string)
	case 2:
		prose = args[0].(string)
		do = args[1].(func())
	default:
		panic(ArgumentError{"args", "invalid amount"})
	}

	if do == nil {
		// SkipTest
		test.reporter.TestSkipped(prose)
	} else {
		// EnterTest
		test.reporter.TestStarted(prose)
		blockSucceeded := true
		defer func() {
			// LeaveTest
			test.reporter.TestFinished(prose, blockSucceeded)
		}()
		defer func() {
			if panicked {
				err := recover()
				// Output error
				if !test.controller.Failed() {
					panicMsg := trace.NewMessage(err)
					test.reporter.PanicInvoked(panicMsg)
				}
				// FailTest
				test.reporter.TestFailed(prose)
				blockSucceeded = false
				test.controller.FailNow()
			} else {
				// SuccessTest
				test.reporter.TestPassed(prose)
			}
		}()

		do()
	}
	panicked = false
}

// Assertion provides a new assertion context
func (test *TestSession) Assertion() asserting.Assertion { return asserting.Assertion(test.Assert) }

// Assert tests the result is successful (true)
func (test *TestSession) Assert(args ...interface{}) {
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
		panic(asserting.AssertionError{Msg: msg})
	}
}
