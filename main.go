package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/VMT1312/blog-gator/internal/config"
	"github.com/VMT1312/blog-gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)

	s := state{
		cfg: &cfg,
		db:  dbQueries,
	}

	cmds := commands{
		callback: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogins)
	cmds.register("register", handlerRegisters)
	cmds.register("reset", handlerResets)
	cmds.register("users", handlerGetUsers)
	cmds.register("agg", handlerFecthFeed)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", func(s *state, cmd command) error {
		return s.feeds()
	})
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))

	args := os.Args
	if len(args) < 2 {
		log.Fatal("Did not provide a command name")
	}

	cmd := command{
		name: args[1],
		args: args[2:],
	}

	err = cmds.run(&s, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
