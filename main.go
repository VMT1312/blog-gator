package main

import (
	"fmt"
	"log"

	"github.com/VMT1312/blog-gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	err = cfg.SetUser("Vinh")
	if err != nil {
		log.Fatal(err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("database URL: %s\n", cfg.DbUrl)
	fmt.Printf("current user: %s\n", cfg.CurrentUserName)
}
