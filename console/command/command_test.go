package command

import (
	"fmt"
	"testing"
)

func TestCommands(t *testing.T) {
	Add("test", test, "test function")
	Add("test1", test1, "test1 function")
	for _, c := range commands {
		c.Run()
	}
}

func test() {
	fmt.Println("test run here")
}

func test1() {
	fmt.Println("test1 run here")
}
