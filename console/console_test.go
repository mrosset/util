package console

import (
	"io"
	"os"
	"testing"
)

//var testfile = "Mac.OSX.Lion.10.7.2.dmg"
var testfile = "console.go"

func TestProgress(t *testing.T) {
	fi, err := os.Stat(testfile)
	if err != nil {
		t.Fatal(err)
	}
	fd, err := os.Open(testfile)
	if err != nil {
		t.Fatal(err)
	}
	defer fd.Close()
	nw, err := os.Create("new.file")
	if err != nil {
		t.Fatal(err)
	}
	defer nw.Close()
	pw := NewProgressBarWriter(testfile, fi.Size(), nw)
	_, err = io.Copy(pw, fd)
	if err != nil {
		t.Fatal(err)
	}
}
