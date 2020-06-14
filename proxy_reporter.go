package lab

// ProxyReporter broadcast reporter calls to all configured reporters
type ProxyReporter struct {
	reporters []Reporter
}

// BuildProxyReporter instantiate a ProxyReporter with the provided reporters
func BuildProxyReporter(reporters ...Reporter) *ProxyReporter {
	return &ProxyReporter{reporters}
}

// Reporters provides the list of reporters the ProxyReporter broadcasts to
func (reporter *ProxyReporter) Reporters() []Reporter {
	return reporter.reporters
}

// Asserted broadcasts Asserted
func (reporter *ProxyReporter) Asserted() {
	for _, r := range reporter.reporters {
		r.Asserted()
	}
}

// ContextEntered broadcasts ContextEntered
func (reporter *ProxyReporter) ContextEntered(prose string) {
	for _, r := range reporter.reporters {
		r.ContextEntered(prose)
	}
}

// ContextExited broadcasts ContextExited
func (reporter *ProxyReporter) ContextExited(prose string) {
	for _, r := range reporter.reporters {
		r.ContextExited(prose)
	}
}

// ContextSkipped broadcasts ContextSkipped
func (reporter *ProxyReporter) ContextSkipped(prose string) {
	for _, r := range reporter.reporters {
		r.ContextSkipped(prose)
	}
}

// ContextSucceeded broadcasts ContextSucceeded
func (reporter *ProxyReporter) ContextSucceeded(prose string) {
	for _, r := range reporter.reporters {
		r.ContextSucceeded(prose)
	}
}

// ContextFailed broadcasts ContextFailed
func (reporter *ProxyReporter) ContextFailed(prose string) {
	for _, r := range reporter.reporters {
		r.ContextFailed(prose)
	}
}

// PanicInvoked broadcasts PanicInvoked
func (reporter *ProxyReporter) PanicInvoked(msg interface{}) {
	for _, r := range reporter.reporters {
		r.PanicInvoked(msg)
	}
}

// TestFailed broadcasts TestFailed
func (reporter *ProxyReporter) TestFailed(prose string) {
	for _, r := range reporter.reporters {
		r.TestFailed(prose)
	}
}

// TestFinished broadcasts TestFinished
func (reporter *ProxyReporter) TestFinished(prose string) {
	for _, r := range reporter.reporters {
		r.TestFinished(prose)
	}
}

// TestPassed broadcasts TestPassed
func (reporter *ProxyReporter) TestPassed(prose string) {
	for _, r := range reporter.reporters {
		r.TestPassed(prose)
	}
}

// TestSkipped broadcasts TestSkipped
func (reporter *ProxyReporter) TestSkipped(prose string) {
	for _, r := range reporter.reporters {
		r.TestSkipped(prose)
	}
}

// TestStarted broadcasts TestStarted
func (reporter *ProxyReporter) TestStarted(prose string) {
	for _, r := range reporter.reporters {
		r.TestStarted(prose)
	}
}
