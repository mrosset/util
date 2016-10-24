package yaml

import (
	"gopkg.in/yaml.v2"
	"io"
	"os"
	"text/tabwriter"
)

func WriteFile(v interface{}, name string) error {
	fd, err := os.Create(name)
	if err != nil {
		return err
	}
	defer fd.Close()
	return WritePretty(v, fd)
}

func Write(v interface{}, w io.Writer) error {
	b, err := yaml.Marshal(v)
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	return err
}

func WritePretty(v interface{}, w io.Writer) error {
	b, err := yaml.Marshal(v)
	if err != nil {
		return err
	}
	tw := tabwriter.NewWriter(w, 4, 0, 1, ' ', 0)
	for _, n := range b {
		tw.Write([]byte{n})
		if string(n) == ":" {
			tw.Write([]byte{'\t'})
		}
	}
	tw.Flush()
	return nil
}
