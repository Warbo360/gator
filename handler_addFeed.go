package main

import (
	"errors"
	"github.com/Warbo360/gator/internal/database"
	"fmt"
	"context"
	"github.com/google/uuid"
	"time"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {

	// Check for two command args, those being the feeds name and feeds url

	if len(cmd.args) != 2 {
		return errors.New("feed name and url needed, usage: addfeed <name> <url>")
	}

	// Builds params for MakeFeed below

	name := cmd.args[0]
	url := cmd.args[1]

 	feedParams := database.MakeFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: name,
		Url: url,
		UserID: user.ID,
	}

	// MakeFeed query which inserts a entry into the feeds table

	newFeed, err := s.db.MakeFeed(context.Background(), feedParams)
	if err != nil {
		return err
	}

	feedUUID, err := s.db.GetFeed(context.Background(), newFeed.Url)
	if err != nil {
		return err
	}

	// Creates a feed follow automatically for when a feed is added by a user

	feedFollowParam := database.CreateFeedFollowParams {
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feedUUID,
	}

	_, err = s.db.CreateFeedFollow(context.Background(), feedFollowParam)
	if err != nil {
		return err
	}

	fmt.Printf("%+v", newFeed)

	return nil
}
