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

// Handler function for aggregating feeds added by a user takes a single arg for how often to scrape the feeds

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

	// Creates a ticker that runs every interval passed by the user, then every interval scrapes the feeds with the
	// helper function below

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

// Helper function abstracting out a singlular scrape of a feed

func scrapeFeeds(s *state) error {

	// SQL query to fetch the oldest updated feed

	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	// SQL query to update the time of the feed

	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return err
	}

	// Function that makes the HTTP request to fetch the feed and unmarshal the resulting XML

	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}

	// In the feed Channel struct, for each Item struct in the slice of Items creates the make post parameters and
	// makes SQL query to create a post value into the posts table
	// Of note as well are the itemDescParser and timeParser helper functions below that construct NullTime and
	// NullString structs for possible empty/nil descriptions and publish dates of the resulting XML

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

		// Skips if trying a post that already exists which is expected to happen in the event the error is
		// *pq.Error.Code == "23505", note if the error is a *pq.Error but the code is not "23505" then we still address
		// the error.

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

// Helper function for parsing timeStrings from resulting feed response into a sql.Nulltime incase fetched feed has none
// or format is not in the listed supported formats below

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

	// For each format see if it matches the format of the timeString, if so set the time in that format and mark it as
	// a valid time, if not just make it an invalid field

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

// Helper function for parsing feed descriptions if they are empty then we just set the field as false, if true we just
// feed the string as is and mark true

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
