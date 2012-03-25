package console

import (
	"path/filepath"
	"testing"
)

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

func TestFileList(t *testing.T) {
	files, err := filepath.Glob("/home/strings/*")
	if err != nil {
		t.Error(err)
	}
	for _, _ = range files {
		//Println(f)
	}
	Flush()
}
