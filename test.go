package lab

import "testing"

// Test represents a test execution
type Test interface {
	Context(...interface{})
	Test(...interface{})
	Assertion() Assertion
}

// BuildTest creates a new test test
func BuildTest(t *testing.T) Test { return NewTestUnit(t) }

// StartTest is a utility function to interact with a test Test without holding its reference. The returned tuple is the Context function, the Test function and the Assert function that work on the Test object
func StartTest(t *testing.T) (func(...interface{}), func(...interface{}), Assertion) {
	test := BuildTest(t)
	return test.Context, test.Test, test.Assertion()
}
