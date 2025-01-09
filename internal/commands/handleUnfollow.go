package commands

import (
	"context"

	"github.com/chrispaul1/blog/internal/database"
)

func HandleUnfollow(s *State, cmd Command, user database.User) error {

	url := cmd.Args[0]
	feed, err := s.DB.GetFeedFromUrl(context.Background(), url)
	if err != nil {
		return err
	}

	deleteFeedstruct := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	_, err = s.DB.DeleteFeedFollow(context.Background(), deleteFeedstruct)
	if err != nil {
		return err
	}
	return nil
}
