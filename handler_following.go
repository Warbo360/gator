package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {

	current_username := s.cfg.CurrentUsername
	user, err := s.db.GetUser(context.Background(), current_username)
	if err != nil {
		return err
	}

	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	fmt.Println("============= Feeds for Current User =============")
	fmt.Printf("User: %v\n", user.Name)
	for _, feed := range feeds {
		fmt.Printf("Feed: %v\n", feed.FeedName)
	}

	return nil
}
