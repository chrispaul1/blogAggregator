package commands

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/chrispaul1/blog/internal/database"
)

func HandleBrowse(s *State, cmd Command, user database.User) error {
	var limit int = 2
	if len(cmd.Args) > 0 {
		userLimit, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error, expecting a number '%v'", err)
			return err
		} else {
			limit = userLimit
		}
	}

	getPostStruct := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	}
	posts, err := s.DB.GetPostsForUser(context.Background(), getPostStruct)
	if err != nil {
		return err
	}
	for i, post := range posts {
		fmt.Printf("\n=== Post : %d ===\n", i)
		fmt.Printf("Title : %s\n", post.Title)
		fmt.Printf("Description : %v\n", post.Description)
		fmt.Printf("Link to Read More : %s\n", post.Url)
		fmt.Printf("Pub Date : %s\n", post.PublishedAt)
		fmt.Printf("-------------\n")
	}

	return nil
}
