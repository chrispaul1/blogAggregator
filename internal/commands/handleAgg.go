package commands

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func Agg(s *State, cmd Command) error {

	if len(cmd.Args) < 1 {
		return errors.New("error: time parameter is required")
	}

	time_string := cmd.Args[0]
	time_between_reqs, err := time.ParseDuration(time_string)
	if err != nil {
		return err
	}

	ticker := time.NewTicker(time_between_reqs)
	for ; ; <-ticker.C {
		if err := ScrapeFeed(s, cmd); err != nil {
			fmt.Fprintf(os.Stderr, "Error scrapping feed: %v\n", err)
			continue
		}
	}

	// feedStruct, err := rss.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error, could not retrieve feed %s", err)
	// 	return err
	// }
	// fmt.Printf("%v\n", feedStruct)
	//fmt.Printf("\n%s", feedStruct.Channel.Title)
	//fmt.Printf("\n%s", feedStruct.Channel.Description)
}
