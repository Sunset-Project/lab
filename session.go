package lab

import (
	"github.com/sunset-project/lab/asserting"
	"github.com/sunset-project/lab/reporting"
)

// Session represents a test session, within a single test function
type Session interface {
	Context(...interface{})
	Test(...interface{})
	Assertion() asserting.Assertion
}

// NewSession creates a new test session
func NewSession(controller TestController, reporter reporting.Reporter) Session {
	return NewTestSession(controller, reporter)
}

// StartSession is a utility function to interact with a test Session without holding its reference. The returned tuple is the Context function, the Test function and the Assert function that work on the `Session`
func StartSession(controller TestController) (func(...interface{}), func(...interface{}), asserting.Assertion) {
	reporter := DefaultSessionReporter()
	session := NewSession(controller, reporter)
	return UseSession(session)
}

// UseSession is a utility function to prepare the test environment to use `lab` DSL. The returned tuple is the Context function, the Test function and the Assert function that work on the `Session`
func UseSession(session Session) (func(...interface{}), func(...interface{}), asserting.Assertion) {
	return session.Context, session.Test, session.Assertion()
}

// DefaultSessionReporter returns the default configured Reporter for `lab`
func DefaultSessionReporter() reporting.Reporter {
	return reporting.ProxyReporter{reporting.StdoutReporter()}
}
