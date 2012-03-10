package magic

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestMagic(t *testing.T) {
	files, _ := filepath.Glob("/home/strings/via/cache/sources/*")
	for _, file := range files {
		m, err := GetFileMagic(file)
		if err != nil {
			t.Error(err)
		}
		fmt.Printf("%-40.40s %v\n", filepath.Base(file), m)
	}
}
