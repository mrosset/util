package console

import (
	"fmt"
	"github.com/str1ngs/util/human"
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
	switch {
	case self.total > 0:
		self.done = self.done + int64(len(b))
		percent := int((self.done * 100) / self.total)
		width := (80 - 9) - 32
		progress := (width * percent) / 100
		bar := strings.Repeat("#", int(progress))
		bps := float64(self.done) / time.Now().Sub(self.start).Seconds()
		fmt.Printf("\r%-20.20s [%-*s] %s/s %3.3s%%", self.prefix, width, bar, human.ByteSize(bps), strconv.Itoa(percent))
	default:
		fmt.Printf("\r%-20.20s", self.prefix)
	}
	return self.w.Write(b)
}

func (self *ProgressBarWriter) Close() error {
	fmt.Println()
	return nil
}
func NewProgressBarWriter(p string, t int64, w io.Writer) *ProgressBarWriter {
	return &ProgressBarWriter{prefix: p, total: t, w: w}
}
