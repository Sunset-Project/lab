package reporting

// ProxyReporter broadcast reporter calls to all configured reporters
type ProxyReporter struct {
	Reporters []Reporter
}

// NewProxyReporter instantiate a ProxyReporter with the provided reporters
func NewProxyReporter(reporters ...Reporter) ProxyReporter {
	return ProxyReporter{Reporters: reporters}
}

// Asserted broadcasts Asserted
func (reporter ProxyReporter) Asserted() {
	for _, r := range reporter.Reporters {
		r.Asserted()
	}
}

// ContextEntered broadcasts ContextEntered
func (reporter ProxyReporter) ContextEntered(prose string) {
	for _, r := range reporter.Reporters {
		r.ContextEntered(prose)
	}
}

// ContextExited broadcasts ContextExited
func (reporter ProxyReporter) ContextExited(prose string) {
	for _, r := range reporter.Reporters {
		r.ContextExited(prose)
	}
}

// ContextSkipped broadcasts ContextSkipped
func (reporter ProxyReporter) ContextSkipped(prose string) {
	for _, r := range reporter.Reporters {
		r.ContextSkipped(prose)
	}
}

// ContextSucceeded broadcasts ContextSucceeded
func (reporter ProxyReporter) ContextSucceeded(prose string) {
	for _, r := range reporter.Reporters {
		r.ContextSucceeded(prose)
	}
}

// ContextFailed broadcasts ContextFailed
func (reporter ProxyReporter) ContextFailed(prose string) {
	for _, r := range reporter.Reporters {
		r.ContextFailed(prose)
	}
}

// PanicInvoked broadcasts PanicInvoked
func (reporter ProxyReporter) PanicInvoked(msg interface{}) {
	for _, r := range reporter.Reporters {
		r.PanicInvoked(msg)
	}
}

// TestFailed broadcasts TestFailed
func (reporter ProxyReporter) TestFailed(prose string) {
	for _, r := range reporter.Reporters {
		r.TestFailed(prose)
	}
}

// TestFinished broadcasts TestFinished
func (reporter ProxyReporter) TestFinished(prose string) {
	for _, r := range reporter.Reporters {
		r.TestFinished(prose)
	}
}

// TestPassed broadcasts TestPassed
func (reporter ProxyReporter) TestPassed(prose string) {
	for _, r := range reporter.Reporters {
		r.TestPassed(prose)
	}
}

// TestSkipped broadcasts TestSkipped
func (reporter ProxyReporter) TestSkipped(prose string) {
	for _, r := range reporter.Reporters {
		r.TestSkipped(prose)
	}
}

// TestStarted broadcasts TestStarted
func (reporter ProxyReporter) TestStarted(prose string) {
	for _, r := range reporter.Reporters {
		r.TestStarted(prose)
	}
}
