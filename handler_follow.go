package main

import (
	"context"
	"fmt"
	"errors"
	"github.com/Warbo360/gator/internal/database"
	"github.com/google/uuid"
	"time"
)

func handlerFollow(s *state, cmd command) error {
	
	if len(cmd.args) != 1 {
		return errors.New("usage: go run . follow <url>")
	}

	url := cmd.args[0]
	username := s.cfg.CurrentUsername
	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return err
	}

	feed, err := s.db.GetFeed(context.Background(), url)
	if err != nil {
		return err
	}

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
