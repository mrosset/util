package command

import (
	"fmt"
	"testing"
)

func TestCommands(t *testing.T) {
	Add("test", test, "test function")
	Add("test1", test1, "test1 function")
	for _, c := range Commands {
		c.Run()
	}
}

func test() error {
	fmt.Println("test run here")
	return nil
}

func test1() error {
	fmt.Println("test1 run here")
	return nil
}
