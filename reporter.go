package lab

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

// DefaultReporter returns the default configured Reporter for `lab`
func DefaultReporter() Reporter {
	return BuildProxyReporter(OutputReporter)
}
