package console

import (
	"fmt"
	"github.com/mrosset/util/human"
	"os"
	"testing"
)

const BUFSIZE = 4096

func TestProgress(t *testing.T) {
	size := int64(1024 * 1024)
	fd, err := os.Open(os.DevNull)
	if err != nil {
		t.Error(err)
	}
	pw := NewProgressBarWriter(os.DevNull, size, fd)
	buf := make([]byte, BUFSIZE)
	for i := int64(0); i < size; i += BUFSIZE {
		pw.Write(buf)
	}
	pw.Close()
	fmt.Println("wrote", human.ByteSize(size), ">", os.DevNull)
}
