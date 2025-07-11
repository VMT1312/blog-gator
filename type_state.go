package main

import (
	"context"
	"fmt"

	"github.com/VMT1312/blog-gator/internal/config"
	"github.com/VMT1312/blog-gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func (s *state) feeds() error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		user, err := s.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return err
		}

		fmt.Printf("Feed name: %s\n", feed.Name)
		fmt.Printf("Feed url: %s\n", feed.Url)
		fmt.Printf("Created user: %s\n", user)
	}

	return nil
}
