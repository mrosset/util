package console

import (
	"testing"
)

func TestConsoleOutput(t *testing.T) {
	var (
		long  = "looooooooooooooooooooooooooooooooooooooooooooooooooooooooooog"
		short = "short"
	)
	for i := 0; i < 4; i++ {
		Println(long, i)
		Println(short, i)
	}
	Flush()
}
