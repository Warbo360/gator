package main

import (
	"context"
	"github.com/Warbo360/gator/internal/database"
)

// Helper function for "logged in" commands like feeds, follow, etc which abstract logic of getting current logged in
// user away from those various handler functions, handler function is passed to this function and returns a new
// function with the resulting user state automatically

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {

        // Uses current state.config to get current username

		currentUsername := s.cfg.CurrentUsername

		// Uses username for SQL query which gets a single user from the users database for who is logged in

		user, err := s.db.GetUser(context.Background(), currentUsername)
		if err != nil {
			return err
		}
		err = handler(s, cmd, user)
		return err
	}
}
