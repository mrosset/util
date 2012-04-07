package command

import (
	"flag"
	"os"
)

type Command struct {
	Name  string
	Run   func()
	Usage string
}

var (
	Commands = []*Command{}
)

func Run() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
		os.Exit(1)
	}
	for _, cmd := range Commands {
		if cmd.Name == args[0] {
			cmd.Run()
			return
		}
	}
	flag.Usage()
}

func Args() []string {
	return flag.Args()[1:]
}

func Add(n string, r func(), u string) {
	Commands = append(Commands, &Command{n, r, u})
}

func ArgsDo(f func(string) error) error {
	for _, a := range Args() {
		err := f(a)
		if err != nil {
			return err
		}
	}
	return nil
}
