package main
// Invoker holds a command and invokes its execution.
type Invoker struct {
	Command Command
}

func (i *Invoker) ExecuteCommand() error {
	return i.Command.Execute()
}

// Client code that uses the commands.
func main() {
	receiver := &Receiver{}

	// Create command objects with receivers.
	commandA := &ConcreteCommandA{Receiver: receiver}
	commandB := &ConcreteCommandB{Receiver: receiver}

	// Create an invoker and associate it with a command.
	invoker := &Invoker{}

	// Execute the command via the invoker.
	invoker.Command = commandA
	invoker.ExecuteCommand()

	invoker.Command = commandB
	invoker.ExecuteCommand()
}
