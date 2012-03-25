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
	for _, j := range a {
		fmt.Fprintf(tw, "%v\t", j)
	}
	fmt.Fprintf(tw, "\n")
}

func Flush() {
	tw.Flush()
}
