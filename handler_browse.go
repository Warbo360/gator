package main

import (
	"fmt"
	"strconv"
	"errors"
	"context"
	"github.com/Warbo360/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {

	limit := 2
	argsN := len(cmd.args)
	if argsN == 0 {
		
	} else if argsN > 0 {
		newLimit, err := parseLimit(cmd.args[0])
		if err != nil {
			return errors.New("usage: browse <number of posts to limit to>")
		}
		limit = newLimit
	}

	posts, err := s.db.GetPosts(context.Background(), database.GetPostsParams{UserID: user.ID, Limit: int32(limit)})
	if err != nil {
		return err
	}
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

func parseLimit(num string) (int, error) {
	numInt, err := strconv.Atoi(num)
	if err != nil {
		return 0, err
	}
	return numInt, nil
}
