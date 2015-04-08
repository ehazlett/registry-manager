package main

import (
	"fmt"
	"log"

	"github.com/codegangsta/cli"
)

func cmdInspectRepo(c *cli.Context) {
	client, err := getClient(c)
	if err != nil {
		log.Fatal(err)
	}

	repo := c.String("name")
	if repo == "" {
		log.Fatal("you must specify a repo")
	}

	r, err := client.Repository(repo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Namespace: %s\n", r.Namespace)
	fmt.Printf("Repository: %s\n", r.Repository)
	fmt.Print("Layers: \n")
	for _, layer := range r.Layers {
		fmt.Printf(" - ID: %s\n   Author: %s\n   Docker Version: %s\n   Size: %d\n",
			layer.ID,
			layer.Author,
			layer.DockerVersion,
			layer.Size,
		)
		fmt.Print("   History: \n")
		for _, cmd := range layer.ContainerConfig.Cmd {
			fmt.Printf("    - %s\n", cmd)
		}
	}
}
