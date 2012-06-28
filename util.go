package util

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"runtime"
)

var (
	errbuf = new(bytes.Buffer)
)

func init() {
	log.SetPrefix("util: ")
	log.SetFlags(log.Lshortfile)
}

func CheckFatal(err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		errFmt := fmt.Sprintf("%s:%v %s", path.Base(file), line, err)
		log.Fatal(errors.New(errFmt))
	}
}

func Run(args ...string) (err error) {
	return RunIn(".", args...)
}

func RunIn(dir string, args ...string) (err error) {
	log.Printf("Running %s %s", args[0], args[1:])
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
