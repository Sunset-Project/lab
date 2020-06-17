package lab

// TestController provides an API to halt a test in case of failure
type TestController interface {
	FailNow()
}

// DiagnosticTestController provides a TestController to track the amount of failures
type DiagnosticTestController struct {
	ImmediateFailures uint
}

// FailNow increases the amount of immediate failures by 1
func (controller *DiagnosticTestController) FailNow() {
	controller.ImmediateFailures = controller.ImmediateFailures + 1
}
