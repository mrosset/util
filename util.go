package unix

import (
	"log"
	"os"
)

func FileExists(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	if !fi.IsDir() || fi.IsDir() || fi.Mode() == os.ModeSymlink {
		return true
	}
	return false
}

func CheckFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
