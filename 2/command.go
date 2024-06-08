package main

import "fmt"

// Command defines the interface for executing a command.
type Command interface {
	Execute() error
}

// ConcreteCommandA is a concrete command that performs some action.
type ConcreteCommandA struct {
	Receiver *Receiver
}

func (c *ConcreteCommandA) Execute() error {
	return c.Receiver.ActionA()
}

// ConcreteCommandB is another concrete command that performs a different action.
type ConcreteCommandB struct {
	Receiver *Receiver
}

func (c *ConcreteCommandB) Execute() error {
	return c.Receiver.ActionB()
}

// Receiver is the object that performs the actions requested by commands.
type Receiver struct{}

func (r *Receiver) ActionA() error {
	fmt.Println("Receiver is performing Action A")
	// Logic for Action A
	return nil
}

func (r *Receiver) ActionB() error {
	fmt.Println("Receiver is performing Action B")
	// Logic for Action B
	return nil
}
