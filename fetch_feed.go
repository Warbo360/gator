package main

import (
	"context"
	"net/http"
	"io"
	"encoding/xml"
	"html"
)

// Fetchs RSS feeds with HTTP requests and unmarshals the XML response if there is one

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {

	// Custom client for configuring header on the request

	client := &http.Client{}

	// Custom request using provided context, GET protocol, passed URL

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}

	// Sets header info on the request to let server know what is requesting the info

	req.Header.Set("User-Agent", "gator")

	// Response from sent request

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Reading response body

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshaling response body XML into RSSFeed struct

	var rssFeed RSSFeed
	if err := xml.Unmarshal(body, &rssFeed); err != nil {
		return nil, err
	}

	// Helper function below that unescapes XML tags

	rssFeed.unescapeHTML()

	return &rssFeed, nil
}

// Helper function that takes in an RSSFeed pointer and modifies channels title and description and well as channel's
// item's title and description with html.UnescapeString()

func (r *RSSFeed) unescapeHTML() {

	r.Channel.Title = html.UnescapeString(r.Channel.Title)
	r.Channel.Description = html.UnescapeString(r.Channel.Description)
	for i := range r.Channel.Item {
		r.Channel.Item[i].Title = html.UnescapeString(r.Channel.Item[i].Title)
		r.Channel.Item[i].Description = html.UnescapeString(r.Channel.Item[i].Description)
	}
}
