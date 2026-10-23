package main

import (
	"context"
	"fmt"
)

// Handler function for the feeds command which prints all the feeds added to the database by all users

func handlerFeeds(s *state, cmd command) error {

	// Access state to db to then call GetFeeds an SQL query selecting all values from the feeds table

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	// Prints all feeds in a formatted way

	for _, feed := range feeds {
		name := feed.Name
		url := feed.Url
		userID := feed.Name_2
		fmt.Printf("|| Feed Name: %v || Feed URL: %v || Feed Added by User: %v ||\n", name, url, userID)
	}

	return nil
}
