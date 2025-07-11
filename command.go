package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/VMT1312/blog-gator/internal/config"
	"github.com/VMT1312/blog-gator/internal/database"
	"github.com/google/uuid"
)

type command struct {
	name string
	args []string
}

func handlerLogins(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New("missing username")
	}

	if s.cfg == nil {
		return errors.New("configuration not initialized")
	}

	if cmd.args[0] == "" {
		return errors.New("username cannot be empty")
	}

	_, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return errors.New("couldn't find user")
		}
		return err
	}

	s.cfg.CurrentUserName = cmd.args[0]

	fmt.Printf("Current user set to: %s\n", s.cfg.CurrentUserName)

	config.Write(*s.cfg)

	return nil
}

func handlerRegisters(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New("missing username")
	}

	param := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	}

	user, err := s.db.CreateUser(context.Background(), param)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return errors.New("user already exists")
		}
		return err
	}

	s.cfg.CurrentUserName = user.Name
	config.Write(*s.cfg)

	fmt.Printf("%s was added successfully to the database\n", cmd.args[0])
	log.Printf(" - ID: %v", user.ID)
	log.Printf(" - Created at: %v", user.CreatedAt)
	log.Printf(" - Updatead at: %v", user.UpdatedAt)
	log.Printf(" - Name: %s", user.Name)

	return nil
}

func handlerResets(s *state, cmd command) error {
	err := s.db.ResetDb(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("Table has been reset successfully")
	return nil
}

func handlerGetUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}

	return nil
}

func handlerFecthFeed(s *state, cmd command) error {
	// if len(cmd.args) < 1 {
	// 	return errors.New("missing url to fetch")
	// }

	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	fmt.Print(feed)

	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		return errors.New("haven't provide both arguments; need name and url of the feed")
	}

	current_user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	arg := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    current_user.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), arg)
	if err != nil {
		return err
	}

	followCmd := command{
		name: "follow",
		args: []string{cmd.args[1]},
	}
	if err = handlerFollow(s, followCmd); err != nil {
		return err
	}

	fmt.Println(feed)

	return nil
}

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New("missing feed url")
	}

	current_user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	arg := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    current_user.ID,
		FeedID:    feed.ID,
	}

	feed_follow, err := s.db.CreateFeedFollow(context.Background(), arg)
	if err != nil {
		return err
	}

	fmt.Println(feed_follow)
	return nil
}

func handlerFollowing(s *state, cmd command) error {
	current_user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	following_feeds, err := s.db.GetFeedFollowsForUser(context.Background(), current_user.ID)
	if err != nil {
		return err
	}

	for _, feed := range following_feeds {
		fmt.Println(feed.FeedName)
	}
	return nil
}
