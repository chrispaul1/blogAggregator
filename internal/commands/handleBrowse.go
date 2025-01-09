package commands

import (
	"context"
	"fmt"
	"os"
	"strconv"
)

func HandleBrowse(s *State, cmd Command) error {
	if len(cmd.Args) > 0 {
		postLimit, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error, expecting a number '%v'", err)
			return err
		}

		posts, err := s.DB.GetPostsForUser(context.Background(), int32(postLimit))
		if err != nil {
			return err
		}
		for _, post := range posts {
			fmt.Printf("Title : %s\n", post.Title)
			fmt.Printf("Description : %v\n", post.Description)
			fmt.Printf("Link : %s\n", post.Url)
			fmt.Printf("Pub Date : %s\n", post.PublishedAt)
		}
	} else {
		posts, err := s.DB.GetPostsForUser(context.Background(), 2)
		if err != nil {
			return err
		}
		for _, post := range posts {
			fmt.Printf("Title : %s\n", post.Title)
			fmt.Printf("Description : %v\n", post.Description)
			fmt.Printf("Link : %s\n", post.Url)
			fmt.Printf("Pub Date : %s\n", post.PublishedAt)
		}
	}
	return nil
}
