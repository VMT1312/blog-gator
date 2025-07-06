package main

import (
	"github.com/VMT1312/blog-gator/internal/config"
	"github.com/VMT1312/blog-gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}
