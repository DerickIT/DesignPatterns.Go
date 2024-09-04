package main

import "fmt"

type Command interface {
	Execute() error
}

type ConcreteCommand struct {
	receiver *Receiver
}

type Receiver struct{}

func (r *Receiver) Action() error {
	println("Receiver Action")
	return nil
}

func (c *ConcreteCommand) Execute() error {
	return c.receiver.Action()
}

type Invoker struct {
	commands []Command
}

func (i *Invoker) AddCommand(command Command) {
	i.commands = append(i.commands, command)
}

func (i *Invoker) ExecuteCommands() {
	for _, command := range i.commands {
		err := command.Execute()
		if err != nil {
			fmt.Println("Error executing command: %v\n", err)
		}
	}
}

func main() {
	receiver := &Receiver{}
	command := &ConcreteCommand{receiver: receiver}

	invoker := &Invoker{}
	invoker.AddCommand(command)

	invoker.ExecuteCommands()
}
