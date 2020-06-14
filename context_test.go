package lab_test

import (
	"testing"

	"github.com/sunset-project/lab"
)

func TestContext(t *testing.T) {
	Context, Test, Assert := lab.StartSession(t)
	Context("Something", func() {
		Test("Else", func() {
			Assert(true)
		})
	})
}

func TestContext2(t *testing.T) {
	Context, Test, Assert := lab.StartSession(t)
	Context("Something2s", func() {
		Test("Else2", func() {
			Assert(false)
			Assert(true)
		})
	})
}
