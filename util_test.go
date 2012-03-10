package util

import (
	"testing"
)

func TestRun(t *testing.T) {
	err := Run("ls", ".", "-l")
	if err != nil {
		t.Error(err)
	}
}

func TestRunFail(t *testing.T) {
	err := Run("git", ".", "asdfasdf")
	if err != nil {
		t.Error(err)
	}
}
