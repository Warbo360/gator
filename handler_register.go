package main

import (
	"errors"
	"fmt"
	"time"
	"github.com/Warbo360/gator/internal/database"
	"github.com/google/uuid"
	"database/sql"
	"context"
)

func handlerRegister(s *state, cmd command) error {

	if len(cmd.args) == 0 {
		return errors.New("name arg required")
	}

	newUserName := cmd.args[0] // New User Name
	_, err := s.db.GetUser(context.Background(), newUserName) // Check to see if this name exists in the database already
	if err == nil {
		return fmt.Errorf("ERROR: User of the name %v already exists", newUserName)
	} else if err == sql.ErrNoRows {

	} else if err != nil {
		return fmt.Errorf("ERROR: Failed to get user: %v", err)
	}

	newUserParams := database.createUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: newUserName,
	}


	newUser, err := s.db.createUser(context.Background(), newUserParams)
	if err != nil {
		return fmt.Errorf("ERROR:Failed to create new user: %v", err)
	}

	err = s.cfg.SetUser(newUserName)
	if err != nil {
		return err
	}

	fmt.Printf("UPDATE %v:%v:%v - User %v was created", newUser.CreatedAt, newUser.UpdatedAt, newUser.ID, newUser.Name)

	return nil

}
