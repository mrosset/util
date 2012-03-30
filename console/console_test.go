package console

import (
	"path"
	"path/filepath"
	"testing"
)

func TestFileList(t *testing.T) {
	files, err := filepath.Glob("/home/strings/*")
	if err != nil {
		t.Error(err)
	}
	for _, f := range files {
		Println(path.Base(f), f)
	}
	Flush()
}

func TestConsoleOutput(t *testing.T) {
	var (
		long  = "looooooooooooooooooooooooooooooooooooooooooooooooooooooooooog"
		short = "short"
	)
	for i := 0; i < 4; i++ {
		Println(long, i)
		Println(short, i)
	}
	Flush()
}
