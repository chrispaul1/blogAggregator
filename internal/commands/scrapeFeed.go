package commands

import (
	"context"
	"database/sql"
	"time"

	"github.com/chrispaul1/blog/internal/database"
	"github.com/chrispaul1/blog/internal/rss"
	"github.com/google/uuid"
	"github.com/lib/pq"
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

	layouts := []string{
		time.RFC3339,
		time.RFC1123,
		time.RFC822,
		"2006-01-02 15:04:05",
		"2006-01-02",
		"02-01-2006",
		"Jan 2, 2006",
	}

	for _, item := range rssFeed.Channel.Item {
		for _, layout := range layouts {
			timeVal, err := time.Parse(layout, item.PubDate)
			if err != nil {
				continue
			} else {
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
					if pgErr, ok := err.(*pq.Error); ok && pgErr.Code == "23505" {
						continue
					} else {
						return err
					}
				}
				break
			}
		}
	}

	return nil
}
