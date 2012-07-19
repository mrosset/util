package json

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"github.com/str1ngs/util/file"
	"io"
	"os"
	"strings"
	"text/tabwriter"
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
	return WritePretty(v, fd)
}

// Read opens a json file and decodes it into interface
func Read(v interface{}, path string) (err error) {
	if !file.Exists(path) {
		return fmt.Errorf("%s does not exist.", path)
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

func Clean(v interface{}, w io.Writer) (err error) {
	buf := new(bytes.Buffer)
	err = WritePretty(v, buf)
	if err != nil {
		return err
	}
	for {
		b, err := buf.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		switch b {
		case '{', '}', '"', ',', '[', ']', '\t':
			continue
		default:
			w.Write([]byte{b})
		}
	}
	return nil
}

// PrintPretty marshal's an interface and ouputs formatted json to writer.
func WritePretty(v interface{}, w io.Writer) (err error) {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	err = json.Indent(buf, b, "", "\t")
	//err = json.Compact(buf, b)
	if err != nil {
		return err
	}
	br := bufio.NewReader(buf)
	tw := tabwriter.NewWriter(w, 4, 0, 1, ' ', 0)
	for {
		b, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		line := string(b) + "\n"
		line = strings.Replace(line, ":", "\t:", 1)
		_, err = tw.Write([]byte(line))
		if err != nil {
			return err
		}
	}
	return tw.Flush()
}
