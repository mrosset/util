package magic

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestMagic(t *testing.T) {
	files, err := filepath.Glob("../testdata/magic/*")
	if err != nil {
		t.Fatal(err)
	}
	if len(files) == 0 {
		t.Fatal("expected files list greater the 0 to test")
	}
	for _, file := range files {
		m, err := GetFileMagic(file)
		if err != nil {
			t.Error(err)
		}
		fmt.Printf("%-40.40s %v\n", filepath.Base(file), m)
	}
}
