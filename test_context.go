package lab

import "testing"

// import "github.com/google/uuid"

type assertion func(bool)

// TestContext organizes test execution in contexts
type TestContext struct {
	t      *testing.T
	Assert assertion
}

// NewTestContext initializes a TestContext with existing Go test
func NewTestContext(t *testing.T) *TestContext {
	context := &TestContext{t, nil}
	context.Assert = context.AssertTrue
	return context
}

// Context opens a new test context
func (session *TestContext) Context(args ...interface{}) {

}

// Test opens a new test
func (session *TestContext) Test(args ...interface{}) {

}

// AssertTrue opens a new test
func (session *TestContext) AssertTrue(result bool) {

}

func (a assertion) Panic(do func()) {

}

func (a assertion) Ok(_ interface{}, ok bool) {

}
