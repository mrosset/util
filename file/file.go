package file

import (
	"io"
	"os"
	"path"
)

type Path string

func (p Path) Expand() string {
	return os.ExpandEnv(string(p))
}

func (p Path) Add(s string) string {
	return path.Join(p.Expand(), string(s))
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
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

func Cat(path string) error {
	fd, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fd.Close()
	_, err = io.Copy(os.Stderr, fd)
	return err
}
