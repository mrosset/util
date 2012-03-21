package console

import (
	"fmt"
	"os"
	"text/tabwriter"
)

var (
	tw = tabwriter.NewWriter(os.Stderr, 8, 0, 1, ' ', 0)
)

func Println(a ...interface{}) {
	var na []interface{}
	for _, i := range a {
		na = append(na, i, "\t")
	}
	fmt.Fprintln(tw, na...)
}

func Flush() {
	tw.Flush()
}
