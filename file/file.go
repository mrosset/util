package file

import (
	"io"
	"os"
)

func Exists(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	if !fi.IsDir() || fi.IsDir() || fi.Mode() == os.ModeSymlink {
		return true
	}
	return false
}

func Copy(path string, w io.Writer) error {
	fd, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fd.Close()
	_, err = io.Copy(w, fd)
	return err
}
