package main

import (
	"context"
	"fmt"
	"errors"
	"time"
)

func handlerAgg(s *state, cmd command) error {

	// Check to see if a time interval was set

	if len(cmd.args) != 1 {
		return errors.New("usage: agg <time_between_requests>")
	}

	// Set interval time to scrapeFeeds

	interval, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Scraping feeds every %v...\n", interval)
	ticker := time.NewTicker(interval)
	for ; ; <-ticker.C {
		err = scrapeFeeds(s)
		if err != nil {
			return err
		}
	}

	return nil
}

func scrapeFeeds(s *state) error {

	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return err
	}

	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}

	for _, item := range rssFeed.Channel.Item {
		fmt.Printf("Feed Title: %v\n", item.Title)
	}
	return nil
}
