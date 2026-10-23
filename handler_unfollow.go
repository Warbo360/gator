package main

import (
	"fmt"
	"github.com/Warbo360/gator/internal/database"
	"errors"
	"context"
)

// Handler function for unfollow command which deletes a feed_follow from the feed_follows table, gets the feed by
// passed URL arg

func handlerUnfollow(s *state, cmd command, user database.User) error {

	if len(cmd.args) != 1 {
		return errors.New("usage: unfollow <url>")
	}

	url := cmd.args[0]

	// SQL query to retrive the feed by passed URL

	feed, err := s.db.GetFeed(context.Background(), url)
	if err != nil {
		return err
	}

	// Deletes the user-feed relation by deleting the feed_follow entry for the user and the feed, note this does not
	// delete the feed itself from the database and would still be indicated to be added by the user who did add it even
	// if they unfollow it, now if the user is deleted from the users table, then the feeds added by that user would
	// also be deleted

	feedFollowParams := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed,
	}

	err = s.db.DeleteFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return err
	}

	fmt.Println("[UPDATE]: Feed sucessfully unfollowed")
	return nil
}
