package magic

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestMagic(t *testing.T) {
	files, _ := filepath.Glob("/home/strings/via/cache/src/*")
	if len(files) == 0 {
		t.Errorf("expected files list greater the 0 to test")
		t.FailNow()
	}
	for _, file := range files {
		m, err := GetFileMagic(file)
		if err != nil {
			t.Error(err)
		}
		fmt.Printf("%-40.40s %v\n", filepath.Base(file), m)
	}
}

func TestContentType(t *testing.T) {
	files, _ := filepath.Glob("/home/strings/Music/Johnny Cash/The Very Best Of/*")
	if len(files) == 0 {
		t.Errorf("expected files list greater the 0 to test")
		t.FailNow()
	}
	for _, file := range files {
		b, err := GetFileMagic(file)
		if err != nil {
			t.Error(err)
		}
		fmt.Printf("%-40.40s %v\n", filepath.Base(file), b)
	}
}
