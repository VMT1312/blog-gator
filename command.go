package main

import (
	"errors"
	"fmt"

	"github.com/VMT1312/blog-gator/internal/config"
)

type command struct {
	name string
	args []string
}

func handlerLogins(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New("missing username")
	}

	if s.cfg_pointer == nil {
		return errors.New("configuration not initialized")
	}

	if cmd.args[0] == "" {
		return errors.New("username cannot be empty")
	}

	s.cfg_pointer.CurrentUserName = cmd.args[0]

	fmt.Printf("Current user set to: %s\n", s.cfg_pointer.CurrentUserName)

	config.Write(*s.cfg_pointer)

	return nil
}
