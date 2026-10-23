package main

import (
	"context"
	"fmt"
	"errors"
	"github.com/Warbo360/gator/internal/database"
	"github.com/google/uuid"
	"time"
)

// Handler function for the follow command, takes a URL arg so users can follow other feeds added by users other than
// themselves

func handlerFollow(s *state, cmd command, user database.User) error {
	
	if len(cmd.args) != 1 {
		return errors.New("usage: go run . follow <url>")
	}

	url := cmd.args[0]

	// SQL query for retrieving a feed by the passed URL

	feed, err := s.db.GetFeed(context.Background(), url)
	if err != nil {
		return err
	}

	// Creates a feedFollow entry in the feed_follows table which symbolizes a follow of a feed from a user each follow
	// is unique in the table

	feedFollowParam := database.CreateFeedFollowParams {
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feed,
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), feedFollowParam)
	if err != nil {
		return err
	}

	fmt.Printf("[UPDATE]: Feed %v followed by %v\n", feedFollow.FeedName, feedFollow.UserName)
	return nil
}
