package main

import (
	"fmt"
	"context"
)

// Handler function for printing all registered users to the CLI while indicating who is currently logged in

func handlerUsers(s *state, cmd command) error {

	// SQL query that gets all users from the users table

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		if user == s.cfg.CurrentUsername {
			fmt.Printf("* %v (current)\n", user)
		} else {
			fmt.Printf("* %v\n", user)
		}
	}
	return nil

}
