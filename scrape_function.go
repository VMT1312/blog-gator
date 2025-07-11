package main

import (
	"context"
	"database/sql"
	"time"

	"github.com/VMT1312/blog-gator/internal/database"
	"github.com/google/uuid"
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
		post := database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: feed.CreatedAt,
			UpdatedAt: feed.UpdatedAt,
			Title:     item.Title,
			Url:       feed.Url,
			Description: sql.NullString{
				String: item.Description,
			},
			PublishedAt: sql.NullString{
				String: item.PubDate,
			},
			FeedID: feed.ID,
		}

		if err := s.db.CreatePost(context.Background(), post); err != nil {
			return err
		}
	}

	return nil
}
