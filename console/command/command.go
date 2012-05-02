package command

import (
	"flag"
	"fmt"
	"github.com/str1ngs/util/console"
	"os"
)

type Command struct {
	Name  string       // Name of the command line argument
	Run   func() error // Function associated with command
	Usage string       // Usage test
}

var (
	commands = []*Command{}
)

// Starts the command pass on command line
func Run() error {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		Usage()
	}
	for _, cmd := range commands {
		if cmd.Name == args[0] {
			return cmd.Run()
		}
	}
	Usage()
	return nil
}

// Returns the arguments after the command argument
func Args() []string {
	return flag.Args()[1:]
}

// Adds a command
func Add(n string, r func() error, u string) {
	commands = append(commands, &Command{n, r, u})
}

// Loops through each argument after the command argument
// and pass it to func f
func ArgsDo(f func(string) error) error {
	for _, a := range Args() {
		err := f(a)
		if err != nil {
			return err
		}
	}
	return nil
}

// Prints out flag usage and then prints out command
// and there usage
func Usage() {
	flag.Usage()
	fmt.Println("Commands: ")
	for _, c := range commands {
		console.Println("     ", c.Name, "      ", c.Usage)
	}
	console.Flush()
	os.Exit(1)
}
