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
		width := (80 - 9) - 32
		progress := (width * percent) / 100
		bar := strings.Repeat("#", int(progress))
		bps := float64(self.done) / time.Now().Sub(self.start).Seconds()
		fmt.Printf("\r%-20.20s [%-*s] %s/s %3.3s%%", self.prefix, width, bar, ByteSize(bps), strconv.Itoa(percent))
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

type ByteSize float64

const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func (b ByteSize) String() string {
	switch {
	case b >= YB:
		return fmt.Sprintf("%5.1fYB", float64(b/YB))
	case b >= ZB:
		return fmt.Sprintf("%5.1fZB", float64(b/ZB))
	case b >= EB:
		return fmt.Sprintf("%5.1fEB", float64(b/EB))
	case b >= PB:
		return fmt.Sprintf("%5.1fPB", float64(b/PB))
	case b >= TB:
		return fmt.Sprintf("%5.1fTB", float64(b/TB))
	case b >= GB:
		return fmt.Sprintf("%5.1fGB", float64(b/GB))
	case b >= MB:
		return fmt.Sprintf("%5.1fMB", float64(b/MB))
	case b >= KB:
		return fmt.Sprintf("%5.1fKB", float64(b/KB))
	}
	return fmt.Sprintf("%.2fB", float64(b))
}
