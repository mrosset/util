package file

import (
	"crypto/md5"
	"fmt"
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

func Copy(w io.Writer, src string) error {
	fd, err := os.Open(src)
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

func Md5(file string) (hash string, err error) {
	h := md5.New()
	fd, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer fd.Close()
	io.Copy(h, fd)
	return fmt.Sprintf("%X", h.Sum(nil)), err
}
