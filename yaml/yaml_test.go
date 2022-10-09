package yaml

import (
	"os"
	"testing"
)

type TestJson struct {
	Name string
	Id   int
}

func TestWritePretty(t *testing.T) {
	var (
		tj = &TestJson{"Test", 0}
	)
	err := WritePretty(tj, os.Stdout)
	if err != nil {
		t.Fatal(err)
	}
}
