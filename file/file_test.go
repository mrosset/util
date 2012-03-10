package file

import (
	"fmt"
	"path/filepath"
	"testing"
)

var (
	existsFiles   = []string{"../util.go", "file.go"}
	notExistFiles = []string{"aaaaaaaa", "bbbbbbbbb"}
)

func TestExists(t *testing.T) {
	for _, f := range existsFiles {
		exists := Exists(f)
		if !exists {
			t.Errorf("expect to find %s got %v", f, exists)
		}
		t.Logf("%s -> %v", f, exists)
	}

	for _, f := range notExistFiles {
		exists := Exists(f)
		if exists {
			t.Errorf("expect not to find %s got %v", f, exists)
		}
		t.Logf("%s -> %v", f, exists)
	}
}

func TestMagic(t *testing.T) {
	files, _ := filepath.Glob("/home/strings//via/cache/sources/*")
	for _, file := range files {
		m, err := GetFileMagic(file)
		if err != nil {
			t.Error(err)
		}
		fmt.Printf("%-40.40s %v\n", filepath.Base(file), m)
	}
}
