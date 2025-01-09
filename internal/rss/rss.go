package rss

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"os"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error, could not get the feed %s", err)
		return nil, err
	}

	req.Header.Set("User-Agent", "gator")

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error,  %s", err)
		return nil, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error, data could not be read %s", err)
		return nil, err
	}

	var newRssFeed RSSFeed
	if err = xml.Unmarshal(data, &newRssFeed); err != nil {
		fmt.Fprintf(os.Stderr, "Error, rss struct could not be filled %s", err)
		return nil, err
	}

	newRssFeed.Channel.Title = html.UnescapeString(newRssFeed.Channel.Title)
	newRssFeed.Channel.Description = html.UnescapeString(newRssFeed.Channel.Description)
	for i := range newRssFeed.Channel.Item {
		newRssFeed.Channel.Item[i].Title = html.UnescapeString(newRssFeed.Channel.Item[i].Title)
		newRssFeed.Channel.Item[i].Description = html.UnescapeString(newRssFeed.Channel.Item[i].Description)
	}

	return &newRssFeed, nil
}
