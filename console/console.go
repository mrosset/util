package console

import (
	"fmt"
	"os"
	"text/tabwriter"
)

var (
	tw = tabwriter.NewWriter(os.Stderr, 1, 0, 1, ' ', 0)
)

func Println(a ...interface{}) {
	for _, i := range a {
		fmt.Fprintf(tw, "%v\t", i)
	}
	fmt.Fprintf(tw, "\n")
}

func Flush() {
	tw.Flush()
}
