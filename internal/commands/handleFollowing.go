package commands

import (
	"context"
	"fmt"

	"github.com/chrispaul1/blog/internal/database"
)

func HandleFollowing(s *State, cmd Command, user database.User) error {
	// currentUser, err := s.DB.GetUser(context.Background(), s.C.User)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error, could not get user struct '%v'\n", err)
	// 	return err
	// }

	allFeeds, err := s.DB.GetFeedFollowsForUser(context.Background(), user.Name)
	if err != nil {
		return err
	}

	for _, item := range allFeeds {
		fmt.Printf("%s\n", item.FeedName)
	}
	return nil
}
