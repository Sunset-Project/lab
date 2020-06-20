package reporting

import "github.com/sunset-project/lab/trace"

// BlockResult represents the success or failure of a Context or Test block
type BlockResult uint

// Blocks are values identifying if Context or Test functions failed
const (
	BlockSucceeded BlockResult = iota
	BlockSkipped
	BlockFailed
)

// Reporter can be used to customize the output of `lab`
type Reporter interface {
	Asserted()
	ContextEntered(prose string)
	ContextExited(prose string, result BlockResult)
	ContextSkipped(prose string)
	ContextSucceeded(prose string)
	ContextFailed(prose string)
	PanicInvoked(msg trace.Message)
	TestFailed(prose string)
	TestFinished(prose string, result BlockResult)
	TestPassed(prose string)
	TestSkipped(prose string)
	TestStarted(prose string)
}
