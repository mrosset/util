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
