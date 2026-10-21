package main

import (
	"context"
	"github.com/Warbo360/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		currentUsername := s.cfg.CurrentUsername
		user, err := s.db.GetUser(context.Background(), currentUsername)
		if err != nil {
			return err
		}
		err = handler(s, cmd, user)
		return err
	}
}
