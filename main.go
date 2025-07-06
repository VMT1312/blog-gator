package main

import (
	"log"
	"os"

	"github.com/VMT1312/blog-gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	s := state{cfg_pointer: &cfg}

	cmds := commands{
		callback: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogins)

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
