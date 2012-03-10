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
	switch {
	case self.total > 0:
		self.done = self.done + int64(len(b))
		percent := int((self.done * 100) / self.total)
		width := (80 - 9) - 30
		progress := (width * percent) / 100
		bar := strings.Repeat("#", int(progress))
		bps := float64(self.done) / time.Now().Sub(self.start).Seconds()
		//fmt.Printf("\r%-20.20s [%-*s] %3.3s%%", self.prefix, width, bar, strconv.Itoa(percent))
		fmt.Printf("\r%-20.20s [%-*s] %s %3.3s%%", self.prefix, width, bar, speed(bps), strconv.Itoa(percent))
	default:
		fmt.Printf("\r%-20.20s", self.prefix)
	}
	return self.w.Write(b)
}

func NewProgressBarWriter(p string, t int64, w io.Writer) *ProgressBarWriter {
	return &ProgressBarWriter{prefix: p, total: t, w: w}
}

func speed(b float64) string {
	switch {
	case b < 1024:
		return fmt.Sprintf("%vB/s", b)
	case b < 1024*1000:
		return fmt.Sprintf("%5.1fKB/s", b/1024)
	case b < 1024*1024*1000:
		return fmt.Sprintf("%5.1fMB/s", b/1024/1024)
	default:
		return fmt.Sprintf("%5.1fGB/s", b/1024/1024/1024)
	}
	return fmt.Sprintf("%5.1fGB/s", b/1024/1024/1024)
}
