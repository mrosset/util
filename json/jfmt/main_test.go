package main

import (
	"bytes"
	"os"
	"testing"
)

var data = []byte(`{"Looooooooong": "stuff"}`)

func TestFormat(t *testing.T) {
	buf := new(bytes.Buffer)
	buf.Write(data)
	err := format(buf, os.Stdout)
	if err != nil {
		t.Error(err)
	}
}
