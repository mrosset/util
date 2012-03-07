package console

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

type ProgressBarWriter struct {
	total  int64
	w      io.Writer
	done   int64
	start  time.Time
	prefix string
}

func (self *ProgressBarWriter) Write(b []byte) (n int, err error) {
	if self.done == 0 {
		self.start = time.Now()
	}
	self.done = self.done + int64(len(b))
	percent := int((self.done * 100) / self.total)
	width := (80 - 9) - 20
	progress := (width * percent) / 100
	bar := strings.Repeat("#", int(progress))
	fmt.Printf("\r%-20.20s [%-*s] %3.3s%%", self.prefix, width, bar, strconv.Itoa(percent))
	if self.done == self.total {
		fmt.Println()
	}
	return self.w.Write(b)
}

func NewProgressBarWriter(p string, t int64, w io.Writer) *ProgressBarWriter {
	return &ProgressBarWriter{prefix: p, total: t, w: w}
}
