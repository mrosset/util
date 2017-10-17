package main

import (
	"github.com/mrosset/util/json"
	"io"
	"log"
	"os"
)

func main() {
	err := format(os.Stdin, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}

func format(in io.Reader, out io.Writer) error {
	return json.Format(in, out)
}
