package lab_test

import (
	"testing"

	"github.com/sunset-project/lab"
)

func TestContext(t *testing.T) {
	Context, Test, Assert := lab.StartTest(t)
	Context("Something", func() {
		Test("Else", func() {
			Assert(true)
		})
	})
}
