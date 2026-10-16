package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {

	// Access state to db to then call GetFeeds

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		name := feed.Name
		url := feed.Url
		userID := feed.Name_2
		fmt.Printf("|| Feed Name: %v || Feed URL: %v || Feed Added by User: %v ||\n", name, url, userID)
	}

	return nil
}
