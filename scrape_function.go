package main

import (
	"context"
	"fmt"
	"time"

	"github.com/VMT1312/blog-gator/internal/database"
)

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	arg := database.MarkFeedFetchedParams{
		UpdatedAt: time.Now(),
		ID:        feed.ID,
	}

	if err := s.db.MarkFeedFetched(context.Background(), arg); err != nil {
		return err
	}

	feedData, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}

	for _, item := range feedData.Channel.Item {
		fmt.Printf("Title: %s\n", item.Title)
	}

	return nil
}
