package walk

import (
	"fmt"
	"testing"
	"time"
)

func TestWalkC(t *testing.T) {
	start := time.Now()
	files, err := WalkC("/home/strings")
	if err != nil {
		t.Error(err)
	}
	fmt.Println("WalkC found", len(files), "files", time.Since(start).Seconds())
}

func TestWalk(t *testing.T) {
	start := time.Now()
	files, err := Walk("/home/strings")
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Walk  found", len(files), "files", time.Since(start).Seconds())
}
