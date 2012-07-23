package magic

import (
	"compress/bzip2"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
)

type MagicBytes struct {
	first  byte
	second byte
}

type FileMagic struct {
	Enum  Magic
	magic MagicBytes
	human string
}

func (fm FileMagic) String() string {
	return fm.human
}

type Magic int

const (
	MagicGzip Magic = iota
	MagicPosixTar
	MagicGnuTar
	MagicBzip
	MagicMP3
	MagicElf
)

var (
	magics = []FileMagic{
		FileMagic{MagicGzip, MagicBytes{31, 139}, "gzip compressed data"},
		FileMagic{MagicPosixTar, MagicBytes{99, 99}, "POSIX tar archive (GNU)"},
		FileMagic{MagicGnuTar, MagicBytes{109, 112}, "tar archive"},
		FileMagic{MagicBzip, MagicBytes{66, 90}, "bzip compressed data"},
		FileMagic{MagicMP3, MagicBytes{73, 68}, "MPEG Layer III audio"},
		FileMagic{MagicElf, MagicBytes{127, 69}, "Elf binary"},
	}
)

func GetBytes(path string) (b []byte, err error) {
	b = make([]byte, 2)
	fd, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	_, err = fd.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func GetFileMagic(path string) (FileMagic, error) {
	b, err := GetBytes(path)
	if err != nil {
		return FileMagic{}, err
	}
	m := MagicBytes{b[0], b[1]}
	for _, mn := range magics {
		if m == mn.magic {
			return mn, nil
		}
	}
	return FileMagic{}, fmt.Errorf("Unknown file magic: %s", path)
}

func GetReader(path string) (r io.Reader, err error) {
	fm, err := GetFileMagic(path)
	if err != nil {
		return nil, err
	}
	fd, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	switch fm.Enum {
	case MagicGzip:
		gz, err := gzip.NewReader(fd)
		if err != nil {
			return nil, err
		}
		return gz, nil
	case MagicGnuTar, MagicPosixTar:
		return fd, nil
	case MagicBzip:
		bz := bzip2.NewReader(fd)
		return bz, nil
	}
	return nil, fmt.Errorf("Unknown reader for: %s", path)
}

func ContentType(path string) (ct string, err error) {
	b := make([]byte, 512)
	fd, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer fd.Close()
	_, err = fd.Read(b)
	if err != nil {
		return "", err
	}
	return http.DetectContentType(b), nil
}
