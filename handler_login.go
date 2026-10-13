package main

import (
	"errors"
	"fmt"
	"context"
)

func handlerLogin(s *state, cmd command) error {

	// Check passed command args slice if empty, if so return error

	if len(cmd.args) == 0 {
		return errors.New("username required")
	}

	// Check the passed state to access the database to check if the passed name is a user in the database

	user, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	// Set the state's config username to the passed 1st arg and error check

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Printf("[UPDATE]: %v is now logged in\n", s.cfg.CurrentUsername)

	return nil
}
