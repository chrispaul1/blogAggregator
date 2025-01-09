package commands

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/chrispaul1/blog/internal/database"
	"github.com/chrispaul1/blog/internal/rss"
	"github.com/google/uuid"
)

func ScrapeFeed(s *State, cmd Command) error {
	feed, err := s.DB.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	markFeedStruct := database.MarkFeedFetchedParams{
		ID: feed.ID,
		LastFetchedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}
	markedFeed, err := s.DB.MarkFeedFetched(context.Background(), markFeedStruct)
	if err != nil {
		return err
	}

	rssFeed, err := rss.FetchFeed(context.Background(), markedFeed.Url)
	if err != nil {
		return err
	}

	fmt.Printf("\n=== %s ===\n", rssFeed.Channel.Title)
	timeFormat := "Tue, 06 sep 2021 11:30:00"
	for _, item := range rssFeed.Channel.Item {
		timeVal, err := time.Parse(timeFormat, item.PubDate)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error, time could not be formated")
			continue
		}
		postStruct := database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title:     item.Title,
			Url:       item.Link,
			Description: sql.NullString{
				String: item.Description,
				Valid:  true,
			},
			PublishedAt: timeVal,
			FeedID:      markedFeed.ID,
		}
		_, err = s.DB.CreatePost(context.Background(), postStruct)
		if err != nil {
			return err
		}
	}

	return nil
}
