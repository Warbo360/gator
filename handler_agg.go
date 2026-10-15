package main

import (
	"context"
	"fmt"
)

const URL string = "https://www.wagslane.dev/index.xml"

func handlerAgg(s *state, cmd command) error {

	// command handler for aggregating service

	feed, err := fetchFeed(context.Background(), URL)
	if err != nil {
		return err
	}
	fmt.Printf("%+v", feed)
	return nil
}
