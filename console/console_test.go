package console

import (
	"os"
	"path"
	"path/filepath"
	"testing"
	"time"
)

func TestFileList(t *testing.T) {
	home := os.ExpandEnv("$HOME/*")
	files, err := filepath.Glob(home)
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

func TestProgressBar(t *testing.T) {
	total := 100
	pb := NewProgressBar("test", 1, 100)
	for i := 0; i < total; i++ {
		time.Sleep(time.Second / 100)
		pb.Step()
	}
}
