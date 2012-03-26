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
	for i, j := range a {
		if i == len(a)-1 {
			fmt.Fprintf(tw, "%v", j)
			continue
		}
		fmt.Fprintf(tw, "%v\t", j)
	}
	fmt.Fprintf(tw, "\n")
}

func Flush() {
	tw.Flush()
}
