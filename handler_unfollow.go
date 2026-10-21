package main

import (
	"fmt"
	"github.com/Warbo360/gator/internal/database"
	"errors"
	"context"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {

	if len(cmd.args) != 1 {
		return errors.New("usage: unfollow <url>")
	}

	url := cmd.args[0]

	feed, err := s.db.GetFeed(context.Background(), url)
	if err != nil {
		return err
	}

	feedFollowParams := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed,
	}

	err = s.db.DeleteFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return err
	}

	fmt.Println("[UPDATE]: Feed sucessfully unfollowed")
	return nil
}
