package lab

import (
	"github.com/sunset-project/lab/asserting"
	"github.com/sunset-project/lab/reporting"
	"github.com/sunset-project/lab/trace"
)

type testSession struct {
	controller TestController
	reporter   reporting.Reporter
}

// NewTestSession prepares a new test unit to test code
func NewTestSession(controller TestController, reporter reporting.Reporter) *testSession {
	if controller == nil {
		panic(ArgumentError{"controller", "is nil"})
	}
	if reporter == nil {
		panic(ArgumentError{"reporter", "is nil"})
	}

	test := &testSession{}
	test.controller = controller
	test.reporter = reporter

	return test
}

// Context opens a new context for this test unit
func (test *testSession) Context(args ...interface{}) {
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

	// EnterContext
	test.reporter.ContextEntered(prose)
	blockResult := reporting.BlockSucceeded
	defer func() {
		// LeaveContext
		test.reporter.ContextExited(prose, blockResult)
	}()

	if do == nil {
		// SkipContext
		test.reporter.ContextSkipped(prose)
		blockResult = reporting.BlockSkipped
	} else {
		defer func() {
			err := recover()

			if panicked {
				// Output error
				panicMsg := trace.NewMessage(err)
				test.reporter.PanicInvoked(panicMsg)
				// FailContext
				test.reporter.ContextFailed(prose)
				blockResult = reporting.BlockFailed
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
func (test *testSession) Test(args ...interface{}) {
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

	// EnterTest
	test.reporter.TestStarted(prose)
	blockResult := reporting.BlockSucceeded
	defer func() {
		// LeaveTest
		test.reporter.TestFinished(prose, blockResult)
	}()

	if do == nil {
		// SkipTest
		test.reporter.TestSkipped(prose)
		blockResult = reporting.BlockSkipped
	} else {
		defer func() {

			if panicked {
				err := recover()
				// Output error
				panicMsg := trace.NewMessage(err)
				test.reporter.PanicInvoked(panicMsg)
				// FailTest
				test.reporter.TestFailed(prose)
				blockResult = reporting.BlockFailed
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
func (test *testSession) Assertion() asserting.Assertion { return asserting.Assertion(test.Assert) }

// Assert tests the result is successful (true)
func (test *testSession) Assert(args ...interface{}) {
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
