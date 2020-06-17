package lab

import (
	"github.com/sunset-project/lab/reporting"
	"github.com/sunset-project/lab/trace"
)

// Test represents a single test in `go test`
type Test struct {
	Controller TestController
	Reporter   reporting.Reporter
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
	test.Controller = controller
	test.Reporter = reporter

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
	test.Reporter.ContextEntered(prose)
	defer func() {
		// LeaveContext
		test.Reporter.ContextExited(prose)
	}()

	if do == nil {
		// SkipContext
		test.Reporter.ContextSkipped(prose)
	} else {
		defer func() {
			err := recover()

			if err == nil {
				// SuccessContext
				test.Reporter.ContextSucceeded(prose)
			} else {
				// Output error
				panicMsg := trace.NewMessage(err)
				test.Reporter.PanicInvoked(panicMsg)
				// FailContext
				test.Reporter.ContextFailed(prose)
				test.Controller.FailNow()
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
	test.Reporter.TestStarted(prose)
	defer func() {
		// LeaveTest
		test.Reporter.TestFinished(prose)
	}()

	if do == nil {
		// SkipTest
		test.Reporter.TestSkipped(prose)
	} else {
		defer func() {
			err := recover()

			if err == nil {
				// SuccessTest
				test.Reporter.TestPassed(prose)
			} else {
				// Output error
				panicMsg := trace.NewMessage(err)
				test.Reporter.PanicInvoked(panicMsg)
				// FailTest
				test.Reporter.TestFailed(prose)
				test.Controller.FailNow()
			}
		}()

		do()
	}
}

// Assertion provides a new assertion context
func (test *Test) Assertion() Assertion { return Assertion(test.Assert) }

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
	test.Reporter.Asserted()

	if !assertOk {
		panic(AssertionError{Msg: msg})
	}
}
