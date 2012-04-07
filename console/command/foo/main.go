package main

import (
	"fmt"
	"github.com/str1ngs/util/console/command"
)

func main() {
	command.Add("echo", echoArgs, "prints a line for each arg")       // add new command : foo echo
	command.Add("reverse", revArgs, "prints each arument in reverse") // add new command : foo reverse
	command.Run()                                                     // run commands
}

func revArgs() {
	command.ArgsDo(revArg) // loops through  all arguments after command
}

func revArg(a string) error {
	r := ""
	for i := len(a) - 1; i >= 0; i-- {
		r += string(a[i])
	}
	fmt.Println(r)
	return nil
}

func echoArgs() {
	command.ArgsDo(echoArg)
}

func echoArg(a string) error {
	fmt.Println(a)
	return nil
}
