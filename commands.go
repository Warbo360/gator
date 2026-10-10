package main

import (
	"errors"
)

// 1.0 struct for commands, these will live in the commands registry

type command struct {
	name string
	args []string
}

// 2.0 struct acting as a commands registry

type commands struct {
	cmds map[string]func(*state, command) error
}

// 2.1 method for commands to run a registered command

func (c *commands) run(s *state, cmd command) error {

	func_val, exists := c.cmds[cmd.name]
	if !exists {
		return errors.New("command is not registered")
	}

	return func_val(s, cmd)
}

// 2.2 method for commands to run a registered command

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmds[name] = f
}
