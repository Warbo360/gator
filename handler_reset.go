package main

import (
	"context"
	"fmt"
)

// Handler function for the reset command, which deletes all users from the users table, in conjunction with the CASCADE
// contraints on the other tables that are dependent on users will delete all entries from all other tables as well
// resetting the database, intended for testing only.

func handlerReset(s *state, cmd command) error {

	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("[UPDATE]: \"users\" SQL table reset succesfully")
	return nil

}
