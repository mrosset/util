package json

import (
	"os"
	"testing"
)

type Test struct {
	Loooooooong string
	Short       string
}

func TestWritePretty(t *testing.T) {
	test := &Test{"Long", "Short"}
	err := WritePretty(&test, os.Stdout)
	if err != nil {
		t.Error(err)
	}
}

func TestGetFail(t *testing.T) {
	var (
		got = Get(nil, "http://localhost:10000")
	)
	if got == nil {
		t.Logf("expect %v -> got %s", nil, got)
		t.Error(got)
	}
}
