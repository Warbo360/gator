package main

import (
	"errors"
	"fmt"
)

// struct for commands, these will live in the commands registry

type command struct {
	name string
	args []string
}

// struct acting as a commands registry

type commands struct {
	cmds map[string]func(*state, command) error
}

// Login handler for setting the config username as the passed arg

func handlerLogin(s *state, cmd command) error {

	// Check passed command args slice if empty, if so return error

	if len(cmd.args) == 0 {
		return errors.New("username required")
	}

	// Set the state's config username to the passed 1st arg and error check

	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}

    fmt.Printf("%v is now logged in\n", s.cfg.CurrentUsername)

	return nil
}

func (c *commands) run(s *state, cmd command) error {

	func_val, exists := c.cmds[cmd.name]
	if !exists {
		return errors.New("command is not registered")
	}

	return func_val(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmds[name] = f
}
