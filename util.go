package util

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
)

func init() {
	log.SetFlags(0)
}

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
		_, file, line, _ := runtime.Caller(1)
		errFmt := fmt.Sprintf("%s:%v %s", path.Base(file), line, err)
		log.Fatal(errors.New(errFmt))
	}
}
