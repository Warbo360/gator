package main

import (
	"context"
	"net/http"
	"io"
	"encoding/xml"
)

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {

	client := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "gator")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rssFeed RSSFeed
	if err := xml.Unmarshal(body, &rssFeed); err != nil {
		return nil, err
	}

	return &rssFeed, nil
}

func (r *RSSFeed) unescapeHTML() {
	// Will be the helper function that will unescape HTML entities with html.UnescapeString()
}
