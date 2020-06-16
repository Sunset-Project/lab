package reporting

import "github.com/sunset-project/lab/trace"

// ProxyReporter broadcast reporter calls to all configured reporters
type ProxyReporter []Reporter

// Asserted broadcasts Asserted
func (reporters ProxyReporter) Asserted() {
	for _, reporter := range reporters {
		reporter.Asserted()
	}
}

// ContextEntered broadcasts ContextEntered
func (reporters ProxyReporter) ContextEntered(prose string) {
	for _, reporter := range reporters {
		reporter.ContextEntered(prose)
	}
}

// ContextExited broadcasts ContextExited
func (reporters ProxyReporter) ContextExited(prose string) {
	for _, reporter := range reporters {
		reporter.ContextExited(prose)
	}
}

// ContextSkipped broadcasts ContextSkipped
func (reporters ProxyReporter) ContextSkipped(prose string) {
	for _, reporter := range reporters {
		reporter.ContextSkipped(prose)
	}
}

// ContextSucceeded broadcasts ContextSucceeded
func (reporters ProxyReporter) ContextSucceeded(prose string) {
	for _, reporter := range reporters {
		reporter.ContextSucceeded(prose)
	}
}

// ContextFailed broadcasts ContextFailed
func (reporters ProxyReporter) ContextFailed(prose string) {
	for _, reporter := range reporters {
		reporter.ContextFailed(prose)
	}
}

// PanicInvoked broadcasts PanicInvoked
func (reporters ProxyReporter) PanicInvoked(msg trace.Message) {
	for _, reporter := range reporters {
		reporter.PanicInvoked(msg)
	}
}

// TestFailed broadcasts TestFailed
func (reporters ProxyReporter) TestFailed(prose string) {
	for _, reporter := range reporters {
		reporter.TestFailed(prose)
	}
}

// TestFinished broadcasts TestFinished
func (reporters ProxyReporter) TestFinished(prose string) {
	for _, reporter := range reporters {
		reporter.TestFinished(prose)
	}
}

// TestPassed broadcasts TestPassed
func (reporters ProxyReporter) TestPassed(prose string) {
	for _, reporter := range reporters {
		reporter.TestPassed(prose)
	}
}

// TestSkipped broadcasts TestSkipped
func (reporters ProxyReporter) TestSkipped(prose string) {
	for _, reporter := range reporters {
		reporter.TestSkipped(prose)
	}
}

// TestStarted broadcasts TestStarted
func (reporters ProxyReporter) TestStarted(prose string) {
	for _, reporter := range reporters {
		reporter.TestStarted(prose)
	}
}
