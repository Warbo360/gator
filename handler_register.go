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
		return errors.New("name argument required")
	}

	newUserName := cmd.args[0] // New User Name
	_, err := s.db.GetUser(context.Background(), newUserName) // Check to see if this name exists in the database already
	if err == nil {
		return fmt.Errorf("user of the name \"%v\" already exists", newUserName)
	} else if err == sql.ErrNoRows {

	} else if err != nil {
		return fmt.Errorf("failed to get user: \"%v\"", err)
	}

	newUserParams := database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: newUserName,
	}


	newUser, err := s.db.CreateUser(context.Background(), newUserParams)
	if err != nil {
		return fmt.Errorf("failed to create new user: \"%v\"", err)
	}

	err = s.cfg.SetUser(newUserName)
	if err != nil {
		return err
	}

	fmt.Printf("[UPDATE] %v:%v:%v - User \"%v\" was created\n", newUser.CreatedAt, newUser.UpdatedAt, newUser.ID, newUser.Name)

	return nil

}
