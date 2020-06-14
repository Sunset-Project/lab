package reporting

// Reporter can be used to customize the output of `lab`
type Reporter interface {
	Asserted()
	ContextEntered(prose string)
	ContextExited(prose string)
	ContextSkipped(prose string)
	ContextSucceeded(prose string)
	ContextFailed(prose string)
	PanicInvoked(msg interface{})
	TestFailed(prose string)
	TestFinished(prose string)
	TestPassed(prose string)
	TestSkipped(prose string)
	TestStarted(prose string)
}
