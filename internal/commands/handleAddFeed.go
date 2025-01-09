package commands

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/chrispaul1/blog/internal/database"
	"github.com/google/uuid"
)

func HandleAddFeed(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Error, not enough arguments, requires name and url\n")
		os.Exit(1)
	}
	// currentUser, err := s.DB.GetUser(context.Background(), s.C.User)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error, could not get user struct '%v'\n", err)
	// 	return err
	// }

	feedStruct := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    user.ID,
	}

	newFeed, err := s.DB.CreateFeed(context.Background(), feedStruct)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: feed was not created '%v'\n", err)
		return err
	}

	feedFollowStruct := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    newFeed.UserID,
		FeedID:    newFeed.ID,
	}
	_, err = s.DB.CreateFeedFollow(context.Background(), feedFollowStruct)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: feed follow struct was not created '%v'\n", err)
		return err
	}

	fmt.Println(newFeed)

	return nil
}

func HandleFeeds(s *State, cmd Command) error {
	feeds, err := s.DB.GetFeeds(context.Background())
	if len(feeds) == 0 {
		fmt.Println("No feeds were found")
		os.Exit(0)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "error, feeds could not be retrieved '%v'\n", err)
		return err
	}

	for _, feed := range feeds {

		fmt.Printf("Feed Name : %s\n", feed.Name)
		fmt.Printf("Url : %s\n", feed.Url)
		user, err := s.DB.GetUserFromID(context.Background(), feed.UserID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error, feed's creator could not be found '%v'\n", err)
		}
		fmt.Printf("Feed Creator : %s\n", user.Name)
	}

	return nil
}
