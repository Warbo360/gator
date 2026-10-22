package main

import (
	"context"
	"fmt"
	"errors"
	"time"
	"github.com/google/uuid"
	"database/sql"
	"github.com/Warbo360/gator/internal/database"
	"github.com/lib/pq"
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
		postParams := database.MakePostParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title: item.Title,
			Url: item.Link,
			Description: itemDescParser(item.Description),
			PublishedAt: timeParser(item.PubDate),
			FeedID: feed.ID,
		}

		err = s.db.MakePost(context.Background(), postParams)
		pqErr, ok := err.(*pq.Error)
		if ok && pqErr.Code == "23505" {
		} else if ok {
			return pqErr
		} else if !ok && err != nil {
			return err
		}
	}

	fmt.Println("[UPDATE]: Feed successfully posted")
	return nil
}

func timeParser(timeString string) sql.NullTime {
	formats := []string{
		time.Layout,
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.Kitchen,
	}
	for _, format := range formats {
		parsedTime, err := time.Parse(format, timeString)
		if err == nil {
			return sql.NullTime{
				Time: parsedTime,
				Valid: true,
			}
		}
	}
	return sql.NullTime{
		Valid: false,
	}
}

func itemDescParser(itemDesc string) sql.NullString {
	if itemDesc == "" {
		return sql.NullString{
			Valid: false,
		}
	}
	return sql.NullString{
		String: itemDesc,
		Valid: true,
	}
}
