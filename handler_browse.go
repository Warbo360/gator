package main

import (
	"fmt"
	"strconv"
	"errors"
	"context"
	"github.com/Warbo360/gator/internal/database"
)

// Handler function for the browse command, will print to the CLI the posts the user has made from followed/added feeds,
// takes a single arg in a limit of posts to see, defaults to 2 if none are given

func handlerBrowse(s *state, cmd command, user database.User) error {

	limit := 2
	argsN := len(cmd.args)
	if argsN == 0 {
	} else if argsN > 0 {
		newLimit, err := parseLimit(cmd.args[0])

		// If an arg is given but can not be parsed (helper function below) halts the program and prints proper usage of
		// the command

		if err != nil {
			return errors.New("usage: browse <number of posts to limit to>")
		}
		limit = newLimit
	}

	// SQL query to select posts made by a user with a limit of limit set above

	posts, err := s.db.GetPosts(context.Background(), database.GetPostsParams{UserID: user.ID, Limit: int32(limit)})
	if err != nil {
		return err
	}

	// Prints to the CLI the retrieved posts from above in a formatted way

	fmt.Println()
	fmt.Printf("Posts for \"%v\":\n", user.Name)
	fmt.Println()
	for idx, post := range posts {
		fmt.Printf("Post %v: %v\n", idx + 1, post.Title)
		fmt.Printf("Created At: %v\n", post.CreatedAt)
		fmt.Printf("Updated At: %v\n", post.UpdatedAt)
		fmt.Printf("URL: %v\n", post.Url)
		if post.PublishedAt.Valid {
			fmt.Printf("Published at: %v\n", post.PublishedAt.Time)
		}
		if post.Description.Valid {
			fmt.Printf("Description: %v\n", post.Description.String)
		}
		fmt.Println()
	}
	return nil
}

// Helper function for parsing string arg of the command into an int

func parseLimit(num string) (int, error) {
	numInt, err := strconv.Atoi(num)
	if err != nil {
		return 0, err
	}
	return numInt, nil
}
