// package command provides a simple way to create cli applications in go.
package command

import (
	"fmt"
	"os"
)

type Command interface {
	Synopsis() string
	Help() string
	Run() error
}

type Commander struct {
	name     string
	commands map[string]Command
}

func New(name string) *Commander {
	cmder := &Commander{
		name: name,
	}

	return cmder
}

func (c *Commander) Register(name string, cmd Command) {
	c.commands[name] = cmd
}

func (c *Commander) Run() error {
	if len(os.Args) > 1 {
		for k, cmd := range c.commands {
			if k == os.Args[1] {
				return cmd.Run()
			}
		}
	}
	c.Usage()
	return nil
}

func (c *Commander) Usage() {
	fmt.Printf("%s <command> [args]\n", c.name)
	for _, cmd := range c.commands {
		fmt.Println(cmd.Synopsis())
	}
}
