package json

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"util/file"
)

// WriteGzJson marshal's a interface, and writes it to a gzipped file.
func WriteGzJson(v interface{}, file string) (err error) {
	fd, err := os.Create(file)
	if err != nil {
		return err
	}
	defer fd.Close()
	return WriteGzIo(v, fd)
}

// ReadGzJson read a gzipped json file and decodes it into an interface.
func ReadGzJson(v interface{}, file string) (err error) {
	fd, err := os.Open(file)
	if err != nil {
		return err
	}
	defer fd.Close()
	return ReadGzIo(v, fd)
}

func ReadGzIo(v interface{}, r io.Reader) (err error) {
	gz, err := gzip.NewReader(r)
	if err != nil {
		return err
	}
	defer gz.Close()
	return json.NewDecoder(gz).Decode(v)
}

func WriteGzIo(v interface{}, w io.Writer) (err error) {
	gz := gzip.NewWriter(w)
	defer gz.Close()
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	err = json.Indent(buf, b, "", "\t")
	if err != nil {
		return err
	}
	_, err = io.Copy(gz, buf)
	return err
}

// Write marshals a interface and writes it to a file
func Write(v interface{}, path string) (err error) {
	fd, err := os.Create(path)
	if err != nil {
		return err
	}
	defer fd.Close()
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	err = json.Indent(buf, b, "", "\t")
	if err != nil {
		return err
	}
	_, err = io.Copy(fd, buf)
	return err
}

// Read opens a json file and decodes it into interface
func Read(v interface{}, path string) (err error) {
	if !file.Exists(path) {
		return fmt.Errorf("%s does not exits.", path)
	}
	fd, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fd.Close()
	err = json.NewDecoder(fd).Decode(v)
	if err != nil {
		return err
	}
	return err
}

// PrintPretty marshal's an interface and ouputs formatted json.
func PrintPretty(v interface{}, w io.Writer) (err error) {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	err = json.Indent(buf, b, "", "\t")
	if err != nil {
		return err
	}
	_, err = io.Copy(w, buf)
	return err
}
