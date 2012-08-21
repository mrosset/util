package console

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

var (
	tw = tabwriter.NewWriter(os.Stdout, 1, 0, 1, ' ', 0)
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

type ProgressBar struct {
	prefix   string
	step     int
	total    int
	progress int
}

func (pb *ProgressBar) Step() {
	pb.progress += pb.step
	width := (80 - 4) - (len(pb.prefix) + 1)
	percent := int((pb.progress * 100) / pb.total)
	progress := (width * percent) / 100
	bar := strings.Repeat("#", int(progress))
	fmt.Printf("\r%s %-*s %3.3s%%", pb.prefix, width, bar, strconv.Itoa(percent))
	if pb.progress == pb.total {
		fmt.Println()
	}
}

func NewProgressBar(prefix string, step, total int) ProgressBar {
	return ProgressBar{prefix, step, total, 0}
}
