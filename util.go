package util

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"runtime"
)

var Verbose = true

func init() {
	log.SetFlags(0)
}

func CheckFatal(err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		errFmt := fmt.Sprintf("%s:%v %s", path.Base(file), line, err)
		log.Fatal(errors.New(errFmt))
	}
}

func Run(bin, dir string, args ...string) (err error) {
	cmd := exec.Command(bin, args...)
	cmd.Dir = dir
	if Verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	return cmd.Run()
}
