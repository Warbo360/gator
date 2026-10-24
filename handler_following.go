package main

import (
	"context"
	"fmt"

	"github.com/Warbo360/gator/internal/database"
)

// Handler function for the following command which prints all feeds being followed by the current user

func handlerFollowing(s *state, cmd command, user database.User) error {
	// SQL query that uses the current user ID to get all followed feeds

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
