package main

import (
	"testing"
)

func TestConcreteCommandA_Execute(t *testing.T) {
	receiver := &Receiver{}
	cmd := &ConcreteCommandA{Receiver: receiver}

	if err := cmd.Execute(); err != nil {
		t.Errorf("Unexpected error executing command A: %v", err)
	}

	// Additional assertions if necessary
}

func TestConcreteCommandB_Execute(t *testing.T) {
	receiver := &Receiver{}
	cmd := &ConcreteCommandB{Receiver: receiver}

	if err := cmd.Execute(); err != nil {
		t.Errorf("Unexpected error executing command B: %v", err)
	}

	// Additional assertions if necessary
}

func TestInvoker_ExecuteCommand(t *testing.T) {
	receiver := &Receiver{}

	commandA := &ConcreteCommandA{Receiver: receiver}
	commandB := &ConcreteCommandB{Receiver: receiver}

	invoker := &Invoker{}

	invoker.Command = commandA
	if err := invoker.ExecuteCommand(); err != nil {
		t.Errorf("Unexpected error executing command A: %v", err)
	}

	invoker.Command = commandB
	if err := invoker.ExecuteCommand(); err != nil {
		t.Errorf("Unexpected error executing command B: %v", err)
	}

	// Additional assertions if necessary
}
