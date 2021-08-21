// package command provides a simple way to create cli applications in go.
package command

import (
	"fmt"
	"os"
)

type Command interface {
	Name() string
	Usage() string
	Run() error
}

type Commander struct {
	name     string
	commands []Command
}

func New(name string) *Commander {
	cmder := &Commander{
		name: name,
	}

	return cmder
}

func (c *Commander) Register(cmd Command) {
	c.commands = append(c.commands, cmd)
}

func (c *Commander) Run() error {
	for _, cmd := range c.commands {
		if cmd.Name() == os.Args[1] {
			return cmd.Run()
		}
	}

	c.Usage()
	return nil
}

func (c *Commander) Usage() {
	fmt.Printf("%s <command> [args]", c.name)
	for _, cmd := range c.commands {
		fmt.Println(cmd.Usage())
	}
}
