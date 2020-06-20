package reporting

import "github.com/sunset-project/lab/trace"

// Reporter can be used to customize the output of `lab`
type Reporter interface {
	Asserted()
	ContextEntered(prose string)
	ContextExited(prose string, success bool)
	ContextSkipped(prose string)
	ContextSucceeded(prose string)
	ContextFailed(prose string)
	PanicInvoked(msg trace.Message)
	TestFailed(prose string)
	TestFinished(prose string, success bool)
	TestPassed(prose string)
	TestSkipped(prose string)
	TestStarted(prose string)
}
