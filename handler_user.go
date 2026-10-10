package main

import (
	"errors"
	"fmt"
)

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
