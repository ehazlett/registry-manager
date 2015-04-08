package main

import (
	"log"

	"github.com/codegangsta/cli"
)

func cmdDeleteRepo(c *cli.Context) {
	client, err := getClient(c)
	if err != nil {
		log.Fatal(err)
	}

	repo := c.String("name")
	if repo == "" {
		log.Fatal("you must specify a repo")
	}

	if err := client.DeleteRepository(repo); err != nil {
		log.Fatal(err)
	}
}
