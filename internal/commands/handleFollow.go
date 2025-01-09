package commands

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/chrispaul1/blog/internal/database"
	"github.com/google/uuid"
)

func HandleFollow(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) < 1 {
		fmt.Fprintf(os.Stderr, "expecting a url")
	}

	url := cmd.Args[0]

	// currentUser, err := s.DB.GetUser(context.Background(), s.C.User)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error, could not get user struct '%v'\n", err)
	// 	return err
	// }

	urlFeed, err := s.DB.GetFeedFromUrl(context.Background(), url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error, could not get feed struct '%v'\n", err)
		return err
	}

	feedFollowStruct := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    urlFeed.ID,
	}

	newFeedFollow, err := s.DB.CreateFeedFollow(context.Background(), feedFollowStruct)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: feed follow struct was not created '%v'\n", err)
		return err
	}

	fmt.Printf("Feed : %s\n", newFeedFollow.FeedName)
	fmt.Printf("Username : %s\n", newFeedFollow.UserName)

	return nil
}
