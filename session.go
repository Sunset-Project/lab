package lab

import "testing"

// Session represents a test session, within a single test function
type Session interface {
	Context(...interface{})
	Test(...interface{})
	Assertion() Assertion
}

// BuildSession creates a new test session
func BuildSession(t *testing.T) Session { return NewTest(t) }

// StartSession is a utility function to interact with a test Session without holding its reference. The returned tuple is the Context function, the Test function and the Assert function that work on the Session object
func StartSession(t *testing.T) (func(...interface{}), func(...interface{}), Assertion) {
	session := BuildSession(t)
	return session.Context, session.Test, session.Assertion()
}
