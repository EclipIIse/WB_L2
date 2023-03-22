package main

import "fmt"

type Command interface {
	Execute() error
}

type GreetCommand struct {
	Name string
}

func (c *GreetCommand) Execute() error {
	fmt.Printf("Hello, %s!\n", c.Name)
	return nil
}

type CommandInvoker struct {
	commands []Command
}

func (ci *CommandInvoker) AddCommand(cmd Command) {
	ci.commands = append(ci.commands, cmd)
}

func (ci *CommandInvoker) ExecuteCommands() error {
	for _, cmd := range ci.commands {
		err := cmd.Execute()
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	greetCmd := &GreetCommand{Name: "Koryun"}
	cmdInvoker := &CommandInvoker{}

	cmdInvoker.AddCommand(greetCmd)

	err := cmdInvoker.ExecuteCommands()
	if err != nil {
		fmt.Printf("Error executing commands: %v\n", err)
	}
}
