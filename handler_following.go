package main

import (
	"context"
	"fmt"
	"github.com/Warbo360/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {

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
