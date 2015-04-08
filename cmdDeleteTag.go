package main

import (
	"log"

	"github.com/codegangsta/cli"
)

func cmdDeleteTag(c *cli.Context) {
	client, err := getClient(c)
	if err != nil {
		log.Fatal(err)
	}

	repo := c.String("name")
	tag := c.String("tag")
	if repo == "" {
		log.Fatal("you must specify a repo")
	}

	if tag == "" {
		log.Fatal("you must specify a tag")
	}

	if err := client.DeleteTag(repo, tag); err != nil {
		log.Fatal(err)
	}
}
