package console

import (
	"fmt"
	"io"
	"time"

	"github.com/mrosset/util/human"
	"github.com/pterm/pterm"
)

type ProgressBarWriter struct {
	total       int64
	w           io.Writer
	done        int64
	start       time.Time
	prefix      string
	progressbar *pterm.ProgressbarPrinter
}

func (pw *ProgressBarWriter) Write(b []byte) (n int, err error) {
	var (
		percent = int((pw.done * 100) / pw.total)
		pb      = pw.progressbar
		bps     = float64(pw.done) / time.Now().Sub(pw.start).Seconds()
	)
	if pw.done == 0 {
		pw.start = time.Now()
	}
	pw.done += int64(len(b))
	switch {
	case pw.total > 0:
		title := fmt.Sprintf("%s %s/s", pw.prefix, human.ByteSize(bps))
		pb.UpdateTitle(title)
	default:
		pb.UpdateTitle(pw.prefix)
	}
	if percent > pw.progressbar.Current {

		pb.Increment()
	}
	return pw.w.Write(b)
}

func (pw *ProgressBarWriter) Close() error {
	pw.progressbar.Increment()
	return nil
}
func NewProgressBarWriter(p string, size int64, w io.Writer) *ProgressBarWriter {
	progress, err := pterm.DefaultProgressbar.WithTotal(100).Start()
	if err != nil {
		panic(err)
	}
	return &ProgressBarWriter{
		prefix:      p,
		total:       size,
		w:           w,
		progressbar: progress,
	}
}
