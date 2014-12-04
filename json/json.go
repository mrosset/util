package json

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/str1ngs/util/file"
	"io"
	"net/http"
	"os"
	"strings"
	"text/tabwriter"
)

var (
	client = new(http.Client)
)

// Marshal's a interface, and writes it to a gzipped file.
func WriteGz(v interface{}, file string) (err error) {
	fd, err := os.Create(file)
	if err != nil {
		return err
	}
	defer fd.Close()
	return WriteGzIo(v, fd)
}

// Read a gzipped json file and decodes it into an interface.
func ReadGz(v interface{}, file string) (err error) {
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

// Reads a json file and decodes it into interface
func Read(v interface{}, path string) (err error) {
	if !file.Exists(path) {
		return fmt.Errorf("%s does not exist.", path)
	}
	fd, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fd.Close()
	return json.NewDecoder(fd).Decode(v)
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

func Format(in io.Reader, out io.Writer) error {
	var v interface{}
	err := json.NewDecoder(in).Decode(&v)
	if err != nil {
		return err
	}
	return WritePretty(v, out)
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
		line = strings.Replace(line, "\":", "\"\t:", 1)
		line = strings.Replace(line, "],", "],\t", 1)
		line = strings.Replace(line, "},", "},\t", 1)
		_, err = tw.Write([]byte(line))
		if err != nil {
			return err
		}
	}
	return tw.Flush()
}

// Decodes a URL to interface
func Get(v interface{}, url string) (err error) {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}
	return json.NewDecoder(resp.Body).Decode(v)
}
