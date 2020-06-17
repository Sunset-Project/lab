package lab

// TestController provides an API to halt a test in case of failure
type TestController interface {
	FailNow()
}

// DiagnosticTestController provides a TestController to track the amount of failures
type DiagnosticTestController struct {
	immediateFailures uint
}

// FailNow increases the amount of immediate failures by 1
func (controller *DiagnosticTestController) FailNow() {
	controller.immediateFailures = controller.immediateFailures + 1
}

// HasFailed is true when the amount of immediate failures is greater than 0
func (controller *DiagnosticTestController) HasFailed() bool {
	return controller.immediateFailures > 0
}

// ImmediateFailures returns the amount of times `FailNow` has been called
func (controller *DiagnosticTestController) ImmediateFailures() uint {
	return controller.immediateFailures
}
