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

func (pbw *ProgressBarWriter) Write(b []byte) (n int, err error) {
	if pbw.done == 0 {
		pbw.start = time.Now()
	}
	switch {
	case pbw.total > 0:
		pbw.done += int64(len(b))
		percent := int((pbw.done * 100) / pbw.total)
		width := (80 - 9) - 40
		progress := (width * percent) / 100
		bar := strings.Repeat("#", int(progress))
		bps := float64(pbw.done) / time.Now().Sub(pbw.start).Seconds()
		fmt.Printf("\r[%-*s] %s/s %3.3s%% %s", width, bar, human.ByteSize(bps), strconv.Itoa(percent), pbw.prefix)
	default:
		fmt.Printf("\r%-20.20s", pbw.prefix)
	}
	return pbw.w.Write(b)
}

func (pbw *ProgressBarWriter) Close() error {
	fmt.Println()
	return nil
}
func NewProgressBarWriter(p string, t int64, w io.Writer) *ProgressBarWriter {
	return &ProgressBarWriter{prefix: p, total: t, w: w}
}
